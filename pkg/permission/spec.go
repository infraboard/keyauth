package permission

import (
	"fmt"

	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/mcube/http/request"
)

// Service 权限查询API
type Service interface {
	QueryPermission(req *QueryPermissionRequest) (*role.PermissionSet, error)
	CheckPermission(req *CheckPermissionrequest) (*endpoint.Endpoint, error)
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
	return &CheckPermissionrequest{
		Session: token.NewSession(),
	}
}

// CheckPermissionrequest todo
type CheckPermissionrequest struct {
	*token.Session
	NamespaceID string
	EnpointID   string
}

// Validate 校验请求合法
func (req *CheckPermissionrequest) Validate() error {
	tk := req.GetToken()
	if tk == nil {
		return fmt.Errorf("token required")
	}

	if req.NamespaceID == "" || req.EnpointID == "" {
		return fmt.Errorf("namespace_id and endpoint_id required")
	}

	return nil
}
