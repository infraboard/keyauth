// Code generated by protoc-gen-go-http. DO NOT EDIT.

package role

import (
	http "github.com/infraboard/mcube/pb/http"
)

// HttpEntry todo
func HttpEntry() *http.EntrySet {
	set := &http.EntrySet{
		Items: []*http.Entry{
			{
				GrpcPath:         "/keyauth.role.RoleService/CreateRole",
				FunctionName:     "CreateRole",
				Path:             "/policies",
				Method:           "POST",
				Resource:         "policy",
				AuthEnable:       true,
				PermissionEnable: false,
				Labels:           map[string]string{"allow": "perm_admin"},
			},
			{
				GrpcPath:         "/keyauth.role.RoleService/QueryRole",
				FunctionName:     "QueryRole",
				Path:             "/policies",
				Method:           "POST",
				Resource:         "policy",
				AuthEnable:       true,
				PermissionEnable: false,
				Labels:           map[string]string{"allow": "perm_admin"},
			},
			{
				GrpcPath:         "/keyauth.role.RoleService/DescribeRole",
				FunctionName:     "DescribeRole",
				Path:             "/policies",
				Method:           "POST",
				Resource:         "policy",
				AuthEnable:       true,
				PermissionEnable: false,
				Labels:           map[string]string{"allow": "perm_admin"},
			},
			{
				GrpcPath:         "/keyauth.role.RoleService/DeleteRole",
				FunctionName:     "DeleteRole",
				Path:             "/policies",
				Method:           "POST",
				Resource:         "policy",
				AuthEnable:       true,
				PermissionEnable: false,
				Labels:           map[string]string{"allow": "perm_admin"},
			},
		},
	}
	return set
}
