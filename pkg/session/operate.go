package session

// // OperateLog 操作日志
// type OperateLog struct {
// 	ID              string `bson:"_id" json:"id"`
// 	LoginLogID      string `bson:"login_id" json:"login_id"`
// 	Domain          string `bson:"domain" json:"domain" alidate:"required"` // 所处域
// 	*OperateLogData `bson:",inline"`
// }

// // OperateLogData todo
// type OperateLogData struct {
// 	Account         string     `bson:"account" json:"account" alidate:"required"`       // 用户
// 	OperateAt       ftime.Time `bson:"operate_at" json:"operate_at" alidate:"required"` // 操作时间
// 	ApplicationID   string     `bson:"application_id" json:"application_id"`            // 用户通过哪个端登录的
// 	ApplicationName string     `bson:"application_name" json:"application_name"`        // 用户通过哪个端登录的
// 	ResourceType    string     `bson:"resource_type" json:"resource_type"`              // 资源类型
// 	Action          string     `bson:"action" json:"action"`                            // 操作资源的动作
// 	Result          Result     `bson:"result" json:"result"`                            // 登录状态 (成功或者失败)
// 	Comment         string     `bson:"comment" json:"comment"`                          // 备注, 用于记录失败原因
// 	ResourceID      string     `bson:"resource_id" json:"resource_id"`                  // 资源ID
// 	ResourceName    string     `bson:"resource_name" json:"resource_name"`              // 资源名称
// }

// // OperateRecordSet todo
// type OperateRecordSet struct {
// 	*request.PageRequest

// 	Total int64         `json:"total"`
// 	Items []*OperateLog `json:"items"`
// }
