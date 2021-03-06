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
	rset, err := policySet.GetRoles(ctx, s.role)
	if err != nil {
		return nil, err
	}

	return rset.Permissions(), nil
}

func (s *service) QueryRoles(ctx context.Context, req *permission.QueryRoleRequest) (
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

	return policySet.GetRoles(ctx, s.role)
}

func (s *service) CheckPermission(ctx context.Context, req *permission.CheckPermissionRequest) (*role.Permission, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate param error, %s", err)
	}

	roleSet, err := s.QueryRoles(ctx, permission.NewQueryRoleRequest(req.NamespaceId))
	if err != nil {
		return nil, err
	}

	ep, err := s.endpoint.DescribeEndpoint(ctx, endpoint.NewDescribeEndpointRequestWithID(req.EndpointId))
	if err != nil {
		return nil, err
	}

	p, ok, err := roleSet.HasPermission(ep)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, exception.NewNotFound("not perm for this enpind")
	}

	return p, nil
}
