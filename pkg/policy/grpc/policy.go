package grpc

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/namespace"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/role"
)

func (s *service) CreatePolicy(ctx context.Context, req *policy.CreatePolicyRequest) (
	*policy.Policy, error) {
	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	ins, err := policy.New(tk, req)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	u, err := ins.CheckDependence(ctx, s.user, s.role, s.namespace)
	if err != nil {
		return nil, err
	}
	ins.UserType = u.Type

	if _, err := s.col.InsertOne(context.TODO(), ins); err != nil {
		return nil, exception.NewInternalServerError("inserted policy(%s) document error, %s",
			ins.Id, err)
	}

	return ins, nil
}

func (s *service) QueryPolicy(ctx context.Context, req *policy.QueryPolicyRequest) (
	*policy.Set, error) {
	r, err := newQueryPolicyRequest(req)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("query policy filter: %s", r.FindFilter())
	resp, err := s.col.Find(context.TODO(), r.FindFilter(), r.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find policy error, error is %s", err)
	}

	set := policy.NewPolicySet()
	// 循环
	for resp.Next(context.TODO()) {
		ins := policy.NewDefaultPolicy()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode policy error, error is %s", err)
		}

		// 补充关联的角色信息
		if req.WithRole {
			descRole := role.NewDescribeRoleRequestWithID(ins.RoleId)
			ins.Role, err = s.role.DescribeRole(ctx, descRole)
			if err != nil {
				return nil, err
			}
		}

		// 关联空间信息
		if req.WithNamespace && ins.NamespaceId != "" && ins.NamespaceId != "*" {
			descNS := namespace.NewNewDescriptNamespaceRequestWithID(ins.NamespaceId)
			ins.Namespace, err = s.namespace.DescribeNamespace(ctx, descNS)
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

func (s *service) DescribePolicy(ctx context.Context, req *policy.DescribePolicyRequest) (
	*policy.Policy, error) {
	r, err := newDescribePolicyRequest(req)
	if err != nil {
		return nil, err
	}

	ins := policy.NewDefaultPolicy()
	s.log.Debugf("describe policy filter: %s", r.FindFilter())
	if err := s.col.FindOne(context.TODO(), r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("policy %s not found", req)
		}

		return nil, exception.NewInternalServerError("find policy %s error, %s", req.Id, err)
	}

	return ins, nil
}

func (s *service) DeletePolicy(ctx context.Context, req *policy.DeletePolicyRequest) (*policy.Policy, error) {
	descReq := policy.NewDescriptPolicyRequest()
	descReq.Id = req.Id
	p, err := s.DescribePolicy(ctx, descReq)
	if err != nil {
		return nil, err
	}

	r, err := newDeletePolicyRequest(req)
	if err != nil {
		return nil, err
	}

	_, err = s.col.DeleteOne(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("delete policy(%s) error, %s", req.Id, err)
	}

	return p, nil
}
