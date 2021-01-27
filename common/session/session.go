package session

import (
	"context"
	"net/http"

	"github.com/infraboard/mcube/exception"
	httpcontext "github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"

	"github.com/infraboard/keyauth/pkg/token"
)

// GetTokenFromHTTPRequest 从上下文中获取Token
func GetTokenFromHTTPRequest(r *http.Request) (*token.Token, error) {
	ctx := httpcontext.GetContext(r)

	if ctx.AuthInfo == nil {
		return nil, exception.NewInternalServerError("authInfo is not in request context, please check auth middleware")
	}

	tk, ok := httpcontext.GetContext(r).AuthInfo.(*token.Token)
	if !ok {
		return nil, exception.NewInternalServerError("authInfo is not token pointer")
	}

	tk.WithRemoteIP(request.GetRemoteIP(r))
	tk.WithUerAgent(r.UserAgent())
	return tk, nil
}

// GetTokenCtxFromHTTPRequest todo
func GetTokenCtxFromHTTPRequest(r *http.Request) (context.Context, error) {
	tk, err := GetTokenFromHTTPRequest(r)
	if err != nil {
		return nil, err
	}

	return WithTokenContext(context.Background(), tk), nil
}

// ContextKeyType key类型
type contextKeyType string

const (
	// ContextKeyType_TOKEN token 上下文key
	contextKeyTypeToken = contextKeyType("token")
)

// WithTokenContext todo
func WithTokenContext(ctx context.Context, tk *token.Token) context.Context {
	return context.WithValue(ctx, contextKeyTypeToken, tk)
}

// GetTokenFromContext 从上下文中获取token
func GetTokenFromContext(ctx context.Context) *token.Token {
	if v, ok := ctx.Value(contextKeyTypeToken).(*token.Token); ok {
		return v
	}

	return token.NewDefaultToken()
}
