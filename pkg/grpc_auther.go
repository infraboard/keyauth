package pkg

import (
	"context"
	"fmt"
	"strconv"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/permission"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
	"github.com/infraboard/keyauth/version"
	"github.com/infraboard/mcube/exception"
	httpb "github.com/infraboard/mcube/pb/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	interceptor = newGrpcAuther()
)

// AuthUnaryServerInterceptor returns a new unary server interceptor for auth.
func AuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return interceptor.Auth
}

func newGrpcAuther() *grpcAuther {
	return &grpcAuther{}
}

// internal todo
type grpcAuther struct {
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

	// 校验用户权限是否合法
	if err := a.validatePermission(rctx, info.FullMethod); err != nil {
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	rctx.ClearInternl()
	resp, err = handler(rctx.ClearInternl().Context(), req)

	//
	switch t := err.(type) {
	case exception.APIException:
		err = status.Errorf(codes.Code(t.ErrorCode()), t.Error())
		// create and set trailer
		trailer := metadata.Pairs(
			ResponseCodeHeader, strconv.Itoa(t.ErrorCode()),
			ResponseReasonHeader, t.Reason(),
		)
		grpc.SetTrailer(ctx, trailer)
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

func (a *grpcAuther) validatePermission(ctx *GrpcInCtx, path string) error {
	var (
		tk  *token.Token
		err error
	)

	entry := GetGrpcPathEntry(path)
	if entry == nil {
		grpc.Errorf(codes.Internal, "entry not nod, check is registry")
	}

	if entry.AuthEnable {
		// 获取需要校验的access token(用户的身份凭证)
		accessToken := ctx.GetAccessToKen()
		req := token.NewValidateTokenRequest()
		if accessToken == "" {
			return grpc.Errorf(codes.Unauthenticated, "access_token meta required")
		}
		req.AccessToken = accessToken

		tk, err = Token.ValidateToken(context.Background(), req)
		if err != nil {
			return err
		}
	}

	if entry.PermissionEnable && tk != nil {
		// 如果是超级管理员不做权限校验, 直接放行
		if tk.UserType.IsIn(types.UserType_SUPPER) {
			return nil
		}

		// 其他比如服务类型, 主账号类型, 子账号类型
		// 如果开启权限认证都需要检查
		client := client.C()
		if client == nil {
			return fmt.Errorf("grpc client service not initial")
		}

		req := permission.NewCheckPermissionrequest()
		req.EndpointId = a.endpointHashID(entry)

		_, err = client.Permission().CheckPermission(ctx.Context(), req)
		if err != nil {
			return exception.NewPermissionDeny("no permission")
		}
	}

	return nil
}

func (a *grpcAuther) endpointHashID(entry *httpb.Entry) string {
	return endpoint.GenHashID(version.ServiceName, entry.GrpcPath)
}
