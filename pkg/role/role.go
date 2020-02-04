/*
admin   管理员
read_olny:
	 *
read_write:
	 *

ops    运维
 read_only:
	<通过Label筛选资源>
	资源:   a, b, c, d
	范围:   *
 read_write:
	资源:   a, b, c, d
	范围:   online

policy:
namespace user     role
admin     admin    admin
admin     yumaojun ops
projectA  yumaojun dev
projectB  yumaojun visitor

tk
 namesapce: projectA   role:    dev
*/

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
