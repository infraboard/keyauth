package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg/policy"
)

func (s *service) CreatePolicy(req *policy.CreatePolicyRequest) (
	*policy.Policy, error) {
	ins, err := policy.New(req)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	if err := ins.CheckDependence(s.user, s.role, s.namespace); err != nil {
		return nil, err
	}

	if _, err := s.col.InsertOne(context.TODO(), ins); err != nil {
		return nil, exception.NewInternalServerError("inserted policy(%s) document error, %s",
			ins.ID, err)
	}

	return ins, nil
}

func (s *service) QueryPolicy(req *policy.QueryPolicyRequest) (
	*policy.Set, error) {
	r := newQueryRequest(req)
	resp, err := s.col.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find policy error, error is %s", err)
	}

	set := policy.NewPolicySet(req.PageRequest)
	// 循环
	for resp.Next(context.TODO()) {
		app := policy.NewDefaultPolicy()
		if err := resp.Decode(app); err != nil {
			return nil, exception.NewInternalServerError("decode policy error, error is %s", err)
		}

		set.Add(app)
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
