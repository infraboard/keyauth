package user

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/infraboard/mcube/http/request"

	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
)

// Service 用户服务
type Service interface {
	// 查询用户
	QueryAccount(types.Type, *QueryAccountRequest) (*Set, error)
	// 创建用户
	CreateAccount(types.Type, *CreateAccountRequest) (*User, error)
	// 获取账号Profile
	DescribeAccount(req *DescriptAccountRequest) (*User, error)
	// 警用账号
	BlockAccount(account, reason string) error
	// DeleteAccount 删除用户
	DeleteAccount(account string) error
	// 更新用户
	UpdateAccountProfile(*UpdateAccountRequest) (*User, error)
	UpdateAccountPassword(*UpdatePasswordRequest) (*Password, error)
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
	Account string `json:"account,omitempty"`
}

func (req *DescriptAccountRequest) String() string {
	return fmt.Sprint(*req)
}

// Validate 校验详情查询
func (req *DescriptAccountRequest) Validate() error {
	if req.Account == "" {
		return errors.New("id or account is required")
	}

	return nil
}

// NewNewQueryAccountRequestFromHTTP todo
func NewNewQueryAccountRequestFromHTTP(r *http.Request) *QueryAccountRequest {
	page := request.NewPageRequestFromHTTP(r)
	query := NewQueryAccountRequest()
	query.SetPageRequest(page)

	qs := r.URL.Query()

	query.WithDepartment = qs.Get("with_department") == "true"
	query.SkipItems = qs.Get("skip_items") == "true"
	query.DepartmentID = qs.Get("department_id")
	query.WithALLSub = qs.Get("with_all_sub") == "true"
	ids := qs.Get("ids")
	if ids != "" {
		query.Accounts = strings.Split(ids, ",")
	}
	return query
}

// NewQueryAccountRequest 列表查询请求
func NewQueryAccountRequest() *QueryAccountRequest {
	return &QueryAccountRequest{
		PageRequest:    request.NewPageRequest(20, 1),
		Session:        token.NewSession(),
		WithDepartment: false,
		SkipItems:      false,
	}
}

// QueryAccountRequest 获取子账号列表
type QueryAccountRequest struct {
	*token.Session
	*request.PageRequest
	Accounts       []string
	NamespaceID    string
	WithDepartment bool
	DepartmentID   string
	WithALLSub     bool
	SkipItems      bool
}

// SetPageRequest todo
func (req *QueryAccountRequest) SetPageRequest(page *request.PageRequest) {
	req.PageRequest = page
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
func NewCreateUserRequest() *CreateAccountRequest {
	return &CreateAccountRequest{
		Session: token.NewSession(),
		Profile: NewProfile(),
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
