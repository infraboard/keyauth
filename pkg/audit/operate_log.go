package audit

import "github.com/infraboard/mcube/types/ftime"

// OperateLog 操作日志
type OperateLog struct {
	ID              string `bson:"_id" json:"id"`
	LoginLogID      string `bson:"login_id" json:"login_id"`
	Domain          string `bson:"domain" json:"domain" alidate:"required"` // 所处域
	*OperateLogData `bson:",inline"`
}

// OperateLogData todo
type OperateLogData struct {
	Account       string     `bson:"account" json:"account" alidate:"required"`       // 用户
	OperateAt     ftime.Time `bson:"operate_at" json:"operate_at" alidate:"required"` // 操作时间
	ApplicationID string     `bson:"application_id" json:"application_id"`            // 用户通过哪个端登录的
	ResourceType  string     `bson:"resource_type" json:"resource_type"`              // 资源类型
	Action        string     `bson:"action" json:"action"`                            // 操作资源的动作
	Result        Result     `bson:"result" json:"result"`                            // 登录状态 (成功或者失败)
}
