package audit

import (
	"github.com/infraboard/keyauth/pkg/ip2region"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/mcube/types/ftime"
)

// LoginLog 登录日志
type LoginLog struct {
	Account          string          `bson:"account" json:"account"`                   // 用户
	LoginAt          ftime.Time      `bson:"login_at" json:"login_at"`                 // 登录时间
	LogoutAt         ftime.Time      `bson:"logout_at" json:"logout_at"`               // 登出时间
	ApplicationID    string          `bson:"application_id" json:"application_id"`     // 用户通过哪个端登录的
	ApplicationName  string          `bson:"application_name" json:"application_name"` // 用户通过哪个端登录的
	UserAgent        string          `bson:"user_agent" json:"user_agent"`             // 用户登录工具信息
	GrantType        token.GrantType `bson:"grant_type" json:"grant_type"`             // 登录方式
	LoginIP          string          `bson:"login_ip" json:"login_ip"`                 // 登录IP
	Result           Result          `bson:"result" json:"result"`                     // 登录状态 (成功或者失败)
	ip2region.IPInfo `bson:",inline"`
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
