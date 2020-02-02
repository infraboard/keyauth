package role

// Role is rbac's role
type Role struct {
	ID          string `json:"id"`
	Name        string `json:"name"`                  // 角色名称
	Description string `json:"description,omitempty"` // 角色描述
	CreateAt    int64  `json:"create_at,omitempty"`   // 创建时间`
	UpdateAt    int64  `json:"update_at,omitempty"`   // 更新时间

}

// CheckPermission 检测该角色是否具有该权限
func (r *Role) CheckPermission() error {
	return nil
}
