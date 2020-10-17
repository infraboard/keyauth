package session

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/mssola/user_agent"
	"github.com/rs/xid"

	"github.com/infraboard/keyauth/pkg/ip2region"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
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
		ID:              xid.New().String(),
		Domain:          tk.Domain,
		Account:         tk.Account,
		AccountType:     tk.UserType,
		ApplicationID:   tk.ApplicationID,
		ApplicationName: tk.ApplicationName,
		GrantType:       tk.GrantType,
		AccessToken:     tk.AccessToken,
		LoginAt:         tk.CreatedAt,
		LoginIP:         tk.GetRemoteIP(),
		log:             zap.L().Named("Session"),
		ip:              ip,
	}
	sess.ParseUserAgent(tk.GetUserAgent())
	sess.ParseLoginAddress(tk.GetRemoteIP())

	return sess, nil
}

// NewDefaultSession todo
func NewDefaultSession() *Session {
	return &Session{
		IPInfo: ip2region.NewDefaultIPInfo(),
	}
}

// Session 登录回话
type Session struct {
	ID              string          `bson:"_id" json:"id"`                                                // sessionID
	Domain          string          `bson:"domain" json:"domain" alidate:"required"`                      // 所处域
	Account         string          `bson:"account" json:"account" validate:"required"`                   // 用户名称
	AccountType     types.Type      `bson:"account_type" json:"account_type" validate:"required"`         // 用户类型
	ApplicationID   string          `bson:"application_id" json:"application_id" validate:"required"`     // 用户通过哪个端登录的
	ApplicationName string          `bson:"application_name" json:"application_name" validate:"required"` // 用户通过哪个端登录的
	GrantType       token.GrantType `bson:"grant_type" json:"grant_type" validate:"required"`             // 登录方式
	LoginAt         ftime.Time      `bson:"login_at" json:"login_at" validate:"required"`                 // 登录时间
	LoginIP         string          `bson:"login_ip" json:"login_ip" validate:"required"`                 // 登录IP
	LogoutAt        ftime.Time      `bson:"logout_at" json:"logout_at"`                                   // 登出时间
	AccessToken     string          `bson:"access_token" json:"access_token"`                             // 当前会话的访问的token

	UserAgent         `bson:",inline"` // 登录端信息
	*ip2region.IPInfo `bson:",inline"` // 登录地

	ip  ip2region.Service // 地址查询服务
	log logger.Logger     //日志服务
}

// ParseLoginAddress todo
func (s *Session) ParseLoginAddress(ip string) {
	if ip == "" {
		return
	}

	ipInfo, err := s.ip.LookupIP(ip)
	if err != nil {
		s.log.Errorf("parse ipinfo error, %s", err)
	}
	s.IPInfo = ipInfo
	return
}

// ParseUserAgent todo
func (s *Session) ParseUserAgent(userAgent string) {
	if userAgent == "" {
		return
	}

	ua := user_agent.New(userAgent)
	s.UserAgent = UserAgent{
		OS:       ua.OS(),
		Platform: ua.Platform(),
	}
	s.EngineName, s.EngineVersion = ua.Engine()
	s.BrowserName, s.BrowserVersion = ua.Browser()
}

// NewSessionSet 实例化
func NewSessionSet(req *request.PageRequest) *Set {
	return &Set{
		PageRequest: req,
		Items:       []*Session{},
	}
}

// Set todo
type Set struct {
	*request.PageRequest

	Total int64      `json:"total"`
	Items []*Session `json:"items"`
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

// UserAgent todo
type UserAgent struct {
	OS             string `bson:"os" json:"os"`
	Platform       string `bson:"platform" json:"platform"`
	EngineName     string `bson:"engine_name" json:"engine_name"`
	EngineVersion  string `bson:"engine_version" json:"engine_version"`
	BrowserName    string `bson:"browser_name" json:"browser_name"`
	BrowserVersion string `bson:"browser_version" json:"browser_version"`
}
