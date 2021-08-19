package grpc

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/role"
)

func (s *service) CreateRole(ctx context.Context, req *role.CreateRoleRequest) (*role.Role, error) {
	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	r, err := role.New(tk, req)
	if err != nil {
		return nil, err
	}

	if _, err := s.col.InsertOne(context.TODO(), r); err != nil {
		return nil, exception.NewInternalServerError("inserted role(%s) document error, %s",
			r.Name, err)
	}

	// 添加权限条目
	permReq := role.NewAddPermissionToRoleRequest()
	permReq.Permissions = req.Permissions
	permReq.RoleId = r.Id
	ps, err := s.AddPermissionToRole(ctx, permReq)
	if err != nil {
		return nil, err
	}
	r.Permissions = ps.Items
	return r, nil
}

func (s *service) QueryRole(ctx context.Context, req *role.QueryRoleRequest) (*role.Set, error) {
	query, err := newQueryRoleRequest(req)
	if err != nil {
		return nil, err
	}

	resp, err := s.col.Find(context.TODO(), query.FindFilter(), query.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find role error, error is %s", err)
	}

	set := role.NewRoleSet()
	// 循环
	for resp.Next(context.TODO()) {
		ins := role.NewDefaultRole()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode role error, error is %s", err)
		}
		set.Add(ins)
	}

	// count
	count, err := s.col.CountDocuments(context.TODO(), query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get token count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (s *service) DescribeRole(ctx context.Context, req *role.DescribeRoleRequest) (*role.Role, error) {
	query, err := newDescribeRoleRequest(req)
	if err != nil {
		return nil, err
	}

	ins := role.NewDefaultRole()
	if err := s.col.FindOne(context.TODO(), query.FindFilter(), query.FindOptions()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("role %s not found", req)
		}

		return nil, exception.NewInternalServerError("find role %s error, %s", req, err)
	}

	if req.WithPermissions {
		queryPerm := role.NewQueryPermissionRequest(request.NewPageRequest(role.RoleMaxPermission, 1))
		queryPerm.RoleId = ins.Id
		ps, err := s.QueryPermission(ctx, queryPerm)
		if err != nil {
			return nil, err
		}
		ins.Permissions = ps.Items
	}

	return ins, nil
}

func (s *service) DeleteRole(ctx context.Context, req *role.DeleteRoleRequest) (*role.Role, error) {
	r, err := s.DescribeRole(ctx, role.NewDescribeRoleRequestWithID(req.Id))
	if err != nil {
		return nil, err
	}

	if r.Type.Equal(role.RoleType_BUILDIN) {
		return nil, fmt.Errorf("build_in role can't be delete")
	}

	if !req.DeletePolicy {
		queryReq := policy.NewQueryPolicyRequest(request.NewPageRequest(20, 1))
		queryReq.RoleId = req.Id
		ps, err := s.policy.QueryPolicy(ctx, queryReq)
		if err != nil {
			return nil, err
		}
		if ps.Total > 0 {
			return nil, exception.NewBadRequest("该角色还关联得有策略, 请先删除关联策略")
		}
	}

	resp, err := s.col.DeleteOne(context.TODO(), bson.M{"_id": req.Id})
	if err != nil {
		return nil, exception.NewInternalServerError("delete role(%s) error, %s", req.Id, err)
	}

	if resp.DeletedCount == 0 {
		return nil, exception.NewNotFound("role(%s) not found", req.Id)
	}

	// 清除角色关联的权限
	permReq := role.NewRemovePermissionFromRoleRequest()
	permReq.RoleId = req.Id
	permReq.RemoveAll = true
	_, err = s.RemovePermissionFromRole(ctx, permReq)
	if err != nil {
		s.log.Errorf("delete role permission error, %s", err)
	}

	// 清除角色关联的策略
	_, err = s.policy.DeletePolicy(ctx, policy.NewDeletePolicyRequestWithRoleID(req.Id))
	if err != nil {
		s.log.Errorf("delete role policy error, %s", err)
	}

	return r, nil
}
