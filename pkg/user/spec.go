package user

import (
	"errors"
	"fmt"

	"github.com/infraboard/mcube/http/request"
)

// Service 用户服务
type Service interface {
	BaseService
	SupperAccountService
	ServiceAccountService
	PrimaryAccountService
	SubAccountService
}

// BaseService 基础功能
type BaseService interface {
	// 更新用户密码
	UpdateAccountPassword(userName, oldPass, newPass string) error
	// 获取账号Profile
	DescribeAccount(req *DescriptAccountRequest) (*User, error)
}

// SupperAccountService 超级管理员账号
type SupperAccountService interface {
	// 新建主账号
	CreateSupperAccount(req *CreateUserRequest) (*User, error)
	// 查询超级账号列表
	QuerySupperAccount(req *QueryAccountRequest) (*UserSet, error)
}

// PrimaryAccountService 主账号服务
type PrimaryAccountService interface {
	// 新建主账号
	CreatePrimayAccount(req *CreateUserRequest) (*User, error)
	// 注销主账号
	DeletePrimaryAccount(userID string) error
}

// SubAccountService 子账号服务
type SubAccountService interface {
	// CreateRAMAccount RAM (Resource Access Management)提供的用户身份管理与访问控制服务
	CreateSubAccount(domainID string, req *CreateUserRequest) (*User, error)
	// 注销子账号
	DeleteSubAccount(userID string) error
	// 查询子账号
	QuerySubAccount(req *QueryAccountRequest) (*UserSet, error)
}

// ServiceAccountService 服务账号
type ServiceAccountService interface {
	CreateServiceAccount(req *CreateUserRequest) (*User, error)
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
	return &CreateUserRequest{}
}
