package verifycode

import (
	"fmt"

	"github.com/infraboard/keyauth/pkg/token"
)

// Service 验证码服务
type Service interface {
	IssueCode(*IssueCodeRequest) (*Code, error)
	CheckCode(*CheckCodeRequest) (*Code, error)
}

// NewIssueCodeRequestByPass todo
func NewIssueCodeRequestByPass() *IssueCodeRequest {
	return &IssueCodeRequest{
		IssueType:          IssueTypePass,
		IssueByPassRequest: IssueByPassRequest{},
	}
}

// NewIssueCodeRequestByToken todo
func NewIssueCodeRequestByToken() *IssueCodeRequest {
	return &IssueCodeRequest{
		IssueType: IssueTypeToken,
		IssueByTokenRequest: IssueByTokenRequest{
			Session: token.NewSession(),
		},
	}
}

// IssueCodeRequest 验证码申请请求
type IssueCodeRequest struct {
	IssueType IssueType `json:"issue_type"`
	IssueByPassRequest
	IssueByTokenRequest
}

// Validate 请求校验
func (req *IssueCodeRequest) Validate() error {
	switch req.IssueType {
	case IssueTypePass:
		return req.ValidateByPass()
	case IssueTypeToken:
		return req.ValidateByToken()
	default:
		return fmt.Errorf("unknown issue type: %s", req.IssueType)
	}
}

// Account todo
func (req *IssueCodeRequest) Account() string {
	switch req.IssueType {
	case IssueTypePass:
		return req.IssueByPassRequest.Account
	case IssueTypeToken:
		return req.IssueByTokenRequest.GetAccount()
	default:
		return ""
	}
}

// IssueByPassRequest todo
type IssueByPassRequest struct {
	Account      string `json:"account" validate:"required"`
	Password     string `json:"password" validate:"required"`
	ClientID     string `json:"client_id" validate:"required"`
	ClientSecret string `json:"client_secret" validate:"required"`
}

// ValidateByPass todo
func (req *IssueByPassRequest) ValidateByPass() error {
	return validate.Struct(req)
}

// IssueByTokenRequest todo
type IssueByTokenRequest struct {
	*token.Session
}

// ValidateByToken todo
func (req *IssueByTokenRequest) ValidateByToken() error {
	if req.GetToken() == nil {
		return fmt.Errorf("token required")
	}

	return nil
}

// CheckCodeRequest 验证码校验请求
type CheckCodeRequest struct {
	Code string
}
