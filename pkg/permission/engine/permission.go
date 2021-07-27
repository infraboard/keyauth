package engine

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/permission"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/role"
)

func (s *service) QueryPermission(ctx context.Context, req *permission.QueryPermissionRequest) (
	*role.PermissionSet, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate param error, %s", err)
	}

	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 获取用户的策略列表
	preq := policy.NewQueryPolicyRequest(request.NewPageRequest(100, 1))
	preq.Account = tk.Account
	preq.NamespaceId = req.NamespaceId

	policySet, err := s.policy.QueryPolicy(ctx, preq)
	if err != nil {
		return nil, err
	}

	// 获取用户的角色列表
	rset, err := policySet.GetRoles(ctx, s.role, true)
	if err != nil {
		return nil, err
	}

	return rset.Permissions(), nil
}

func (s *service) QueryRole(ctx context.Context, req *permission.QueryRoleRequest) (
	*role.Set, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate param error, %s", err)
	}

	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 获取用户的策略列表
	preq := policy.NewQueryPolicyRequest(request.NewPageRequest(100, 1))
	preq.Account = tk.Account
	preq.NamespaceId = req.NamespaceId

	policySet, err := s.policy.QueryPolicy(ctx, preq)
	if err != nil {
		return nil, err
	}

	return policySet.GetRoles(ctx, s.role, req.WithPermission)
}

func (s *service) CheckPermission(ctx context.Context, req *permission.CheckPermissionRequest) (*role.Permission, error) {
	if req.EndpointId == "" {
		req.EndpointId = endpoint.GenHashID(req.ServiceId, req.Path)
	}

	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate param error, %s", err)
	}

	roleReq := permission.NewQueryRoleRequest(req.NamespaceId)
	roleReq.WithPermission = true
	roleSet, err := s.QueryRole(ctx, roleReq)
	if err != nil {
		return nil, err
	}

	ep, err := s.endpoint.DescribeEndpoint(ctx, endpoint.NewDescribeEndpointRequestWithID(req.EndpointId))
	if err != nil {
		return nil, err
	}
	s.log.Debugf("check roles %s has permission access endpoint [%s]", roleSet.RoleNames(), ep.Entry)

	// 不需要鉴权
	if !ep.Entry.PermissionEnable {
		return role.NewDeaultPermission(), nil
	}

	p, ok, err := roleSet.HasPermission(ep)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, exception.NewPermissionDeny("in namespace %s, role %s has no permission access endpoint: %s",
			req.NamespaceId,
			roleSet.RoleNames(),
			ep.Entry.Path,
		)
	}

	return p, nil
}
