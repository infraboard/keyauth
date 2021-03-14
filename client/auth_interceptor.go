package client

import (
	"context"
	"strconv"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/grpc/gcontext"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	httpb "github.com/infraboard/mcube/pb/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/permission"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
	"github.com/infraboard/keyauth/version"
)

type PathEntryHandleFunc func(path string) *httpb.Entry

// NewGrpcKeyauthAuther todo
func NewGrpcKeyauthAuther(hf PathEntryHandleFunc, c *Client) *GrpcAuther {
	return &GrpcAuther{
		hf: hf,
		c:  c,
	}
}

// GrpcAuther todo
type GrpcAuther struct {
	hf PathEntryHandleFunc
	l  logger.Logger
	c  *Client
}

// AuthUnaryServerInterceptor returns a new unary server interceptor for auth.
func (a *GrpcAuther) AuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return a.auth
}

// SetLogger todo
func (a *GrpcAuther) SetLogger(l logger.Logger) {
	a.l = l
}

// Auth impl interface
func (a *GrpcAuther) auth(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	// 重上下文中获取认证信息
	rctx, err := gcontext.GetGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 校验调用的客户端凭证是否有效
	if err := a.validateServiceCredential(rctx); err != nil {
		return nil, err
	}

	// 校验用户权限是否合法
	if err := a.validatePermission(rctx, info.FullMethod); err != nil {
		return nil, err
	}

	switch t := err.(type) {
	case exception.APIException:
		err = status.Errorf(codes.Code(t.ErrorCode()), t.Error())
		trailer := metadata.Pairs(
			gcontext.ResponseCodeHeader, strconv.Itoa(t.ErrorCode()),
			gcontext.ResponseReasonHeader, t.Reason(),
			gcontext.ResponseDescHeader, t.Error(),
		)
		if err := grpc.SetTrailer(ctx, trailer); err != nil {
			a.log().Errorf("send grpc trailer error, %s", err)
		}
	}

	return resp, err
}

func (a *GrpcAuther) validateServiceCredential(ctx *gcontext.GrpcInCtx) error {
	clientID := ctx.GetClientID()
	clientSecret := ctx.GetClientSecret()

	if clientID == "" && clientSecret == "" {
		return grpc.Errorf(codes.Unauthenticated, "client_id or client_secret is \"\"")
	}

	vsReq := micro.NewValidateClientCredentialRequest(clientID, clientSecret)
	_, err := a.c.Micro().ValidateClientCredential(context.Background(), vsReq)
	if err != nil {
		return grpc.Errorf(codes.Unauthenticated, "service auth error, %s", err)
	}

	return nil
}

func (a *GrpcAuther) validatePermission(ctx *gcontext.GrpcInCtx, path string) error {
	var (
		tk  *token.Token
		err error
	)

	entry := a.hf(path)
	if entry == nil {
		return grpc.Errorf(codes.Internal, "entry not found, check is registry")
	}

	if entry.AuthEnable {
		// 获取需要校验的access token(用户的身份凭证)
		accessToken := ctx.GetAccessToKen()
		req := token.NewValidateTokenRequest()
		if accessToken == "" {
			return grpc.Errorf(codes.Unauthenticated, "access_token meta required")
		}
		req.AccessToken = accessToken

		tk, err = a.c.Token().ValidateToken(context.Background(), req)
		if err != nil {
			return err
		}
	}

	if entry.PermissionEnable && tk != nil {
		// 如果是超级管理员不做权限校验, 直接放行
		if tk.UserType.IsIn(types.UserType_SUPPER) {
			return nil
		}

		req := permission.NewCheckPermissionrequest()
		req.EndpointId = a.endpointHashID(entry)
		req.NamespaceId = ctx.GetNamespace()
		_, err = a.c.Permission().CheckPermission(gcontext.NewGrpcOutCtxFromIn(ctx).Context(), req)
		if err != nil {
			return exception.NewPermissionDeny("no permission, %s", err)
		}
	}

	return nil
}

func (a *GrpcAuther) endpointHashID(entry *httpb.Entry) string {
	return endpoint.GenHashID(version.ServiceName, entry.GrpcPath)
}

func (a *GrpcAuther) log() logger.Logger {
	if a == nil {
		a.l = zap.L().Named("GRPC Auther")
	}

	return a.l
}
