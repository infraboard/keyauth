package audit

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/mssola/user_agent"
	"github.com/rs/xid"

	"github.com/infraboard/keyauth/pkg/ip2region"
	"github.com/infraboard/keyauth/pkg/token"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewLoginLog todo
func NewLoginLog(data *LoginLogData) *LoginLog {
	log := &LoginLog{
		ID:           xid.New().String(),
		Domain:       data.GetToken().Domain,
		LoginLogData: data,
	}

	return log
}

// NewDefaultLoginLog todo
func NewDefaultLoginLog() *LoginLog {
	return &LoginLog{
		LoginLogData: NewDefaultLoginLogData(),
		IPInfo:       ip2region.NewDefaultIPInfo(),
	}
}

// LoginLog 登录日志
type LoginLog struct {
	ID                string `bson:"_id" json:"id"`
	Domain            string `bson:"domain" json:"domain" alidate:"required"` // 所处域
	*LoginLogData     `bson:",inline"`
	UserAgent         `bson:",inline"`
	*ip2region.IPInfo `bson:",inline"`
}

// ParseLoginIP todo
func (l *LoginLog) ParseLoginIP(ip ip2region.Service) error {
	if l.LoginIP == "" {
		return nil
	}

	ipInfo, err := ip.LookupIP(l.LoginIP)
	if err != nil {
		return fmt.Errorf("parse ipinfo error, %s", err)
	}
	l.IPInfo = ipInfo
	return nil
}

// ParseUserAgent todo
func (l *LoginLog) ParseUserAgent() {
	if l.userAgent == "" {
		return
	}

	ua := user_agent.New(l.userAgent)
	l.UserAgent = UserAgent{
		OS:       ua.OS(),
		Platform: ua.Platform(),
	}
	l.EngineName, l.EngineVersion = ua.Engine()
	l.BrowserName, l.BrowserVersion = ua.Browser()
}

// NewLoginRecordSet 实例化
func NewLoginRecordSet(req *request.PageRequest) *LoginRecordSet {
	return &LoginRecordSet{
		PageRequest: req,
		Items:       []*LoginLog{},
	}
}

// LoginRecordSet todo
type LoginRecordSet struct {
	*request.PageRequest

	Total int64       `json:"total"`
	Items []*LoginLog `json:"items"`
}

// Add 添加应用
func (s *LoginRecordSet) Add(item *LoginLog) {
	s.Items = append(s.Items, item)
}

// Length 长度
func (s *LoginRecordSet) Length() int {
	return len(s.Items)
}

// IsEmpty 长度
func (s *LoginRecordSet) IsEmpty() bool {
	return len(s.Items) == 0
}

// NewDefaultLoginLogData todo
func NewDefaultLoginLogData() *LoginLogData {
	return &LoginLogData{
		Session: token.NewSession(),
		LoginAt: ftime.Now(),
	}
}

// NewDefaultLogoutLogData todo
func NewDefaultLogoutLogData() *LoginLogData {
	return &LoginLogData{
		Session:  token.NewSession(),
		LogoutAt: ftime.Now(),
	}
}

// LoginLogData todo
type LoginLogData struct {
	*token.Session  `bson:"-" json:"-"`
	Account         string          `bson:"account" json:"account" alidate:"required"`                   // 用户
	LoginAt         ftime.Time      `bson:"login_at" json:"login_at" alidate:"required"`                 // 登录时间
	LogoutAt        ftime.Time      `bson:"logout_at" json:"logout_at"`                                  // 登出时间
	ApplicationID   string          `bson:"application_id" json:"application_id" alidate:"required"`     // 用户通过哪个端登录的
	ApplicationName string          `bson:"application_name" json:"application_name" alidate:"required"` // 用户通过哪个端登录的
	GrantType       token.GrantType `bson:"grant_type" json:"grant_type" alidate:"required"`             // 登录方式
	LoginIP         string          `bson:"login_ip" json:"login_ip" alidate:"required"`                 // 登录IP
	userAgent       string          `bson:"-"`
}

// ActionType 是否是登录日志
func (l *LoginLogData) ActionType() ActionType {
	if !l.LoginAt.T().IsZero() {
		return LoginAction
	}

	if !l.LogoutAt.T().IsZero() {
		return LogoutAction
	}

	return LoginAction
}

// Validate 校验必填参数
func (l *LoginLogData) Validate() error {
	if l.LoginAt.T().IsZero() && l.LogoutAt.T().IsZero() {
		return fmt.Errorf("login_at or logout_at need one")
	}

	return validate.Struct(l)
}

// WithUserAgent 记录UserAgent
func (l *LoginLogData) WithUserAgent(ua string) {
	l.userAgent = ua
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
