package policy

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// Service 策略服务
type Service interface {
	CreatePolicy(req *CreatePolicyRequest) (*Policy, error)
	QueryPolicy(req *QueryPolicyRequest) (*Set, error)
}

// NewQueryPolicyRequest 列表查询请求
func NewQueryPolicyRequest(pageReq *request.PageRequest) *QueryPolicyRequest {
	return &QueryPolicyRequest{
		PageRequest: pageReq,
	}
}

// QueryPolicyRequest 获取子账号列表
type QueryPolicyRequest struct {
	*request.PageRequest

	UserID    string `json:"user_id,omitempty"`
	RoleName  string `json:"role_name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

// NewCreateUserRequest 创建请求
func NewCreateUserRequest() *QueryPolicyRequest {
	return &QueryPolicyRequest{}
}

// Validate 校验请求是否合法
func (req *QueryPolicyRequest) Validate() error {
	return validate.Struct(req)
}

// DescribePolicyRequest todo
type DescribePolicyRequest struct {
	ID string
}
