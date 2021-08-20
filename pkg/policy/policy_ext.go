package policy

import (
	"context"
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/infraboard/mcube/exception"
	page "github.com/infraboard/mcube/pb/page"
	"github.com/infraboard/mcube/types/ftime"

	"github.com/infraboard/keyauth/pkg/namespace"
	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user"
)

// New 新实例
func New(tk *token.Token, req *CreatePolicyRequest) (*Policy, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	if tk == nil {
		return nil, exception.NewUnauthorized("token required")
	}

	p := &Policy{
		CreateAt:    ftime.Now().Timestamp(),
		UpdateAt:    ftime.Now().Timestamp(),
		Creater:     tk.Account,
		Domain:      tk.Domain,
		NamespaceId: req.NamespaceId,
		Account:     req.Account,
		RoleId:      req.RoleId,
		Scope:       req.Scope,
		ExpiredTime: req.ExpiredTime,
		Type:        req.Type,
	}
	p.genID()

	return p, nil
}

// NewDefaultPolicy todo
func NewDefaultPolicy() *Policy {
	return &Policy{}
}

func (p *Policy) genID() {
	h := fnv.New32a()
	hashedStr := fmt.Sprintf("%s-%s-%s-%s",
		p.Domain, p.NamespaceId, p.Account, p.RoleId)

	h.Write([]byte(hashedStr))
	p.Id = fmt.Sprintf("%x", h.Sum32())
}

// CheckDependence todo
func (p *Policy) CheckDependence(ctx context.Context, u user.UserServiceServer, r role.RoleServiceServer, ns namespace.NamespaceServiceServer) (*user.User, error) {
	account, err := u.DescribeAccount(ctx, user.NewDescriptAccountRequestWithAccount(p.Account))
	if err != nil {
		return nil, fmt.Errorf("check user error, %s", err)
	}

	_, err = r.DescribeRole(ctx, role.NewDescribeRoleRequestWithID(p.RoleId))
	if err != nil {
		return nil, fmt.Errorf("check role error, %s", err)
	}

	if !p.IsAllNamespace() {
		_, err = ns.DescribeNamespace(ctx, namespace.NewNewDescriptNamespaceRequestWithID(p.NamespaceId))
		if err != nil {
			return nil, fmt.Errorf("check namespace error, %s", err)
		}
	}

	return account, nil
}

// IsAllNamespace 是否是对账所有namespace的测试
func (p *Policy) IsAllNamespace() bool {
	return p.NamespaceId == "*"
}

// NewCreatePolicyRequest 请求实例
func NewCreatePolicyRequest() *CreatePolicyRequest {
	return &CreatePolicyRequest{}
}

// Validate 校验请求合法
func (req *CreatePolicyRequest) Validate() error {
	return validate.Struct(req)
}

// NewPolicySet todo
func NewPolicySet() *Set {
	return &Set{
		Items: []*Policy{},
	}
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
func (s *Set) GetRoles(ctx context.Context, r role.RoleServiceServer, withPermission bool) (*role.Set, error) {
	set := role.NewRoleSet()
	for i := range s.Items {
		req := role.NewDescribeRoleRequestWithID(s.Items[i].RoleId)
		req.WithPermissions = withPermission

		ins, err := r.DescribeRole(ctx, req)
		if err != nil {
			return nil, err
		}
		// 继承policy上的范围限制
		ins.Scope = s.Items[i].Scope
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
			rns = append(rns, item.RoleId)
		}
	}

	if len(rns) == 0 {
		rns = append(rns, "vistor")
	}

	return rns
}

// GetScope todo
func (s *Set) GetScope(account string) string {
	scopes := []string{}
	for i := range s.Items {
		item := s.Items[i]
		if item.Account == account {
			scopes = append(scopes, item.Scope)
		}
	}
	return strings.Join(scopes, " ")
}

func (s *Set) GetNamespace() (nss []string) {
	nmap := map[string]struct{}{}
	for i := range s.Items {
		nmap[s.Items[i].NamespaceId] = struct{}{}
	}

	for k := range nmap {
		nss = append(nss, k)
	}

	return
}

func (s *Set) GetNamespaceWithPage(page *page.PageRequest) (nss []string, total int64) {
	nmap := map[string]struct{}{}
	for i := range s.Items {
		// 如果policy的namespace为* , 表示所有namespace
		if s.Items[i].NamespaceId == "*" {
			return []string{}, 0
		}

		nmap[s.Items[i].NamespaceId] = struct{}{}
	}

	offset := page.PageSize*page.PageNumber - 1
	end := offset + page.PageSize

	var count uint64 = 0
	for k := range nmap {
		if count >= offset && count < end {
			nss = append(nss, k)
		}

		count++
	}

	return nss, int64(len(nmap))
}
