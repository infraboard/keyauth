package policy

import (
	"fmt"
	"hash/fnv"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"

	"github.com/infraboard/keyauth/pkg/namespace"
	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/user/types"
)

// New 新实例
func New(req *CreatePolicyRequest) (*Policy, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	tk := req.GetToken()
	if tk == nil {
		return nil, exception.NewUnauthorized("token required")
	}

	p := &Policy{
		CreateAt:            ftime.Now(),
		UpdateAt:            ftime.Now(),
		Creater:             tk.Account,
		Domain:              tk.Domain,
		CreatePolicyRequest: req,
	}
	p.genID()

	return p, nil
}

// NewDefaultPolicy todo
func NewDefaultPolicy() *Policy {
	return &Policy{
		CreatePolicyRequest: NewCreatePolicyRequest(),
	}
}

// Policy 权限策略
type Policy struct {
	ID       string     `bson:"_id" json:"id"`              // 策略ID
	CreateAt ftime.Time `bson:"create_at" json:"create_at"` // 创建时间
	UpdateAt ftime.Time `bson:"update_at" json:"update_at"` // 更新时间
	Domain   string     `bson:"domain" json:"domain"`       // 策略所属域
	Creater  string     `bson:"creater" json:"creater"`     // 创建者ID
	UserType types.Type `bson:"user_type" json:"user_type"` // 用户类型

	*CreatePolicyRequest `bson:",inline"`

	Role      *role.Role           `bson:"-" json:"role,omitempty"`      // 关联的角色对象
	Namespace *namespace.Namespace `bson:"-" json:"namespace,omitempty"` // 关联的空间信息
}

func (p *Policy) genID() {
	h := fnv.New32a()
	hashedStr := fmt.Sprintf("%s-%s-%s-%s",
		p.Domain, p.NamespaceID, p.Account, p.RoleID)

	h.Write([]byte(hashedStr))
	p.ID = fmt.Sprintf("%x", h.Sum32())
}

// CheckDependence todo
func (req *CreatePolicyRequest) CheckDependence(u user.Service, r role.Service, ns namespace.Service) (*user.User, error) {
	account, err := u.DescribeAccount(user.NewDescriptAccountRequestWithAccount(req.Account))
	if err != nil {
		return nil, fmt.Errorf("check user error, %s", err)
	}

	_, err = r.DescribeRole(role.NewDescribeRoleRequestWithID(req.RoleID))
	if err != nil {
		return nil, fmt.Errorf("check role error, %s", err)
	}

	if !req.IsAllNamespace() {
		_, err = ns.DescribeNamespace(namespace.NewNewDescriptNamespaceRequestWithID(req.NamespaceID))
		if err != nil {
			return nil, fmt.Errorf("check namespace error, %s", err)
		}
	}

	return account, nil
}

// func NewCreatePolicyRequestFromHTTP(r *http.Request) (*CreatePolicyRequest, error) {
// 	tk, err := pkg.GetTokenFromContext(r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req := NewCreatePolicyRequest()
// 	if err := request.GetDataFromRequest(r, req); err != nil {
// 		response.Failed(w, err)
// 		return nil,
// 	}
// 	req.WithToken(tk)
// }

// NewCreatePolicyRequest 请求实例
func NewCreatePolicyRequest() *CreatePolicyRequest {
	return &CreatePolicyRequest{
		Session: token.NewSession(),
	}
}

// CreatePolicyRequest 创建策略的请求
type CreatePolicyRequest struct {
	*token.Session `bson:"-" json:"-"`
	NamespaceID    string     `bson:"namespace_id" json:"namespace_id" validate:"lte=120"` // 范围
	Account        string     `bson:"account" json:"account" validate:"required,lte=120"`  // 用户ID
	RoleID         string     `bson:"role_id" json:"role_id" validate:"required,lte=40"`   // 角色名称
	Scope          string     `bson:"scope" json:"scope"`                                  // 范围控制
	ExpiredTime    ftime.Time `bson:"expired_time" json:"expired_time"`                    // 策略过期时间
	Type           Type       `bson:"type" json:"type"`                                    // 策略的类型
}

// Validate 校验请求合法
func (req *CreatePolicyRequest) Validate() error {
	return validate.Struct(req)
}

// IsAllNamespace 是否是对账所有namespace的测试
func (req *CreatePolicyRequest) IsAllNamespace() bool {
	return req.NamespaceID == "*"
}

// NewPolicySet todo
func NewPolicySet(req *request.PageRequest) *Set {
	return &Set{
		PageRequest: req,
		Items:       []*Policy{},
	}
}

// Set 列表
type Set struct {
	*request.PageRequest

	Total int64     `json:"total"`
	Items []*Policy `json:"items"`
}

// Users 策略包含的所有用户ID, 已去重
func (s *Set) Users() []string {
	users := map[string]struct{}{}
	for i := range s.Items {
		users[s.Items[i].Account] = struct{}{}
	}

	set := make([]string, 0, len(users))
	for k := range users {
		set = append(set, k)
	}

	return set
}

// Add 添加
func (s *Set) Add(e *Policy) {
	s.Items = append(s.Items, e)
}

// Length todo
func (s *Set) Length() int {
	return len(s.Items)
}

// GetRoles todo
func (s *Set) GetRoles(r role.Service) (*role.Set, error) {
	set := role.NewRoleSet(nil)
	for i := range s.Items {
		req := role.NewDescribeRoleRequestWithID(s.Items[i].RoleID)
		req.WithPermissions = true

		ins, err := r.DescribeRole(req)
		if err != nil {
			return nil, err
		}
		set.Add(ins)
	}
	return set, nil
}

// UserRoles 获取用户的角色
func (s *Set) UserRoles(account string) []string {
	rns := []string{}
	for i := range s.Items {
		item := s.Items[i]
		if item.Account == account {
			rns = append(rns, item.RoleID)
		}
	}

	if len(rns) == 0 {
		rns = append(rns, "vistor")
	}

	return rns
}
