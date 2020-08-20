package audit

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/mssola/user_agent"

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
		LoginLogData: data,
	}

	return log
}

// LoginLog 登录日志
type LoginLog struct {
	*LoginLogData
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

// LoginLogData todo
type LoginLogData struct {
	Account         string          `bson:"account" json:"account" alidate:"required"`                   // 用户
	LoginAt         ftime.Time      `bson:"login_at" json:"login_at" alidate:"required"`                 // 登录时间
	LogoutAt        ftime.Time      `bson:"logout_at" json:"logout_at"`                                  // 登出时间
	ApplicationID   string          `bson:"application_id" json:"application_id" alidate:"required"`     // 用户通过哪个端登录的
	ApplicationName string          `bson:"application_name" json:"application_name" alidate:"required"` // 用户通过哪个端登录的
	GrantType       token.GrantType `bson:"grant_type" json:"grant_type" alidate:"required"`             // 登录方式
	LoginIP         string          `bson:"login_ip" json:"login_ip" alidate:"required"`                 // 登录IP
	Result          Result          `bson:"result" json:"result" alidate:"required"`                     // 登录状态 (成功或者失败)
	Comment         string          `bson:"comment" json:"comment"`                                      // 备注 主要用于描述失败原因
	userAgent       string          `bson:"-"`
}

// Validate 校验必填参数
func (l *LoginLogData) Validate() error {
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

// OperateLog 操作日志
type OperateLog struct {
	Account       string     `bson:"account" json:"account"`               // 用户
	OperateAt     ftime.Time `bson:"operate_at" json:"operate_at"`         // 操作时间
	ApplicationID string     `bson:"application_id" json:"application_id"` // 用户通过哪个端登录的
	ResourceType  string     `bson:"resource_type" json:"resource_type"`   // 资源类型
	Action        string     `bson:"action" json:"action"`                 // 操作资源的动作
	Result        Result     `bson:"result" json:"result"`                 // 登录状态 (成功或者失败)
}
