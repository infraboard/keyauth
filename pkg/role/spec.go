package role

import (
	"github.com/go-playground/validator"
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
		PageRequest:         pageReq,
		DescribeRoleRequest: &DescribeRoleRequest{},
	}
}

// QueryRoleRequest 查询请求
type QueryRoleRequest struct {
	*request.PageRequest
	*DescribeRoleRequest
	Type Type
}

// DescribeRoleRequest role详情
type DescribeRoleRequest struct {
	Name string `json:"name,omitempty" validate:"required,lte=64"`
}
