package engine

import (
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/permission"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/mcube/http/request"
)

func (s *service) QueryPermission(req *permission.QueryPermissionRequest) (
	*role.PermissionSet, error) {
	tk := req.GetToken()

	// 获取用户的策略列表
	preq := policy.NewQueryPolicyRequest(request.NewPageRequest(100, 1))
	preq.UserID = tk.UserID
	preq.NamespaceID = req.NamespaceID
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

func (s *service) CheckPermission(req *permission.CheckPermissionrequest) (*endpoint.Endpoint, error) {
	return nil, nil
}
