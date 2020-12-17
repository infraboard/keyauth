package mongo

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg/session"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/verifycode"
)

func (s *service) IssueToken(req *token.IssueTokenRequest) (*token.Token, error) {
	// 检查安全性
	if err := s.securityCheck(req); err != nil {
		return nil, exception.NewBadRequest("security check failed, %s", err)
	}

	// 颁发Token
	tk, err := s.issuer.IssueToken(req)
	if err != nil {
		s.checker.UpdateFailedRetry(req)
		return nil, err
	}
	tk.WithRemoteIP(req.GetRemoteIP())
	tk.WithUerAgent(req.GetUserAgent())

	// 登录会话
	sess, err := s.session.Login(tk)
	if err != nil {
		return nil, err
	}
	tk.SessionID = sess.ID

	// 保存入库
	if err := s.saveToken(tk); err != nil {
		return nil, err
	}

	return tk, nil
}

func (s *service) securityCheck(req *token.IssueTokenRequest) error {
	// 如果有校验码, 则直接通过校验码检测用户身份安全
	if req.VerifyCode != "" {
		s.log.Debugf("verify code provided, check code ...")
		_, err := s.code.CheckCode(verifycode.NewCheckCodeRequest(req.VerifyCode))
		if err != nil {
			return exception.NewPermissionDeny("verify code invalidate, error, %s", err)
		}
		s.log.Debugf("verfiy code check passed")
		return nil
	}

	// 连续登录失败检测
	if err := s.checker.MaxFailedRetryCheck(req); err != nil {
		return exception.NewBadRequest("max retry error, %s", err)
	}

	// 异地登录检测
	err := s.checker.OtherPlaceLoggedInChecK(req)
	if err != nil {
		return err
	}

	// 长时间未登录检测
	err = s.checker.NotLoginDaysChecK(req)
	if err != nil {
		return err
	}

	// IP保护检测
	err = s.checker.IPProtectCheck(req)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) ValidateToken(req *token.ValidateTokenRequest) (*token.Token, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	tk, err := s.describeToken(newDescribeTokenRequest(req.DescribeTokenRequest))
	if err != nil {
		return nil, exception.NewUnauthorized(err.Error())
	}

	if tk.IsBlock {
		return nil, s.makeBlockExcption(tk.BlockType, tk.BlockMessage())
	}

	// 校验Token是否过期
	if req.AccessToken != "" {
		if tk.CheckAccessIsExpired() {
			return nil, exception.NewAccessTokenExpired("access_token: %s has expired", tk.AccessToken)
		}
	}

	if req.RefreshToken != "" {
		if tk.CheckRefreshIsExpired() {
			return nil, exception.NewRefreshTokenExpired("refresh_token: %s expoired", tk.RefreshToken)
		}
	}

	tk.Desensitize()
	return tk, nil
}

func (s *service) makeBlockExcption(bt token.BlockType, message string) exception.APIException {
	switch bt {
	case token.OtherClientLoggedIn:
		return exception.NewOtherClientsLoggedIn(message)
	case token.SessionTerminated:
		return exception.NewSessionTerminated(message)
	case token.OtherPlaceLoggedIn:
		return exception.NewOtherPlaceLoggedIn(message)
	case token.OtherIPLoggedIn:
		return exception.NewOtherIPLoggedIn(message)
	default:
		return exception.NewInternalServerError("unknow block type: %s, message: %s", bt, message)
	}
}

func (s *service) BlockToken(req *token.BlockTokenRequest) (*token.Token, error) {
	tk, err := s.DescribeToken(token.NewDescribeTokenRequestWithAccessToken(req.AccessToken))
	if err != nil {
		return nil, fmt.Errorf("query session access token error, %s", err)
	}

	tk.IsBlock = true
	tk.BlockType = req.BlcokType
	tk.BlockReason = req.BlockReson
	tk.BlockAt = ftime.Now()

	if err := s.updateToken(tk); err != nil {
		return nil, err
	}
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

	// 退出会话
	if req.LogoutSession {
		logoutReq := session.NewLogoutRequest(tk.SessionID)
		if err := s.session.Logout(logoutReq); err != nil {
			return exception.NewInternalServerError("logout session error, %s", err)
		}
	}

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
