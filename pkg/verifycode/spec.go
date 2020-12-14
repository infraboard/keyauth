package verifycode

import "github.com/infraboard/keyauth/pkg/token"

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

// IssueByPassRequest todo
type IssueByPassRequest struct {
	Account      string `json:"account"`
	Password     string `json:"password"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// IssueByTokenRequest todo
type IssueByTokenRequest struct {
	*token.Session
}

// CheckCodeRequest 验证码校验请求
type CheckCodeRequest struct {
	Code string
}
