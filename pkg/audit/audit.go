package audit

import (
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/mcube/types/ftime"
)

// LoginLog 登录日志
type LoginLog struct {
	Account       string          `json:"account"`        // 用户
	LoginAt       ftime.Time      `json:"login_at"`       // 登录时间
	ApplicationID string          `json:"application_id"` // 用户通过哪个端登录的
	UserAgent     string          `json:"user_agent"`     // 用户登录工具信息
	GrantType     token.GrantType `json:"grant_type"`     // 登录方式
	LoginIP       string          `json:"login_ip"`       // 登录IP
	Result        Result          `json:"result"`         // 登录状态 (成功或者失败)
}

// OperateLog 操作日志
type OperateLog struct {
	Account       string     `json:"account"`        // 用户
	OperateAt     ftime.Time `json:"operate_at"`     // 操作时间
	ApplicationID string     `json:"application_id"` // 用户通过哪个端登录的
	ResourceType  string     `json:"resource_type"`  // 资源类型
	Action        string     `json:"action"`         // 操作资源的动作
	Result        Result     `json:"result"`         // 登录状态 (成功或者失败)
}
