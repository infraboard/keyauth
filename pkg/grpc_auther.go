package pkg

import (
	"context"
	"strconv"
	"strings"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pb/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
	"github.com/infraboard/keyauth/version"
)

var (
	interceptor = newGrpcAuther()
)

// 检测是不是owner请求
type OwnerChecker interface {
	CheckOwner(account string) bool
}

// AuthUnaryServerInterceptor returns a new unary server interceptor for auth.
func AuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return interceptor.Auth
}

func newGrpcAuther() *grpcAuther {
	return &grpcAuther{}
}

// internal todo
type grpcAuther struct {
	l logger.Logger
}

func (a *grpcAuther) Auth(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	// 重上下文中获取认证信息
	rctx, err := GetGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 校验调用的客户端凭证是否有效
	if err := a.validateServiceCredential(rctx); err != nil {
		return nil, err
	}

	entry := GetGrpcPathEntry(info.FullMethod)
	if entry == nil {
		return nil, grpc.Errorf(codes.Internal, "entry gprc path: %s not found, check is registry", info.FullMethod)
	}

	if entry.AuthEnable {
		tk, err := a.checkToken(rctx)
		if err != nil {
			return nil, err
		}

		// 权限校验
		if err := a.validatePermission(tk, entry, resp); err != nil {
			return nil, err
		}
	}

	rctx.ClearInternl()
	resp, err = handler(rctx.ClearInternl().Context(), req)

	switch t := err.(type) {
	case exception.APIException:
		err = status.Errorf(codes.Code(t.ErrorCode()), t.Error())
		// create and set trailer
		trailer := metadata.Pairs(
			ResponseCodeHeader, strconv.Itoa(t.ErrorCode()),
			ResponseReasonHeader, t.Reason(),
			ResponseDescHeader, t.Error(),
		)
		if err := grpc.SetTrailer(ctx, trailer); err != nil {
			a.log().Errorf("send grpc trailer error, %s", err)
		}
	}
	return resp, err
}

func (a *grpcAuther) validateServiceCredential(ctx *GrpcInCtx) error {
	clientID := ctx.GetClientID()
	clientSecret := ctx.GetClientSecret()

	if clientID == "" && clientSecret == "" {
		return grpc.Errorf(codes.Unauthenticated, "client_id or client_secret is \"\"")
	}

	if Micro == nil {
		return grpc.Errorf(codes.Internal, "micro service is not initial")
	}

	vsReq := micro.NewValidateClientCredentialRequest(clientID, clientSecret)
	_, err := Micro.ValidateClientCredential(context.Background(), vsReq)
	if err != nil {
		return grpc.Errorf(codes.Unauthenticated, "service auth error, %s", err)
	}

	return nil
}

// 获取需要校验的access token(用户的身份凭证)
func (a *grpcAuther) checkToken(ctx *GrpcInCtx) (*token.Token, error) {
	accessToken := ctx.GetAccessToKen()
	req := token.NewValidateTokenRequest()
	if accessToken == "" {
		return nil, grpc.Errorf(codes.Unauthenticated, "access_token meta required")
	}
	req.AccessToken = accessToken
	return Token.ValidateToken(context.Background(), req)
}

func (a *grpcAuther) validatePermission(tk *token.Token, entry *http.Entry, req interface{}) error {
	// 如果是超级管理员不做权限校验, 直接放行
	if tk.UserType.IsIn(types.UserType_SUPPER) {
		return nil
	}

	// 检测owner
	if v, ok := req.(OwnerChecker); ok {
		if !v.CheckOwner(tk.Account) {
			return grpc.Errorf(codes.PermissionDenied, "only owner can operate")
		}
	}

	if v, ok := entry.Labels["allow"]; ok {
		types := strings.Split(v, ",")
		a.log().Debugf("allows: %v", types)
	}

	return nil
}

func (a *grpcAuther) endpointHashID(entry *http.Entry) string {
	return endpoint.GenHashID(version.ServiceName, entry.GrpcPath)
}

func (a *grpcAuther) log() logger.Logger {
	if a == nil {
		a.l = zap.L().Named("GRPC Auther")
	}

	return a.l
}
