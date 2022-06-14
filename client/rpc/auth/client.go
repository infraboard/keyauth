package auth

import (
	"context"

	"google.golang.org/grpc/metadata"
)

// 客户端携带的凭证
func NewAuthentication(clientId, clientSecret string) *Authentication {
	return &Authentication{
		clientID:     clientId,
		clientSecret: clientSecret,
	}
}

// Authentication todo
type Authentication struct {
	clientID     string
	clientSecret string
}

// SetClientCredentials todo
func (a *Authentication) SetClientCredentials(clientID, clientSecret string) {
	a.clientID = clientID
	a.clientSecret = clientSecret
}

// GetRequestMetadata todo
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (
	map[string]string, error,
) {
	return map[string]string{
		ClientHeaderKey: a.clientID,
		ClientSecretKey: a.clientSecret,
	}, nil
}

// RequireTransportSecurity todo
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

func GetClientCredential(ctx context.Context) (clientId, clientSecret string) {
	// 重上下文中获取认证信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return
	}

	cids := md.Get(ClientHeaderKey)
	sids := md.Get(ClientSecretKey)
	if len(cids) > 0 {
		clientId = cids[0]
	}
	if len(sids) > 0 {
		clientSecret = sids[0]
	}

	return
}
