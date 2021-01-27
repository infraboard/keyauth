package mongo

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/common/session"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/role"
)

func (s *service) CreateRole(ctx context.Context, req *role.CreateRoleRequest) (*role.Role, error) {
	tk := session.GetTokenFromContext(ctx)

	r, err := role.New(tk, req)
	if err != nil {
		return nil, err
	}

	if _, err := s.col.InsertOne(context.TODO(), r); err != nil {
		return nil, exception.NewInternalServerError("inserted role(%s) document error, %s",
			r.Data.Name, err)
	}

	return r, nil
}

func (s *service) QueryRole(ctx context.Context, req *role.QueryRoleRequest) (*role.Set, error) {
	tk := session.GetTokenFromContext(ctx)

	query, err := newQueryRoleRequest(tk, req)
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

	return ins, nil
}

func (s *service) DeleteRole(ctx context.Context, req *role.DeleteRoleRequest) (*role.Role, error) {
	r, err := s.DescribeRole(ctx, role.NewDescribeRoleRequestWithID(req.Id))
	if err != nil {
		return nil, err
	}

	if r.Data.Type.Equal(role.RoleType_BUILDIN) {
		return nil, fmt.Errorf("build_in role can't be delete")
	}

	resp, err := s.col.DeleteOne(context.TODO(), bson.M{"_id": req.Id})
	if err != nil {
		return nil, exception.NewInternalServerError("delete role(%s) error, %s", req.Id, err)
	}

	if resp.DeletedCount == 0 {
		return nil, exception.NewNotFound("role(%s) not found", req.Id)
	}

	// 清除角色管理的策略
	_, err = s.policy.DeletePolicy(ctx, policy.NewDeletePolicyRequestWithRoleID(req.Id))
	if err != nil {
		s.log.Errorf("delete role policy error, %s", err)
	}

	return r, nil
}
