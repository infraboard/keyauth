package pkg

import (
	"context"
	"fmt"
	"log"

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
	var clientID, clientSecret string
	if val, ok := md["client_id"]; ok {
		clientID = val[0]
	}
	if val, ok := md["client_secret"]; ok {
		clientSecret = val[0]
	}
	if clientID == "" && clientSecret == "" {
		return nil, grpc.Errorf(codes.Unauthenticated, "client_id or client_secret is \"\"")
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	return handler(ctx, req)
}
