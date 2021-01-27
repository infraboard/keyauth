package engine

import (
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"

	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/permission"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/role"
)

func (s *service) QueryPermission(req *permission.QueryPermissionRequest) (
	*role.PermissionSet, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate param error, %s", err)
	}

	tk := req.GetToken()

	// 获取用户的策略列表
	preq := policy.NewQueryPolicyRequest(request.NewPageRequest(100, 1))
	preq.Account = tk.Account
	preq.NamespaceId = req.NamespaceID

	policySet, err := s.policy.QueryPolicy(preq)
	if err != nil {
		return nil, err
	}

	// 获取用户的角色列表
	rset, err := policySet.GetRoles(s.role)
	if err != nil {
		return nil, err
	}

	return rset.Permissions(), nil
}

func (s *service) QueryRoles(req *permission.QueryPermissionRequest) (
	*role.Set, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate param error, %s", err)
	}

	tk := req.GetToken()

	// 获取用户的策略列表
	preq := policy.NewQueryPolicyRequest(request.NewPageRequest(100, 1))
	preq.Account = tk.Account
	preq.NamespaceId = req.NamespaceID

	policySet, err := s.policy.QueryPolicy(preq)
	if err != nil {
		return nil, err
	}

	return policySet.GetRoles(s.role)
}

func (s *service) CheckPermission(req *permission.CheckPermissionrequest) (*role.Permission, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate param error, %s", err)
	}

	rset, err := s.QueryRoles(req.QueryPermissionRequest)
	if err != nil {
		return nil, err
	}

	ep, err := s.endpoint.DescribeEndpoint(endpoint.NewDescribeEndpointRequestWithID(req.EnpointID))
	if err != nil {
		return nil, err
	}

	p, ok, err := rset.HasPermission(ep)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, exception.NewNotFound("not perm for this enpind")
	}

	return p, nil
}
