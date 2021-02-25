package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/infraboard/keyauth/pkg/micro"
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
	log.Println("filter:", info)

	// 获取认证信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing credentials")
	}

	// 校验调用的客户端凭证是否有效
	if err := a.validateServiceCredential(md); err != nil {
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
		return grpc.Errorf(codes.Internal, "micro service is nil")
	}

	vsReq := micro.NewValidateClientCredentialRequest(clientID, clientSecret)
	_, err := Micro.ValidateClientCredential(context.Background(), vsReq)
	if err != nil {
		return grpc.Errorf(codes.Unauthenticated, "service auth error, %s", err)
	}

	return nil
}
