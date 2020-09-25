package role

import (
	"fmt"
	"strings"

	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"

	"github.com/infraboard/keyauth/pkg/endpoint"
)

const (
	// MaxPermissionCount 一个角色最多可以容纳的权限条数
	MaxPermissionCount = 500
)

// New 新创建一个Role
func New(req *CreateRoleRequest) (*Role, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	tk := req.GetToken()
	if tk == nil {
		return nil, fmt.Errorf("token required")
	}

	if !tk.UserType.Is(types.SupperAccount) && !req.IsCumstomType() {
		return nil, fmt.Errorf("only supper account can create global and build role")
	}

	return &Role{
		ID:                xid.New().String(),
		CreateAt:          ftime.Now(),
		UpdateAt:          ftime.Now(),
		Domain:            tk.Domain,
		Creater:           tk.Account,
		CreateRoleRequest: req,
	}, nil
}

// NewDefaultRole 默认实例
func NewDefaultRole() *Role {
	return &Role{
		CreateRoleRequest: NewCreateRoleRequest(),
	}
}

// Role is rbac's role
type Role struct {
	ID                 string     `bson:"_id" json:"id"`                        // 角色ID
	CreateAt           ftime.Time `bson:"create_at" json:"create_at,omitempty"` // 创建时间`
	UpdateAt           ftime.Time `bson:"update_at" json:"update_at,omitempty"` // 更新时间
	Domain             string     `bson:"domain" json:"domain,omitempty"`       // 角色所属域
	Creater            string     `bson:"creater" json:"creater"`               // 创建人
	*CreateRoleRequest `bson:",inline"`
}

// HasPermission 权限判断
func (r *Role) HasPermission(ep *endpoint.Endpoint) (*Permission, bool, error) {
	var (
		rok, lok bool
	)
	for i := range r.Permissions {
		rok = r.Permissions[i].MatchResource(ep.Resource)
		lok = r.Permissions[i].MatchLabel(ep.Labels)
		if rok && lok {
			return r.Permissions[i], true, nil
		}
	}
	return nil, false, nil
}

// NewCreateRoleRequest 实例化请求
func NewCreateRoleRequest() *CreateRoleRequest {
	return &CreateRoleRequest{
		Session:     token.NewSession(),
		Permissions: []*Permission{},
		Type:        CustomType,
	}
}

// CreateRoleRequest 创建应用请求
type CreateRoleRequest struct {
	*token.Session `bson:"-" json:"-"`
	Type           Type          `bson:"type" json:"type"`                                            // 角色类型
	Name           string        `bson:"name" json:"name,omitempty" validate:"required,lte=30"`       // 应用名称
	Description    string        `bson:"description" json:"description,omitempty" validate:"lte=400"` // 应用简单的描述
	Permissions    []*Permission `bson:"permissions" json:"permissions,omitempty"`                    // 读权限
}

// IsCumstomType todo
func (req *CreateRoleRequest) IsCumstomType() bool {
	return req.Type.Is(CustomType)
}

// Validate 请求校验
func (req *CreateRoleRequest) Validate() error {
	tk := req.GetToken()
	if tk == nil {
		return fmt.Errorf("token required")
	}

	pc := len(req.Permissions)
	if pc > MaxPermissionCount {
		return fmt.Errorf("role permission overed max count: %d",
			MaxPermissionCount)
	}

	errs := []string{}
	for i := range req.Permissions {
		if err := req.Permissions[i].Validate(); err != nil {
			errs = append(errs, err.Error())
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("validate permission error, %s", strings.Join(errs, ","))
	}

	return validate.Struct(req)
}

// CheckPermission 检测该角色是否具有该权限
func (r *Role) CheckPermission() error {
	return nil
}

// NewRoleSet 实例化make
func NewRoleSet(req *request.PageRequest) *Set {
	return &Set{
		PageRequest: req,
		Items:       []*Role{},
	}
}

// Set 角色集合
type Set struct {
	*request.PageRequest

	Total int64   `json:"total"`
	Items []*Role `json:"items"`
}

// Permissions todo
func (s *Set) Permissions() *PermissionSet {
	ps := NewPermissionSet(nil)

	for i := range s.Items {
		ps.Add(s.Items[i].Permissions...)
	}

	return ps
}

// Add todo
func (s *Set) Add(item *Role) {
	s.Items = append(s.Items, item)
}

// HasPermission todo
func (s *Set) HasPermission(ep *endpoint.Endpoint) (*Permission, bool, error) {
	for i := range s.Items {
		p, ok, err := s.Items[i].HasPermission(ep)
		if err != nil {
			return nil, false, err
		}
		if ok {
			return p, ok, nil
		}
	}

	return nil, false, nil
}

// NewDefaultPermission todo
func NewDefaultPermission() *Permission {
	return &Permission{
		Effect: Allow,
	}
}

// Permission 权限
type Permission struct {
	Effect       EffectType `bson:"effect" json:"effect,omitempty"`               // 效力
	ResourceName string     `bson:"resource_name" json:"resource_name,omitempty"` // 资源列表
	LabelKey     string     `bson:"label_key" json:"label_key,omitempty"`         // 维度
	MatchAll     bool       `bson:"match_all" json:"match_all"`                   // 适配所有值
	LabelValues  []string   `bson:"label_values" json:"label_values,omitempty"`   // 标识值
}

// Validate todo
func (p *Permission) Validate() error {
	if p.ResourceName == "" || p.LabelKey == "" {
		return fmt.Errorf("permisson required resource_name and label_key")
	}

	if len(p.LabelValues) == 0 {
		return fmt.Errorf("permission label_values required")
	}

	return nil
}

// ID 计算唯一ID
func (p *Permission) ID(namespace string) string {
	return namespace + "." + p.ResourceName
}

// MatchResource 检测资源是否匹配
func (p *Permission) MatchResource(r string) bool {
	if p.ResourceName == "*" {
		return true
	}
	return p.ResourceName == r
}

// MatchLabel 匹配Label
func (p *Permission) MatchLabel(label map[string]string) bool {
	for k, v := range label {
		if p.LabelKey == "*" || p.LabelKey == k {
			if p.MatchAll {
				return true
			}
			for i := range p.LabelValues {
				if p.LabelValues[i] == v {
					return true
				}
			}
		}
	}

	return false
}

// NewPermissionSet todo
func NewPermissionSet(req *request.PageRequest) *PermissionSet {
	return &PermissionSet{
		PageRequest: req,
		Items:       []*Permission{},
	}
}

// PermissionSet 用户列表
type PermissionSet struct {
	*request.PageRequest

	Total int64         `json:"total"`
	Items []*Permission `json:"items"`
}

// Add todo
func (s *PermissionSet) Add(items ...*Permission) {
	s.Items = append(s.Items, items...)
}
