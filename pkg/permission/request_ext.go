package permission

import (
	"fmt"

	"github.com/infraboard/mcube/http/request"
)

// NewQueryPermissionRequest todo
func NewQueryPermissionRequest(page *request.PageRequest) *QueryPermissionRequest {
	return &QueryPermissionRequest{
		Page: &page.PageRequest,
	}
}

// Validate 校验请求合法
func (req *QueryPermissionRequest) Validate() error {
	if req.NamespaceId == "" {
		return fmt.Errorf("namespace required")
	}

	return nil
}

// NewCheckPermissionrequest todo
func NewCheckPermissionrequest() *CheckPermissionrequest {
	return &CheckPermissionrequest{
		Page: &request.NewPageRequest(100, 1).PageRequest,
	}
}

// Validate 校验请求合法
func (req *CheckPermissionrequest) Validate() error {
	if req.NamespaceId == "" {
		return fmt.Errorf("namespace required")
	}

	if req.EndpointId == "" {
		return fmt.Errorf("endpoint_id required when check")
	}

	return nil
}

// NewQueryRoleRequest todo
func NewQueryRoleRequest(namespaceid string) *QueryRoleRequest {
	return &QueryRoleRequest{
		Page:        &request.NewPageRequest(100, 1).PageRequest,
		NamespaceId: namespaceid,
	}
}

// Validate 校验请求合法
func (req *QueryRoleRequest) Validate() error {
	if req.NamespaceId == "" {
		return fmt.Errorf("namespace required")
	}

	return nil
}
