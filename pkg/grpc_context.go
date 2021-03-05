package pkg

import (
	"context"
	"fmt"
	"net/http"

	"github.com/infraboard/mcube/http/request"
	"google.golang.org/grpc/metadata"

	"github.com/infraboard/keyauth/pkg/token"
)

// NewGrpcCtx todo
func NewGrpcCtx() *GrpcCtx {
	return &GrpcCtx{md: metadata.Pairs()}
}

// GetGrpcCtx todo
func GetGrpcCtx(ctx context.Context) (*GrpcCtx, error) {
	// 获取认证信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("ctx is not an grpc incoming context")
	}

	return &GrpcCtx{md: md}, nil
}

// GetTokenFromGrpcCtx todo
func GetTokenFromGrpcCtx(ctx context.Context) (*token.Token, error) {
	rctx, err := GetGrpcCtx(ctx)
	if err != nil {
		return nil, err
	}

	return rctx.GetToken()
}

// GrpcCtx todo
type GrpcCtx struct {
	md metadata.MD
}

// Get todo
func (c *GrpcCtx) get(key string) string {
	if val, ok := c.md[key]; ok {
		return val[0]
	}

	return ""
}

func (c *GrpcCtx) set(key, value string) {
	c.md.Set(key, value)
}

// Context todo
func (c *GrpcCtx) Context() context.Context {
	return metadata.NewOutgoingContext(context.Background(), c.md)
}

// GetAccessToKen todo
func (c *GrpcCtx) GetAccessToKen() string {
	return c.get("x-oauth-token")
}

// GetToken todo
func (c *GrpcCtx) GetToken() (*token.Token, error) {
	ctx := context.Background()
	req := token.NewDescribeTokenRequestWithAccessToken(c.GetAccessToKen())
	return Token.DescribeToken(ctx, req)
}

// GetClientID todo
func (c *GrpcCtx) GetClientID() string {
	return c.get("client-id")
}

// GetClientSecret todo
func (c *GrpcCtx) GetClientSecret() string {
	return c.get("client-secret")
}

// SetAccessToken todo
func (c *GrpcCtx) SetAccessToken(ak string) {
	c.set("x-oauth-token", ak)
}

// SetRemoteIP todo
func (c *GrpcCtx) SetRemoteIP(ip string) {
	c.set("x-real-ip", ip)
}

// SetUserAgent todo
func (c *GrpcCtx) SetUserAgent(ua string) {
	c.set("user-agent", ua)
}

// GetGrpcCtxFromHTTPRequest 从上下文中获取Token
func GetGrpcCtxFromHTTPRequest(r *http.Request) (*GrpcCtx, error) {
	rc := NewGrpcCtx()
	rc.SetAccessToken(r.Header.Get("X-OAUTH-TOKEN"))
	rc.SetRemoteIP(request.GetRemoteIP(r))
	rc.SetUserAgent(r.UserAgent())
	return rc, nil
}
