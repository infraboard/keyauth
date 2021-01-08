package role

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/keyauth/pkg/token/session"
	"github.com/infraboard/mcube/http/request"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// Service 角色服务
type Service interface {
	CreateRole(req *CreateRoleRequest) (*Role, error)
	QueryRole(req *QueryRoleRequest) (*Set, error)
	DescribeRole(req *DescribeRoleRequest) (*Role, error)
	DeleteRole(name string) error
}

// NewQueryRoleRequestFromHTTP 列表查询请求
func NewQueryRoleRequestFromHTTP(r *http.Request) *QueryRoleRequest {
	page := request.NewPageRequestFromHTTP(r)

	req := NewQueryRoleRequest(page)
	qs := r.URL.Query()
	req.WithPermissions = strings.TrimSpace(qs.Get("with_permissions")) == "true"

	return req
}

// NewQueryRoleRequest 列表查询请求
func NewQueryRoleRequest(pageReq *request.PageRequest) *QueryRoleRequest {
	return &QueryRoleRequest{
		Session:         session.NewSession(),
		PageRequest:     pageReq,
		WithPermissions: false,
	}
}

// QueryRoleRequest 查询请求
type QueryRoleRequest struct {
	*session.Session
	*request.PageRequest

	Type            *Type
	WithPermissions bool
}

// Validate todo
func (req *QueryRoleRequest) Validate() error {
	if req.GetToken() == nil {
		return fmt.Errorf("token required")
	}

	return nil
}

// NewDescribeRoleRequestWithID todo
func NewDescribeRoleRequestWithID(id string) *DescribeRoleRequest {
	return &DescribeRoleRequest{
		Session:         session.NewSession(),
		ID:              id,
		WithPermissions: false,
	}
}

// NewDescribeRoleRequestWithName todo
func NewDescribeRoleRequestWithName(name string) *DescribeRoleRequest {
	return &DescribeRoleRequest{
		Session:         session.NewSession(),
		Name:            name,
		WithPermissions: false,
	}
}

// DescribeRoleRequest role详情
type DescribeRoleRequest struct {
	*session.Session
	ID              string `json:"id"`
	Name            string `json:"name,omitempty" validate:"required,lte=64"`
	WithPermissions bool
	Type            *Type
}

// Validate todo
func (req *DescribeRoleRequest) Validate() error {
	if req.ID == "" && req.Name == "" {
		return fmt.Errorf("id or name required")
	}

	return nil
}
