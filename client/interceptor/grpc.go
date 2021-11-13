package interceptor

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/infraboard/keyauth/app/micro"
	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/common/header"
)

// GrpcAuthUnaryServerInterceptor returns a new unary server interceptor for auth.
func GrpcAuthUnaryServerInterceptor(c *client.Client) grpc.UnaryServerInterceptor {
	return newGrpcAuther(c).Auth
}

func newGrpcAuther(c *client.Client) *grpcAuther {
	return &grpcAuther{
		log:     zap.L().Named("Grpc Auther"),
		keyauth: c,
	}
}

// internal todo
type grpcAuther struct {
	log     logger.Logger
	keyauth *client.Client
}

func (a *grpcAuther) Auth(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	// 重上下文中获取认证信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("ctx is not an grpc incoming context")
	}

	clientId, clientSecret := a.GetClientCredentialsFromMeta(md)

	// 校验调用的客户端凭证是否有效
	if err := a.validateServiceCredential(clientId, clientSecret); err != nil {
		return nil, err
	}

	resp, err = handler(ctx, req)
	return resp, err
}

func (a *grpcAuther) GetClientCredentialsFromMeta(md metadata.MD) (
	clientId, clientSecret string) {
	cids := md.Get(header.ClientHeaderKey)
	sids := md.Get(header.ClientSecretKey)
	if len(cids) > 0 {
		clientId = cids[0]
	}
	if len(sids) > 0 {
		clientSecret = sids[0]
	}
	return
}

func (a *grpcAuther) validateServiceCredential(clientId, clientSecret string) error {
	if clientId == "" && clientSecret == "" {
		return status.Errorf(codes.Unauthenticated, "client_id or client_secret is \"\"")
	}

	vsReq := micro.NewValidateClientCredentialRequest(clientId, clientSecret)
	_, err := a.keyauth.Micro().ValidateClientCredential(context.Background(), vsReq)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "service auth error, %s", err)
	}

	return nil
}
