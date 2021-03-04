package pkg

import (
	"context"
	"fmt"

	"github.com/infraboard/keyauth/common/session"
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/permission"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
	"github.com/infraboard/keyauth/version"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pb/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var (
	interceptor = &grpcAuther{}
)

// AuthUnaryServerInterceptor returns a new unary server interceptor for auth.
func AuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return interceptor.Auth
}

// internal todo
type grpcAuther struct{}

func (a *grpcAuther) Auth(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	// 获取认证信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing credentials")
	}

	// 校验调用的客户端凭证是否有效
	if err := a.validateServiceCredential(md); err != nil {
		return nil, err
	}

	// 校验用户权限是否合法
	if err := a.validatePermission(info.FullMethod, md); err != nil {
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	return handler(ctx, req)
}

func (a *grpcAuther) validateServiceCredential(md metadata.MD) error {
	var clientID, clientSecret string
	if val, ok := md["client_id"]; ok {
		clientID = val[0]
	}
	if val, ok := md["client_secret"]; ok {
		clientSecret = val[0]
	}

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

func (a *grpcAuther) validatePermission(path string, md metadata.MD) error {
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
		var accessToken string
		if val, ok := md["access_token"]; ok {
			accessToken = val[0]
		}

		req := token.NewValidateTokenRequest()
		if accessToken == "" {
			return grpc.Errorf(codes.Unauthenticated, "access_token meta required")
		}
		req.AccessToken = accessToken

		tk, err = Token.ValidateToken(nil, req)
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
		if Permission == nil {
			return fmt.Errorf("permission service not load")
		}

		req := permission.NewCheckPermissionrequest()
		req.EndpointId = a.endpointHashID(entry)
		ctx := session.WithTokenContext(context.Background(), tk)
		_, err = Permission.CheckPermission(ctx, req)
		if err != nil {
			return exception.NewPermissionDeny("no permission")
		}
	}

	return nil
}

func (a *grpcAuther) endpointHashID(entry *http.Entry) string {
	return endpoint.GenHashID(version.ServiceName, entry.GrpcPath)
}
