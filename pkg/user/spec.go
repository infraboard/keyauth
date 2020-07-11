package user

import (
	"errors"
	"fmt"

	"github.com/infraboard/mcube/http/request"

	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
)

// Service 用户服务
type Service interface {
	// 查询用户
	QueryAccount(types.Type, *QueryAccountRequest) (*Set, error)
	// 创建用户
	CreateAccount(types.Type, *CreateUserRequest) (*User, error)
	// 更新用户密码
	UpdateAccountPassword(userName, oldPass, newPass string) error
	// 获取账号Profile
	DescribeAccount(req *DescriptAccountRequest) (*User, error)
	// 警用账号
	BlockAccount(id, reason string) error
	// DeleteAccount 删除用户
	DeleteAccount(id string) error
}

// NewDescriptAccountRequest 查询详情请求
func NewDescriptAccountRequest() *DescriptAccountRequest {
	return &DescriptAccountRequest{}
}

// NewDescriptAccountRequestWithID 查询详情请求
func NewDescriptAccountRequestWithID(id string) *DescriptAccountRequest {
	return &DescriptAccountRequest{ID: id}
}

// DescriptAccountRequest 查询用户详情请求
type DescriptAccountRequest struct {
	ID      string `json:"id,omitempty"`
	Account string `json:"account,omitempty"`
}

func (req *DescriptAccountRequest) String() string {
	return fmt.Sprint(*req)
}

// Validate 校验详情查询
func (req *DescriptAccountRequest) Validate() error {
	if req.ID == "" && req.Account == "" {
		return errors.New("id or account is required")
	}

	return nil
}

// NewQueryAccountRequest 列表查询请求
func NewQueryAccountRequest(pageReq *request.PageRequest) *QueryAccountRequest {
	return &QueryAccountRequest{
		PageRequest: pageReq,
	}
}

// QueryAccountRequest 获取子账号列表
type QueryAccountRequest struct {
	*request.PageRequest
	*DescriptAccountRequest

	DomainID  string `json:"domain_id,omitempty"`
	ProjectID string `json:"project_id,omitempty"`
}

// NewCreateUserRequest 创建请求
func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Session: token.NewSession(),
	}
}
