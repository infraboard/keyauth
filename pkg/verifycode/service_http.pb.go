// Code generated by protoc-gen-go-http. DO NOT EDIT.

package verifycode

import (
	http "github.com/infraboard/mcube/pb/http"
)

// HttpEntry todo
func HttpEntry() *http.EntrySet {
	set := &http.EntrySet{
		Items: []*http.Entry{
			{
				GrpcPath:         "/keyauth.verifycode.VerifyCodeService/IssueCode",
				FunctionName:     "IssueCode",
				Path:             "/verify_code/issue",
				Method:           "POST",
				Resource:         "verify_code",
				AuthEnable:       false,
				PermissionEnable: false,
				Labels:           map[string]string{},
			},
			{
				GrpcPath:         "/keyauth.verifycode.VerifyCodeService/CheckCode",
				FunctionName:     "CheckCode",
				Path:             "/verify_code/check",
				Method:           "POST",
				Resource:         "verify_code",
				AuthEnable:       false,
				PermissionEnable: false,
				Labels:           map[string]string{},
			},
		},
	}
	return set
}
