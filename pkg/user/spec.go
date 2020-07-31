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
	// 获取账号Profile
	DescribeAccount(req *DescriptAccountRequest) (*User, error)
	// 警用账号
	BlockAccount(id, reason string) error
	// DeleteAccount 删除用户
	DeleteAccount(id string) error
	// 更新用户
	UpdateAccountProfile(u *User) error
	UpdateAccountPassword(req *UpdatePasswordRequest) (*Password, error)
}

// NewDescriptAccountRequest 查询详情请求
func NewDescriptAccountRequest() *DescriptAccountRequest {
	return &DescriptAccountRequest{}
}

// NewDescriptAccountRequestWithAccount 查询详情请求
func NewDescriptAccountRequestWithAccount(accout string) *DescriptAccountRequest {
	return &DescriptAccountRequest{Account: accout}
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
		Session:     token.NewSession(),
	}
}

// QueryAccountRequest 获取子账号列表
type QueryAccountRequest struct {
	*token.Session
	*request.PageRequest
	IDs         []string
	NamespaceID string
}

// Validate 校验查询参数
func (req *QueryAccountRequest) Validate() error {
	tk := req.GetToken()
	if tk == nil {
		return fmt.Errorf("token required")
	}

	return nil
}

// NewCreateUserRequest 创建请求
func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Session: token.NewSession(),
	}
}

// NewUpdatePasswordRequest todo
func NewUpdatePasswordRequest() *UpdatePasswordRequest {
	return &UpdatePasswordRequest{
		Session: token.NewSession(),
	}
}

// UpdatePasswordRequest todo
type UpdatePasswordRequest struct {
	*token.Session `json:"-"`
	OldPass        string `json:"old_pass,omitempty"`
	NewPass        string `json:"new_pass,omitempty"`
}

// Validate tood
func (req *UpdatePasswordRequest) Validate() error {
	tk := req.GetToken()
	if tk == nil {
		return fmt.Errorf("token required")
	}

	if req.OldPass == req.NewPass {
		return fmt.Errorf("old_pass equal new_pass")
	}

	if req.NewPass == "" || req.OldPass == "" {
		return fmt.Errorf("old_pass and new_pass required")
	}

	return nil
}
