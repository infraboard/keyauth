// Code generated by protoc-gen-go-http. DO NOT EDIT.

package token

import (
	http "github.com/infraboard/mcube/pb/http"
)

// HttpEntry todo
func HttpEntry() *http.EntrySet {
	set := &http.EntrySet{
		Items: []*http.Entry{
			{
				GrpcPath:         "/keyauth.token.TokenService/IssueToken",
				FunctionName:     "IssueToken",
				Path:             "/oauth2/tokens",
				Method:           "POST",
				Resource:         "token",
				AuthEnable:       false,
				PermissionEnable: false,
				AuditLog:         false,
				Labels:           map[string]string{},
			},
			{
				GrpcPath:         "/keyauth.token.TokenService/ValidateToken",
				FunctionName:     "ValidateToken",
				Path:             "/oauth2/tokens",
				Method:           "GET",
				Resource:         "token",
				AuthEnable:       false,
				PermissionEnable: false,
				AuditLog:         false,
				Labels:           map[string]string{},
			},
			{
				GrpcPath:         "/keyauth.token.TokenService/DescribeToken",
				FunctionName:     "DescribeToken",
				Path:             "/applications/:id/tokens",
				Method:           "GET",
				Resource:         "token",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         false,
				Labels:           map[string]string{},
			},
			{
				GrpcPath:         "/keyauth.token.TokenService/RevolkToken",
				FunctionName:     "RevolkToken",
				Path:             "/applications/:id/tokens",
				Method:           "DELETE",
				Resource:         "token",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         false,
				Labels:           map[string]string{},
			},
			{
				GrpcPath:         "/keyauth.token.TokenService/BlockToken",
				FunctionName:     "BlockToken",
				Path:             "/applications/:id/tokens",
				Method:           "PUT",
				Resource:         "token",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         false,
				Labels:           map[string]string{},
			},
			{
				GrpcPath:         "/keyauth.token.TokenService/QueryToken",
				FunctionName:     "QueryToken",
				Path:             "/applications/:id/tokens",
				Method:           "GET",
				Resource:         "token",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         false,
				Labels:           map[string]string{},
			},
		},
	}
	return set
}
