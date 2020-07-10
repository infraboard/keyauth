package policy

import (
	"crypto/sha1"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"

	"github.com/infraboard/keyauth/pkg/user"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// Service 策略服务
type Service interface {
	CreatePolicy(createrID string, req *CreatePolicyRequest) (*Policy, error)
	QueryPolicy(req *QueryPolicyRequest) (*PolicySet, error)
}

// NewCreatePolicyRequest 请求实例
func NewCreatePolicyRequest() *CreatePolicyRequest {
	return &CreatePolicyRequest{}
}

// CreatePolicyRequest 创建策略的请求
type CreatePolicyRequest struct {
	UserType    user.Type   `json:"user_type" validate:"required,lte=60"` // 用户类型
	UserID      string      `json:"user_id" validate:"required,lte=120"`  // 用户ID
	RoleName    string      `json:"role_name" validate:"required,lte=40"` // 角色名称
	ExpiredTime *ftime.Time `json:"expired_time"`                         // 策略过期时间
	Scope       string      `json:"scope" validate:"lte=120"`             // 范围
}

// Validate 校验请求合法
func (req *CreatePolicyRequest) Validate() error {
	return validate.Struct(req)
}

func (req *CreatePolicyRequest) hashedID() string {
	inst := sha1.New()
	hashedStr := fmt.Sprintf("%s-%s-%s-%s",
		req.Scope, req.UserID, req.UserType, req.RoleName)
	inst.Write([]byte(hashedStr))
	return fmt.Sprintf("%x", inst.Sum([]byte("")))
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

	UserID   string `json:"user_id,omitempty"`
	RoleName string `json:"role_name,omitempty"`
	Scope    string `json:"scope,omitempty"`
}

// NewCreateUserRequest 创建请求
func NewCreateUserRequest() *QueryPolicyRequest {
	return &QueryPolicyRequest{}
}

// Validate 校验请求是否合法
func (req *QueryPolicyRequest) Validate() error {
	return validate.Struct(req)
}
