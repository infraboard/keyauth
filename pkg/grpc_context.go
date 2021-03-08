package pkg

import (
	"context"
	"fmt"
	"net/http"

	"github.com/infraboard/mcube/http/request"
	"google.golang.org/grpc/metadata"

	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
)

const (
	// InternalCallTokenHeader todo
	InternalCallTokenHeader = "internal-call-token"
)

// NewGrpcCtx todo
func NewGrpcCtx() *GrpcCtx {
	return &GrpcCtx{md: metadata.Pairs()}
}

// NewInternalMockGrpcCtx todo
func NewInternalMockGrpcCtx(account string) *GrpcCtx {
	ctx := NewGrpcCtx()
	ctx.SetIsInternalCall(account, domain.AdminDomainName)
	return ctx
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

	// 是不是内部调用, 如果是内部调用直接mock token
	if rctx.IsInternalCall() {
		return rctx.InternalCallToken(), nil
	}

	return rctx.GetToken()
}

// GrpcCtx todo
type GrpcCtx struct {
	md metadata.MD
}

// Get todo
func (c *GrpcCtx) get(key string) string {
	return c.getWithIndex(key, 0)
}

// Get todo
func (c *GrpcCtx) getWithIndex(key string, index int) string {
	if val, ok := c.md[key]; ok {
		if len(val) > index {
			return val[index]
		}
	}

	return ""
}

func (c *GrpcCtx) set(key string, values ...string) {
	c.md.Set(key, values...)
}

// ClearInternl todo
func (c *GrpcCtx) ClearInternl() *GrpcCtx {
	delete(c.md, InternalCallTokenHeader)
	return c
}

// Context todo
func (c *GrpcCtx) Context() context.Context {
	return metadata.NewOutgoingContext(context.Background(), c.md)
}

// InContext todo
func (c *GrpcCtx) InContext() context.Context {
	return metadata.NewIncomingContext(context.Background(), c.md)
}

// GetAccessToKen todo
func (c *GrpcCtx) GetAccessToKen() string {
	return c.get("x-oauth-token")
}

// GetToken todo
func (c *GrpcCtx) GetToken() (*token.Token, error) {
	req := token.NewDescribeTokenRequestWithAccessToken(c.GetAccessToKen())
	ctx := NewInternalMockGrpcCtx("internal").Context()
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

// SetIsInternalCall 内部调用不需要认证
func (c *GrpcCtx) SetIsInternalCall(account, domain string) {
	c.set(InternalCallTokenHeader, account, domain)
}

// IsInternalCall todo
func (c *GrpcCtx) IsInternalCall() bool {
	if _, ok := c.md[InternalCallTokenHeader]; ok {
		return true
	}

	return false
}

// InternalCallToken 是不是内部调用
func (c *GrpcCtx) InternalCallToken() *token.Token {
	tk := &token.Token{UserType: types.UserType_INTERNAL}
	tk.Account = c.getWithIndex(InternalCallTokenHeader, 0)
	tk.Domain = c.getWithIndex(InternalCallTokenHeader, 1)
	return tk
}

// GetGrpcCtxFromHTTPRequest 从上下文中获取Token
func GetGrpcCtxFromHTTPRequest(r *http.Request) (*GrpcCtx, error) {
	rc := NewGrpcCtx()
	rc.SetAccessToken(r.Header.Get("X-OAUTH-TOKEN"))
	rc.SetRemoteIP(request.GetRemoteIP(r))
	rc.SetUserAgent(r.UserAgent())
	return rc, nil
}
