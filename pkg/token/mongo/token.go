package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg/audit"
	"github.com/infraboard/keyauth/pkg/token"
)

func (s *service) IssueToken(req *token.IssueTokenRequest) (*token.Token, error) {
	fl := s.inspectRequest(req)
	if err := fl.CheckBlook(); err != nil {
		return nil, exception.NewBadRequest("inspect request error, %s", err)
	}

	tk, err := s.issuer.IssueToken(req)
	if err != nil {
		s.saveAbnormalLogin(req, fl)
		return nil, err
	}

	if _, err := s.col.InsertOne(context.TODO(), tk); err != nil {
		return nil, exception.NewInternalServerError("inserted token(%s) document error, %s",
			tk.AccessToken, err)
	}

	s.saveLoginLog(req, tk)
	return tk, nil
}

func (s *service) inspectRequest(req *token.IssueTokenRequest) *FailedLogin {
	fl := NewFailedLogin()
	s.cache.Get(req.AbnormalUserCheckKey(), fl)
	return fl
}

// 记录登录失败的次数
func (s *service) saveAbnormalLogin(req *token.IssueTokenRequest, fl *FailedLogin) {
	fl.Inc()
	s.cache.PutWithTTL("abnormal_"+req.Username, fl, s.retryTTL)
}

func (s *service) saveLoginLog(req *token.IssueTokenRequest, tk *token.Token) {
	data := audit.NewDefaultLoginLogData()

	data.Account = tk.Account
	data.AccountType = tk.UserType
	data.ApplicationID = tk.ApplicationID
	data.ApplicationName = tk.ApplicationName
	data.GrantType = tk.GrantType
	data.LoginIP = req.GetRemoteIP()

	data.WithUserAgent(req.GetUserAgent())
	data.WithToken(tk)

	s.audit.SaveLoginRecord(data)
	return
}

func (s *service) saveLogoutLog(tk *token.Token) {
	data := audit.NewDefaultLogoutLogData()
	data.Account = tk.Account
	data.ApplicationID = tk.ApplicationID
	data.ApplicationName = tk.ApplicationName
	data.GrantType = tk.GrantType
	data.GrantType = *tk.GetStartGrantType()

	data.WithToken(tk)
	if tk.CheckRefreshIsExpired() {
		data.LogoutAt = tk.RefreshExpiredAt
	}
	s.audit.SaveLoginRecord(data)
	return
}

func (s *service) ValidateToken(req *token.ValidateTokenRequest) (*token.Token, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	tk, err := s.describeToken(newDescribeTokenRequest(req.DescribeTokenRequest))
	if err != nil {
		return nil, exception.NewUnauthorized(err.Error())
	}

	// 校验Token是否过期
	if req.AccessToken != "" {
		if tk.CheckAccessIsExpired() {
			return nil, exception.NewAccessTokenExpired("access_token: %s has expired", tk.AccessToken)
		}
	}

	if req.RefreshToken != "" {
		if tk.CheckRefreshIsExpired() {
			// 如果token过期了记录退出日志
			s.saveLogoutLog(tk)
			return nil, exception.NewRefreshTokenExpired("refresh_token: %s expoired", tk.RefreshToken)
		}
	}

	tk.Desensitize()
	return tk, nil
}

func (s *service) DescribeToken(req *token.DescribeTokenRequest) (*token.Token, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	tk, err := s.describeToken(newDescribeTokenRequest(req))
	if err != nil {
		return nil, exception.NewUnauthorized(err.Error())
	}

	return tk, nil
}

func (s *service) QueryToken(req *token.QueryTokenRequest) (*token.Set, error) {
	query := newQueryRequest(req)
	resp, err := s.col.Find(context.TODO(), query.FindFilter(), query.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find token error, error is %s", err)
	}

	tokenSet := token.NewTokenSet(req.PageRequest)
	// 循环
	for resp.Next(context.TODO()) {
		tk := new(token.Token)
		if err := resp.Decode(tk); err != nil {
			return nil, exception.NewInternalServerError("decode token error, error is %s", err)
		}
		tokenSet.Add(tk)
	}

	// count
	count, err := s.col.CountDocuments(context.TODO(), query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get token count error, error is %s", err)
	}
	tokenSet.Total = count

	return tokenSet, nil

}

func (s *service) RevolkToken(req *token.RevolkTokenRequest) error {
	if err := req.Validate(); err != nil {
		return exception.NewBadRequest(err.Error())
	}

	// 检测撤销token的客户端是否合法
	app, err := s.issuer.CheckClient(req.ClientID, req.ClientSecret)
	if err != nil {
		return exception.NewUnauthorized(err.Error())
	}

	// 检测被撤销token的合法性
	descReq := newDescribeTokenRequest(req.DescribeTokenRequest)
	tk, err := s.describeToken(descReq)
	if err != nil {
		return err
	}

	if err := tk.CheckTokenApplication(app.ID); err != nil {
		return exception.NewPermissionDeny(err.Error())
	}

	// 记录退出日志
	s.saveLogoutLog(tk)
	return s.destoryToken(descReq)
}

func (s *service) destoryToken(req *describeTokenRequest) error {
	resp, err := s.col.DeleteOne(context.TODO(), req.FindFilter())
	if err != nil {
		return exception.NewInternalServerError("delete token(%s) error, %s", req, err)
	}

	if resp.DeletedCount == 0 {
		return exception.NewNotFound("token(%s) not found", req)
	}

	return nil
}

func (s *service) describeToken(req *describeTokenRequest) (*token.Token, error) {
	tk := new(token.Token)

	if err := s.col.FindOne(context.TODO(), req.FindFilter()).Decode(tk); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("token %s not found", req)
		}

		return nil, exception.NewInternalServerError("find token %s error, %s", req, err)
	}

	return tk, nil
}
