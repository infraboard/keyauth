package verifycode

import (
	"fmt"
	"hash/fnv"

	"github.com/infraboard/keyauth/pkg/token"
)

// Service 验证码服务
type Service interface {
	IssueCode(*IssueCodeRequest) (string, error)
	CheckCode(*CheckCodeRequest) error
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
		return req.IssueByPassRequest.Username
	case IssueTypeToken:
		return req.IssueByTokenRequest.GetAccount()
	default:
		return ""
	}
}

// IssueByPassRequest todo
type IssueByPassRequest struct {
	Username     string `json:"username" validate:"required"`
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

// NewCheckCodeRequest todo
func NewCheckCodeRequest(username, number string) *CheckCodeRequest {
	return &CheckCodeRequest{
		Username: username,
		Number:   number,
	}
}

// CheckCodeRequest 验证码校验请求
type CheckCodeRequest struct {
	Username string `bson:"username" json:"username" validate:"required"`
	Number   string `bson:"number" json:"number" validate:"required"`
}

// Validate todo
func (req *CheckCodeRequest) Validate() error {
	return validate.Struct(req)
}

// HashID todo
func (req *CheckCodeRequest) HashID() string {
	hash := fnv.New32a()
	hash.Write([]byte(req.Username))
	hash.Write([]byte(req.Number))
	return fmt.Sprintf("%x", hash.Sum32())
}
