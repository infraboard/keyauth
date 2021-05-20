package user

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
)

// NewDescriptAccountRequest 查询详情请求
func NewDescriptAccountRequest() *DescribeAccountRequest {
	return &DescribeAccountRequest{}
}

// NewDescriptAccountRequestWithAccount 查询详情请求
func NewDescriptAccountRequestWithAccount(accout string) *DescribeAccountRequest {
	return &DescribeAccountRequest{Account: accout}
}

// Validate 校验详情查询
func (req *DescribeAccountRequest) Validate() error {
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

	query.DepartmentId = qs.Get("department_id")
	query.Keywords = qs.Get("keywords")
	query.NamespaceId = qs.Get("namespace_id")

	query.WithDepartment = qs.Get("with_department") == "true"
	query.SkipItems = qs.Get("skip_items") == "true"
	query.WithAllSub = qs.Get("with_all_sub") == "true"

	accounts := qs.Get("accounts")
	if accounts != "" {
		query.Accounts = strings.Split(accounts, ",")
	}
	return query
}

// NewQueryAccountRequest 列表查询请求
func NewQueryAccountRequest() *QueryAccountRequest {
	return &QueryAccountRequest{
		Page:           &request.NewPageRequest(20, 1).PageRequest,
		WithDepartment: false,
		SkipItems:      false,
	}
}

// SetPageRequest todo
func (req *QueryAccountRequest) SetPageRequest(page *request.PageRequest) {
	req.Page = &page.PageRequest
}

// Validate 校验查询参数
func (req *QueryAccountRequest) Validate() error {
	return nil
}

// NewCreateUserRequestWithLDAPSync todo
func NewCreateUserRequestWithLDAPSync(username, password string) *CreateAccountRequest {
	req := NewCreateUserRequest()
	req.CreateType = CreateType_LDAP_SYNC
	req.Account = username
	req.Password = password
	return req
}

// NewCreateUserRequest 创建请求
func NewCreateUserRequest() *CreateAccountRequest {
	return &CreateAccountRequest{
		Profile:     NewProfile(),
		ExpiresDays: DefaultExiresDays,
	}
}

// NewUpdatePasswordRequest todo
func NewUpdatePasswordRequest() *UpdatePasswordRequest {
	return &UpdatePasswordRequest{}
}

// IsReset 密码是否需要被重置, 如果不是自己设置的密码 都需要被用户自己重置
func (req *UpdatePasswordRequest) IsReset() bool {
	return true
}

// Validate tood
func (req *UpdatePasswordRequest) Validate() error {
	if req.Account == "" {
		return fmt.Errorf("account required")
	}

	if req.OldPass == req.NewPass {
		return fmt.Errorf("old_pass equal new_pass")
	}

	if req.NewPass == "" || req.OldPass == "" {
		return fmt.Errorf("old_pass and new_pass required")
	}

	if req.Account == req.NewPass {
		return fmt.Errorf("password must not equal account")
	}
	return nil
}

// 实现checkowner方法
func (req *UpdatePasswordRequest) CheckOwner(account string) bool {
	return req.Account == account
}

// NewGeneratePasswordRequest todo
func NewGeneratePasswordRequest() *GeneratePasswordRequest {
	return &GeneratePasswordRequest{}
}

// NewGeneratePasswordResponse todo
func NewGeneratePasswordResponse(password string) *GeneratePasswordResponse {
	return &GeneratePasswordResponse{
		Password: password,
	}
}

// NewBlockAccountRequest todo
func NewBlockAccountRequest(account, reason string) *BlockAccountRequest {
	return &BlockAccountRequest{
		Account: account,
		Reason:  reason,
	}
}

func (req *BlockAccountRequest) Validate() error {
	if req.Account == "" {
		return exception.NewBadRequest("block account required!")
	}

	return nil
}

func (req *UnBlockAccountRequest) Validate() error {
	if req.Account == "" {
		return exception.NewBadRequest("unblock account required!")
	}

	return nil
}
