package user

import "github.com/infraboard/mcube/http/request"

// Service 用户服务
type Service interface {
	PrimaryAccountService
	RAMAccountService

	// 更新用户密码
	UpdateAccountPassword(userName, oldPass, newPass string) error
	// 获取账号Profile
	DescribeAccount(req *DescriptAccountRequest) (*User, error)
}

// PrimaryAccountService 主账号服务
type PrimaryAccountService interface {
	// 新建主账号
	CreatePrimayAccount(req *CreateUserRequest) (*User, error)
	// 注销主账号
	DeletePrimaryAccount(userID string) error
}

// RAMAccountService 子账号服务
type RAMAccountService interface {
	QueryRAMAccount(req *QueryRAMAccountRequest) ([]*User, error)
	// CreateRAMAccount 创建子账号
	// RAM (Resource Access Management)提供的用户身份管理与访问控制服务
	CreateRAMAccount(domainID string, req *CreateUserRequest) (*User, error)
	// 注销子账号
	DeleteRAMAccount(userID string) error
}

// NewDescriptAccountRequest 查询详情请求
func NewDescriptAccountRequest() *DescriptAccountRequest {
	return &DescriptAccountRequest{}
}

// DescriptAccountRequest 查询用户详情请求
type DescriptAccountRequest struct {
	ID      string `json:"id,omitempty"`
	Account string `json:"account,omitempty"`
}

// QueryRAMAccountRequest 获取子账号列表
type QueryRAMAccountRequest struct {
	*request.PageRequest
	*DescriptAccountRequest

	DomainID  string `json:"domain_id,omitempty"`
	ProjectID string `json:"project_id,omitempty"`
}
