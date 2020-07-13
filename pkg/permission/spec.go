package permission

import (
	"fmt"

	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/mcube/http/request"
)

// Service 权限查询API
type Service interface {
	QueryPermission(req *QueryPermissionRequest) (*role.PermissionSet, error)
	QueryRoles(req *QueryPermissionRequest) (*role.Set, error)
	CheckPermission(req *CheckPermissionrequest) (*role.Permission, error)
}

// NewQueryPermissionRequest todo
func NewQueryPermissionRequest(page *request.PageRequest) *QueryPermissionRequest {
	return &QueryPermissionRequest{
		PageRequest: page,
		Session:     token.NewSession(),
	}
}

// QueryPermissionRequest 查询用户权限
type QueryPermissionRequest struct {
	*token.Session
	*request.PageRequest
	NamespaceID string
}

// Validate 校验请求合法
func (req *QueryPermissionRequest) Validate() error {
	tk := req.GetToken()
	if tk == nil {
		return fmt.Errorf("token required")
	}

	if req.NamespaceID == "" {
		return fmt.Errorf("namespace required")
	}

	return nil
}

// NewCheckPermissionrequest todo
func NewCheckPermissionrequest() *CheckPermissionrequest {
	query := NewQueryPermissionRequest(request.NewPageRequest(100, 1))
	return &CheckPermissionrequest{
		QueryPermissionRequest: query,
	}
}

// CheckPermissionrequest todo
type CheckPermissionrequest struct {
	*QueryPermissionRequest
	EnpointID string
}

// Validate 校验请求合法
func (req *CheckPermissionrequest) Validate() error {
	if err := req.QueryPermissionRequest.Validate(); err != nil {
		return err
	}

	if req.EnpointID == "" {
		return fmt.Errorf("endpoint_id required when check")
	}

	return nil
}
