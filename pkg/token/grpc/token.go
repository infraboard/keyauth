package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/namespace"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/session"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
	"github.com/infraboard/keyauth/pkg/verifycode"
)

func (s *service) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	// 连续登录失败检测
	if err := s.loginBeforeCheck(ctx, req); err != nil {
		return nil, exception.NewBadRequest("安全检测失败, %s", err)
	}

	// 颁发Token
	tk, err := s.issuer.IssueToken(ctx, req)
	if err != nil {
		s.checker.UpdateFailedRetry(ctx, req)
		return nil, err
	}
	tk.WithRemoteIP(req.GetRemoteIp())
	tk.WithUerAgent(req.GetUserAgent())

	// 安全登录检测
	if err := s.securityCheck(ctx, req.VerifyCode, tk); err != nil {
		return nil, err
	}

	// 登录会话
	if req.IsLoginRequest() {
		sess, err := s.session.Login(ctx, tk)
		if err != nil {
			return nil, err
		}
		tk.SessionId = sess.Id
	}

	// 保存入库
	if err := s.saveToken(tk); err != nil {
		return nil, err
	}

	return tk, nil
}

func (s *service) loginBeforeCheck(ctx context.Context, req *token.IssueTokenRequest) error {
	// 连续登录失败检测
	if err := s.checker.MaxFailedRetryCheck(ctx, req); err != nil {
		return exception.NewBadRequest("%s", err)
	}

	// IP保护检测
	err := s.checker.IPProtectCheck(ctx, req)
	if err != nil {
		return err
	}

	s.log.Debug("security check complete")
	return nil
}

func (s *service) securityCheck(ctx context.Context, code string, tk *token.Token) error {
	// 如果有校验码, 则直接通过校验码检测用户身份安全
	if code != "" {
		s.log.Debugf("verify code provided, check code ...")
		_, err := s.code.CheckCode(ctx, verifycode.NewCheckCodeRequest(tk.Account, code))
		if err != nil {
			return exception.NewPermissionDeny("verify code invalidate, error, %s", err)
		}
		s.log.Debugf("verfiy code check passed")
		return nil
	}

	// 异地登录检测
	err := s.checker.OtherPlaceLoggedInChecK(ctx, tk)
	if err != nil {
		return exception.NewVerifyCodeRequiredError("异地登录检测失败: %s", err)
	}

	// 长时间未登录检测
	err = s.checker.NotLoginDaysChecK(ctx, tk)
	if err != nil {
		return exception.NewVerifyCodeRequiredError("长时间未登录检测失败: %s", err)
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

	// 校验Access Token是否过期
	if req.AccessToken != "" {
		if tk.CheckAccessIsExpired() {
			return nil, exception.NewAccessTokenExpired("access_token: %s has expired", tk.AccessToken)
		}
	}

	// 校验RefreshToken
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
	tk, err := s.DescribeToken(ctx, token.NewDescribeTokenRequestWithAccessToken(req.AccessToken))
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

func (s *service) ChangeNamespace(ctx context.Context, req *token.ChangeNamespaceRequest) (*token.Token, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate change namespace error, %s", err)
	}

	tk, err := s.DescribeToken(ctx, token.NewDescribeTokenRequestWithAccessToken(req.Token))
	if err != nil {
		return nil, err
	}

	_, err = s.ns.DescribeNamespace(ctx, namespace.NewNewDescriptNamespaceRequestWithID(req.Namespace))
	if err != nil {
		return nil, err
	}

	fmt.Println(tk.UserType, types.UserType_DOMAIN_ADMIN, types.UserType_SUPPER, tk.Namespace)
	if !tk.UserType.IsIn(types.UserType_DOMAIN_ADMIN, types.UserType_SUPPER) && !tk.HasNamespace(req.Namespace) {
		return nil, exception.NewPermissionDeny("your has no permission to access namespace %s", req.Namespace)
	}

	tk.Namespace = req.Namespace
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

	// 查询用户可以访问的空间
	query := policy.NewQueryPolicyRequest(request.NewPageRequest(policy.MaxUserPolicy, 1))
	query.Account = tk.Account
	ps, err := s.policy.QueryPolicy(pkg.NewInternalMockGrpcCtx(tk.Account).Context(), query)
	if err != nil {
		return nil, err
	}
	if ps.Total > policy.MaxUserPolicy {
		s.log.Warnf("user policy large than max policy count %d, total: %d", policy.MaxUserPolicy, ps.Total)
	}
	tk.AvailableNamespace = ps.GetNamespace()

	return tk, nil
}

func (s *service) QueryToken(ctx context.Context, req *token.QueryTokenRequest) (*token.Set, error) {
	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}
	req.Account = tk.Account

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
	app, err := s.issuer.CheckClient(ctx, req.ClientId, req.ClientSecret)
	if err != nil {
		return nil, exception.NewUnauthorized(err.Error())
	}

	// 检测被撤销token的合法性
	descReq := newDescribeTokenRequest(req.MakeDescribeTokenRequest())
	tk, err := s.describeToken(descReq)
	if err != nil {
		return nil, err
	}

	if err := tk.CheckTokenApplication(app.Id); err != nil {
		return nil, exception.NewPermissionDeny(err.Error())
	}

	// 退出会话
	if req.LogoutSession && tk.SessionId != "" {
		logoutReq := session.NewLogoutRequest(tk.SessionId)
		if _, err := s.session.Logout(ctx, logoutReq); err != nil {
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

func (s *service) DeleteToken(ctx context.Context, req *token.DeleteTokenRequest) (
	*token.DeleteTokenResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	deleteReq := newDeleteTokenRequest(req)
	s.log.Debugf("delete token filter: %s", deleteReq.FindFilter())
	resp, err := s.col.DeleteOne(context.TODO(), deleteReq.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("delete token(%s) error, %s", req, err)
	}

	if resp.DeletedCount == 0 {
		return nil, exception.NewNotFound("token %s not found", req.AccessToken)
	}

	dr := token.NewDeleteTokenResponse()
	dr.Message = fmt.Sprintf("delete %d token", resp.DeletedCount)
	return dr, nil
}
