package role

import (
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"

	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
)

// New 新创建一个Role
func New(tk *token.Token, req *CreateRoleRequest) (*Role, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	if !tk.UserType.IsIn(types.UserType_SUPPER) && !req.IsCumstomType() {
		return nil, fmt.Errorf("only supper account can create global and build role")
	}

	r := &Role{
		Id:          xid.New().String(),
		CreateAt:    ftime.Now().Timestamp(),
		UpdateAt:    ftime.Now().Timestamp(),
		Domain:      tk.Domain,
		Creater:     tk.Account,
		Type:        req.Type,
		Name:        req.Name,
		Meta:        req.Meta,
		Description: req.Description,
	}
	r.Permissions = NewPermission(r.Id, tk.Account, req.Permissions)
	return r, nil
}

// NewDefaultRole 默认实例
func NewDefaultRole() *Role {
	return &Role{
		Permissions: []*Permission{},
		Type:        RoleType_CUSTOM,
		Meta:        map[string]string{},
	}
}

// HasPermission 权限判断
func (r *Role) HasPermission(ep *endpoint.Endpoint) (*Permission, bool, error) {
	var (
		rok, lok bool
	)
	for i := range r.Permissions {
		rok = r.Permissions[i].MatchResource(ep.ServiceId, ep.Entry.Resource)
		lok = r.Permissions[i].MatchLabel(ep.Entry.Labels)
		if rok && lok {
			return r.Permissions[i], true, nil
		}
	}
	return nil, false, nil
}

func NewDeaultPermission() *Permission {
	return &Permission{}
}

func NewPermission(roleID, creater string, perms []*CreatePermssionRequest) []*Permission {
	set := []*Permission{}
	for i := range perms {
		set = append(set, &Permission{
			Id:           PermissionHash(roleID, perms[i]),
			RoleId:       roleID,
			CreateAt:     ftime.Now().Timestamp(),
			Creater:      creater,
			Effect:       perms[i].Effect,
			ServiceId:    perms[i].ServiceId,
			ResourceName: perms[i].ResourceName,
			LabelKey:     perms[i].LabelKey,
			MatchAll:     perms[i].MatchAll,
			LabelValues:  perms[i].LabelValues,
		})
	}
	return set
}

func PermissionHash(namesapce string, perm *CreatePermssionRequest) string {
	h := fnv.New32a()

	h.Write([]byte(namesapce + perm.Effect.String() + perm.ServiceId + perm.ResourceName))
	return fmt.Sprintf("%x", h.Sum32())
}

// NewCreateRoleRequest 实例化请求
func NewCreateRoleRequest() *CreateRoleRequest {
	return &CreateRoleRequest{
		Permissions: []*CreatePermssionRequest{},
		Type:        RoleType_CUSTOM,
	}
}

// IsCumstomType todo
func (req *CreateRoleRequest) IsCumstomType() bool {
	return req.Type.Equal(RoleType_CUSTOM)
}

// Validate 请求校验
func (req *CreateRoleRequest) Validate() error {
	pc := len(req.Permissions)
	if pc > RoleMaxPermission {
		return fmt.Errorf("role permission overed max count: %d",
			RoleMaxPermission)
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
func NewRoleSet() *Set {
	return &Set{
		Items: []*Role{},
	}
}

// Permissions todo
func (s *Set) Permissions() *PermissionSet {
	ps := NewPermissionSet()

	for i := range s.Items {
		ps.Add(s.Items[i].Permissions...)
	}

	return ps
}

// Add todo
func (s *Set) Add(item *Role) {
	s.Total++
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
			// 补充权限访问范围
			p.Scope = s.Items[i].Scope
			return p, ok, nil
		}
	}

	return nil, false, nil
}

func (s *Set) RoleNames() []string {
	set := []string{}
	for i := range s.Items {
		set = append(set, s.Items[i].Name)

	}

	return set
}

// NewDefaultPermission todo
func NewDefaultPermission() *CreatePermssionRequest {
	return &CreatePermssionRequest{
		Effect: EffectType_ALLOW,
	}
}

// Validate todo
func (p *CreatePermssionRequest) Validate() error {
	if p.ServiceId == "" || p.ResourceName == "" || p.LabelKey == "" {
		return fmt.Errorf("permisson required service_id, resource_name and label_key")
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
func (p *Permission) MatchResource(serviceID, resourceName string) bool {
	// 服务匹配
	if p.ServiceId != "*" && p.ServiceId != serviceID {
		return false
	}

	// 资源匹配
	if p.ResourceName != "*" && p.ResourceName != resourceName {
		return false
	}

	return true
}

// MatchLabel 匹配Label
func (p *Permission) MatchLabel(label map[string]string) bool {
	for k, v := range label {
		// 匹配key
		if p.LabelKey == "*" || p.LabelKey == k {
			// 匹配value
			if p.isMatchAllValue() {
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

func (p *Permission) isMatchAllValue() bool {
	if p.MatchAll {
		return true
	}

	for i := range p.LabelValues {
		if p.LabelValues[i] == "*" {
			return true
		}
	}

	return false
}

// NewPermissionSet todo
func NewPermissionSet() *PermissionSet {
	return &PermissionSet{
		Items: []*Permission{},
	}
}

// Add todo
func (s *PermissionSet) Add(items ...*Permission) {
	s.Items = append(s.Items, items...)
}

func (req *QueryPermissionRequest) Validate() error {
	return validate.Struct(req)
}
