package pkg

import (
	"github.com/infraboard/keyauth/pkg/token"
)

// NewContext new上下文
func NewContext() *Context {
	return new(Context)
}

// Context 请求上下文信息
type Context struct {
	tk *token.Token
}

// WithToken 携带token
func (c *Context) WithToken(tk *token.Token) {
	c.tk = tk
}

// GetToken 获取token
func (c *Context) GetToken() *token.Token {
	return c.tk
}
