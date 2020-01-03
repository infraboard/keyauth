package user

import (
	"errors"
	"fmt"

	"github.com/infraboard/mcube/http/request"
)

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
	// CreateRAMAccount RAM (Resource Access Management)提供的用户身份管理与访问控制服务
	CreateRAMAccount(domainID string, req *CreateUserRequest) (*User, error)
	// 注销子账号
	DeleteRAMAccount(userID string) error
	// 查询子账号
	QueryRAMAccount(req *QueryRAMAccountRequest) ([]*User, error)
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

// QueryRAMAccountRequest 获取子账号列表
type QueryRAMAccountRequest struct {
	*request.PageRequest
	*DescriptAccountRequest

	DomainID  string `json:"domain_id,omitempty"`
	ProjectID string `json:"project_id,omitempty"`
}

// NewCreateUserRequest 创建请求
func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{}
}

// Validate 校验请求是否合法
func (req *CreateUserRequest) Validate() error {
	return validate.Struct(req)
}

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Account     string `bson:"account" json:"account,omitempty" validate:"required,lte=60"` // 用户账号名称
	Mobile      string `bson:"mobile" json:"mobile,omitempty" validate:"lte=30"`            // 手机号码, 用户可以通过手机进行注册和密码找回, 还可以通过手机号进行登录
	Email       string `bson:"email" json:"email,omitempty" validate:"lte=30"`              // 邮箱, 用户可以通过邮箱进行注册和照明密码
	Phone       string `bson:"phone" json:"phone,omitempty" validate:"lte=30"`              // 用户的座机号码
	Address     string `bson:"address" json:"address,omitempty" validate:"lte=120"`         // 用户住址
	RealName    string `bson:"real_name" json:"real_name,omitempty" validate:"lte=10"`      // 用户真实姓名
	NickName    string `bson:"nick_name" json:"nick_name,omitempty" validate:"lte=30"`      // 用户昵称, 用于在界面进行展示
	Gender      string `bson:"gender" json:"gender,omitempty" validate:"lte=10"`            // 性别
	Avatar      string `bson:"avatar" json:"avatar,omitempty" validate:"lte=300"`           // 头像
	Language    string `bson:"language" json:"language,omitempty" validate:"lte=40"`        // 用户使用的语言
	City        string `bson:"city" json:"city,omitempty" validate:"lte=40"`                // 用户所在的城市
	Province    string `bson:"province" json:"province,omitempty" validate:"lte=40"`        // 用户所在的省
	ExpiresDays int    `bson:"expires_days" json:"expires_days,omitempty"`                  // 用户多久未登录时(天), 冻结改用户, 防止僵尸用户的账号被利用
	Password    string `bson:"-" json:"password,omitempty" validate:"required,lte=80"`      // 密码相关信息
}
