package session

import (
	"github.com/infraboard/keyauth/pkg/token"
)

// NewSession todo
func NewSession() *Session {
	return &Session{}
}

// Session 请求上下文信息
type Session struct {
	tk *token.Token
}

// WithToken 携带token
func (s *Session) WithToken(tk *token.Token) {
	s.tk = tk
}

// WithTokenGetter geter
func (s *Session) WithTokenGetter(gt Getter) {
	s.tk = gt.GetToken()
}

// GetToken 获取token
func (s *Session) GetToken() *token.Token {
	return s.tk
}

// GetAccount todo
func (s *Session) GetAccount() string {
	if s.tk == nil {
		return "Nil"
	}

	return s.tk.Account
}

// Getter 获取token
type Getter interface {
	GetToken() *token.Token
}
