package verifycode

import (
	"fmt"
	"hash/fnv"
)

// NewIssueCodeRequestByPass todo
func NewIssueCodeRequestByPass() *IssueCodeRequest {
	return &IssueCodeRequest{
		IssueType: IssueType_PASS,
	}
}

// NewIssueCodeRequestByToken todo
func NewIssueCodeRequestByToken() *IssueCodeRequest {
	return &IssueCodeRequest{
		IssueType: IssueType_TOKEN,
	}
}

// Validate 请求校验
func (req *IssueCodeRequest) Validate() error {
	switch req.IssueType {
	case IssueType_PASS:
		return req.ValidateByPass()
	case IssueType_TOKEN:
		return req.ValidateByToken()
	default:
		return fmt.Errorf("unknown issue type: %s", req.IssueType)
	}
}

// Account todo
func (req *IssueCodeRequest) Account() string {
	switch req.IssueType {
	case IssueType_PASS:
		return req.Username
	case IssueType_TOKEN:
		return ""
		// return req.IssueByToken.GetAccount()
	default:
		return ""
	}
}

// ValidateByPass todo
func (req *IssueCodeRequest) ValidateByPass() error {
	return validate.Struct(req)
}

// ValidateByToken todo
func (req *IssueCodeRequest) ValidateByToken() error {
	return nil
}

// NewCheckCodeRequest todo
func NewCheckCodeRequest(username, number string) *CheckCodeRequest {
	return &CheckCodeRequest{
		Username: username,
		Number:   number,
	}
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

// NewIssueCodeResponse todo
func NewIssueCodeResponse(message string) *IssueCodeResponse {
	return &IssueCodeResponse{Message: message}
}
