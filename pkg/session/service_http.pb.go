// Code generated by protoc-gen-go-http. DO NOT EDIT.

package session

import (
	http "github.com/infraboard/mcube/pb/http"
)

// HttpEntry todo
func HttpEntry() *http.EntrySet {
	set := &http.EntrySet{
		Items: []*http.Entry{
			{
				GrpcPath:     "/keyauth.session.UserService/Login",
				FunctionName: "Login",
			},
			{
				GrpcPath:     "/keyauth.session.UserService/Logout",
				FunctionName: "Logout",
			},
			{
				GrpcPath:         "/keyauth.session.UserService/DescribeSession",
				FunctionName:     "DescribeSession",
				Path:             "/sessions",
				Method:           "GET",
				Resource:         "session",
				AuthEnable:       true,
				PermissionEnable: false,
				Labels:           map[string]string{"allow": "audit_admin"},
			},
			{
				GrpcPath:         "/keyauth.session.UserService/QuerySession",
				FunctionName:     "QuerySession",
				Path:             "/sessions",
				Method:           "GET",
				Resource:         "session",
				AuthEnable:       true,
				PermissionEnable: false,
				Labels:           map[string]string{"allow": "audit_admin"},
			},
			{
				GrpcPath:     "/keyauth.session.AdminService/QueryUserLastSession",
				FunctionName: "QueryUserLastSession",
			},
		},
	}
	return set
}
