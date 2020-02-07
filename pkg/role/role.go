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

// Type 角色类型
type Type string

const (
	// BuildInType 内建角色, 系统初始时创建
	BuildInType Type = "build_in"
	// CustomType 用户自定义角色
	CustomType Type = "custom"
)

// EffectType 授权效力包括两种：允许（Allow）和拒绝（Deny）
type EffectType string

const (
	// Allow 允许访问
	Allow EffectType = "allow"
	// Deny 拒绝访问
	Deny EffectType = "deny"
)

// Role is rbac's role
type Role struct {
	Type        Type        `json:"type"`
	Name        string      `json:"name"`                  // 角色名称
	Description string      `json:"description,omitempty"` // 角色描述
	CreateAt    int64       `json:"create_at,omitempty"`   // 创建时间`
	UpdateAt    int64       `json:"update_at,omitempty"`   // 更新时间
	Read        *Permission `json:"read,omitempty"`        // 读权限
	Write       *Permission `json:"write,omitempty"`       // 写权限
}

// Permission 权限
type Permission struct {
	Effect   EffectType `json:"effect,omitempty"`
	IsALL    bool       `json:"is_all"`
	Resource []string   `json:"resource,omitempty"` // 资源列表
	Scope    []string   `json:"scope,omitempty"`    // 范围列表
}

// CheckPermission 检测该角色是否具有该权限
func (r *Role) CheckPermission() error {
	return nil
}
