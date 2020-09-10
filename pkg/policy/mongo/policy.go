package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg/namespace"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/role"
)

func (s *service) CreatePolicy(t policy.Type, req *policy.CreatePolicyRequest) (
	*policy.Policy, error) {
	ins, err := policy.New(t, req)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	u, err := ins.CheckDependence(s.user, s.role, s.namespace)
	if err != nil {
		return nil, err
	}
	ins.UserType = u.Type

	if _, err := s.col.InsertOne(context.TODO(), ins); err != nil {
		return nil, exception.NewInternalServerError("inserted policy(%s) document error, %s",
			ins.ID, err)
	}

	return ins, nil
}

func (s *service) QueryPolicy(req *policy.QueryPolicyRequest) (
	*policy.Set, error) {
	r, err := newQueryPolicyRequest(req)
	if err != nil {
		return nil, err
	}

	resp, err := s.col.Find(context.TODO(), r.FindFilter(), r.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find policy error, error is %s", err)
	}

	set := policy.NewPolicySet(req.PageRequest)
	// 循环
	for resp.Next(context.TODO()) {
		ins := policy.NewDefaultPolicy()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode policy error, error is %s", err)
		}

		// 补充关联的角色信息
		if req.WithRole {
			descRole := role.NewDescribeRoleRequestWithID(ins.RoleID)
			ins.Role, err = s.role.DescribeRole(descRole)
			if err != nil {
				return nil, err
			}
		}

		// 关联空间信息
		if req.WithNamespace {
			descNS := namespace.NewNewDescriptNamespaceRequestWithID(ins.NamespaceID)
			ins.Namespace, err = s.namespace.DescribeNamespace(descNS)
			if err != nil {
				return nil, err
			}
		}

		set.Add(ins)
	}

	// count
	count, err := s.col.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get policy count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (s *service) DescribePolicy(req *policy.DescribePolicyRequest) (
	*policy.Policy, error) {
	r, err := newDescribePolicyRequest(req)
	if err != nil {
		return nil, err
	}

	ins := policy.NewDefaultPolicy()
	if err := s.col.FindOne(context.TODO(), r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("policy %s not found", req)
		}

		return nil, exception.NewInternalServerError("find policy %s error, %s", req.ID, err)
	}

	return ins, nil
}
