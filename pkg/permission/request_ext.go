package permission

import (
	"fmt"
	"net/http"

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
func NewCheckPermissionRequest() *CheckPermissionRequest {
	return &CheckPermissionRequest{
		Page: &request.NewPageRequest(100, 1).PageRequest,
	}
}

// Validate 校验请求合法
func (req *CheckPermissionRequest) Validate() error {
	if req.NamespaceId == "" {
		return fmt.Errorf("namespace required")
	}

	if req.EndpointId == "" || (req.ServiceId == "" && req.Path == "") {
		return fmt.Errorf("endpoint_id or (service_id and path) required when check")
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

// NewCheckPermissionRequestFromHTTP 从HTTP请求中加载分页请求
func NewCheckPermissionRequestFromHTTP(req *http.Request) *CheckPermissionRequest {
	qs := req.URL.Query()
	r := NewCheckPermissionRequest()

	r.ServiceId = qs.Get("service_id")
	r.Path = qs.Get("path")

	return r
}
