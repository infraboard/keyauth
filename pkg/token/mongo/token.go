package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg/session"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/verifycode"
)

func (s *service) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	// 连续登录失败检测
	if err := s.loginBeforeCheck(req); err != nil {
		return nil, exception.NewBadRequest("安全检测失败, %s", err)
	}

	// 颁发Token
	tk, err := s.issuer.IssueToken(req)
	if err != nil {
		s.checker.UpdateFailedRetry(req)
		return nil, err
	}
	tk.WithRemoteIP(req.GetRemoteIp())
	tk.WithUerAgent(req.GetUserAgent())

	// 安全登录检测
	if err := s.securityCheck(req.VerifyCode, tk); err != nil {
		return nil, err
	}

	// 登录会话
	sess, err := s.session.Login(tk)
	if err != nil {
		return nil, err
	}
	tk.SessionId = sess.ID

	// 保存入库
	if err := s.saveToken(tk); err != nil {
		return nil, err
	}

	return tk, nil
}

func (s *service) loginBeforeCheck(req *token.IssueTokenRequest) error {
	// 连续登录失败检测
	if err := s.checker.MaxFailedRetryCheck(req); err != nil {
		return exception.NewBadRequest("max retry error, %s", err)
	}

	// IP保护检测
	err := s.checker.IPProtectCheck(req)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) securityCheck(code string, tk *token.Token) error {
	// 如果有校验码, 则直接通过校验码检测用户身份安全
	if code != "" {
		s.log.Debugf("verify code provided, check code ...")
		err := s.code.CheckCode(verifycode.NewCheckCodeRequest(tk.Account, code))
		if err != nil {
			return exception.NewPermissionDeny("verify code invalidate, error, %s", err)
		}
		s.log.Debugf("verfiy code check passed")
		return nil
	}

	// 异地登录检测
	err := s.checker.OtherPlaceLoggedInChecK(tk)
	if err != nil {
		return exception.NewVerifyCodeRequiredError("异常检测: %s", err)
	}

	// 长时间未登录检测
	err = s.checker.NotLoginDaysChecK(tk)
	if err != nil {
		return exception.NewVerifyCodeRequiredError("异常检测: %s", err)
	}

	return nil
}

func (s *service) ValidateToken(ctx context.Context, req *token.ValidateTokenRequest) (*token.Token, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	tk, err := s.describeToken(newDescribeTokenRequest(req.MakeDescribeTokenRequest()))
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
	case token.BlockType_OTHER_CLIENT_LOGGED_IN:
		return exception.NewOtherClientsLoggedIn(message)
	case token.BlockType_SESSION_TERMINATED:
		return exception.NewSessionTerminated(message)
	case token.BlockType_OTHER_PLACE_LOGGED_IN:
		return exception.NewOtherPlaceLoggedIn(message)
	case token.BlockType_OTHER_IP_LOGGED_IN:
		return exception.NewOtherIPLoggedIn(message)
	default:
		return exception.NewInternalServerError("unknow block type: %s, message: %s", bt, message)
	}
}

func (s *service) BlockToken(ctx context.Context, req *token.BlockTokenRequest) (*token.Token, error) {
	tk, err := s.DescribeToken(nil, token.NewDescribeTokenRequestWithAccessToken(req.AccessToken))
	if err != nil {
		return nil, fmt.Errorf("query session access token error, %s", err)
	}

	tk.IsBlock = true
	tk.BlockType = req.BlockType
	tk.BlockReason = req.BlockReason
	tk.BlockAt = time.Now().UnixNano() / 1000000

	if err := s.updateToken(tk); err != nil {
		return nil, err
	}
	return tk, nil
}

func (s *service) DescribeToken(ctx context.Context, req *token.DescribeTokenRequest) (*token.Token, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	tk, err := s.describeToken(newDescribeTokenRequest(req))
	if err != nil {
		return nil, exception.NewUnauthorized(err.Error())
	}

	return tk, nil
}

func (s *service) QueryToken(ctx context.Context, req *token.QueryTokenRequest) (*token.Set, error) {
	query := newQueryRequest(req)
	resp, err := s.col.Find(context.TODO(), query.FindFilter(), query.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find token error, error is %s", err)
	}

	tokenSet := token.NewTokenSet()
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

func (s *service) RevolkToken(ctx context.Context, req *token.RevolkTokenRequest) (*token.Token, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	// 检测撤销token的客户端是否合法
	app, err := s.issuer.CheckClient(req.ClientId, req.ClientSecret)
	if err != nil {
		return nil, exception.NewUnauthorized(err.Error())
	}

	// 检测被撤销token的合法性
	descReq := newDescribeTokenRequest(req.MakeDescribeTokenRequest())
	tk, err := s.describeToken(descReq)
	if err != nil {
		return nil, err
	}

	if err := tk.CheckTokenApplication(app.ID); err != nil {
		return nil, exception.NewPermissionDeny(err.Error())
	}

	// 退出会话
	if req.LogoutSession {
		logoutReq := session.NewLogoutRequest(tk.SessionId)
		if err := s.session.Logout(logoutReq); err != nil {
			return nil, exception.NewInternalServerError("logout session error, %s", err)
		}
	}

	return tk, s.destoryToken(descReq)
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
