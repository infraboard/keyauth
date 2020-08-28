package role

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/mcube/http/request"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// Service 角色服务
type Service interface {
	CreateRole(t Type, req *CreateRoleRequest) (*Role, error)
	QueryRole(req *QueryRoleRequest) (*Set, error)
	DescribeRole(req *DescribeRoleRequest) (*Role, error)
	DeleteRole(name string) error
}

// NewQueryRoleRequest 列表查询请求
func NewQueryRoleRequest(pageReq *request.PageRequest) *QueryRoleRequest {
	return &QueryRoleRequest{
		PageRequest: pageReq,
	}
}

// QueryRoleRequest 查询请求
type QueryRoleRequest struct {
	*request.PageRequest
	Type Type
}

// Validate todo
func (req *QueryRoleRequest) Validate() error {
	return nil
}

// NewDescribeRoleRequestWithID todo
func NewDescribeRoleRequestWithID(id string) *DescribeRoleRequest {
	return &DescribeRoleRequest{
		Session: token.NewSession(),
		ID:      id,
	}
}

// DescribeRoleRequest role详情
type DescribeRoleRequest struct {
	*token.Session
	ID   string `json:"id"`
	Name string `json:"name,omitempty" validate:"required,lte=64"`
}

// Valiate todo
func (req *DescribeRoleRequest) Valiate() error {
	if req.ID == "" && req.Name == "" {
		return fmt.Errorf("id or name required")
	}

	return nil
}
