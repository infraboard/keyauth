package session

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/exception"
	"github.com/mssola/user_agent"
	"github.com/rs/xid"

	"github.com/infraboard/keyauth/pkg/ip2region"
	"github.com/infraboard/keyauth/pkg/token"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewSession todo
func NewSession(ip ip2region.Service, tk *token.Token) (*Session, error) {
	if err := tk.IsAvailable(); err != nil {
		return nil, exception.NewPermissionDeny("token is not available, %s", err)
	}

	sess := &Session{
		Id:              xid.New().String(),
		Domain:          tk.Domain,
		Account:         tk.Account,
		UserType:        tk.UserType,
		ApplicationId:   tk.ApplicationId,
		ApplicationName: tk.ApplicationName,
		GrantType:       tk.GrantType,
		AccessToken:     tk.AccessToken,
		LoginAt:         tk.CreateAt,
		LoginIp:         tk.GetRemoteIP(),
		UserAgent:       &UserAgent{},
		IpInfo:          &IPInfo{},
	}
	sess.ParseUserAgent(tk.GetUserAgent())

	return sess, nil
}

// NewDefaultSession todo
func NewDefaultSession() *Session {
	return &Session{
		UserAgent: &UserAgent{},
		IpInfo:    &IPInfo{},
	}
}

// ParseUserAgent todo
func (s *Session) ParseUserAgent(userAgent string) {
	if userAgent == "" {
		return
	}

	ua := user_agent.New(userAgent)
	s.UserAgent = &UserAgent{
		Os:       ua.OS(),
		Platform: ua.Platform(),
	}
	s.UserAgent.EngineName, s.UserAgent.EngineVersion = ua.Engine()
	s.UserAgent.BrowserName, s.UserAgent.BrowserVersion = ua.Browser()
}

// NewSessionSet 实例化
func NewSessionSet() *Set {
	return &Set{
		Items: []*Session{},
	}
}

// Add 添加应用
func (s *Set) Add(item *Session) {
	s.Items = append(s.Items, item)
}

// Length 长度
func (s *Set) Length() int {
	return len(s.Items)
}

// IsEmpty 长度
func (s *Set) IsEmpty() bool {
	return len(s.Items) == 0
}
