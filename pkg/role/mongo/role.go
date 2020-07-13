package mongo

import (
	"context"

	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) CreateRole(t role.Type, req *role.CreateRoleRequest) (*role.Role, error) {
	r, err := role.New(t, req)
	if err != nil {
		return nil, err
	}

	if _, err := s.col.InsertOne(context.TODO(), r); err != nil {
		return nil, exception.NewInternalServerError("inserted role(%s) document error, %s",
			r.Name, err)
	}

	return r, nil
}

func (s *service) QueryRole(req *role.QueryRoleRequest) (*role.Set, error) {
	query, err := newQueryRoleRequest(req)
	if err != nil {
		return nil, err
	}

	resp, err := s.col.Find(context.TODO(), query.FindFilter(), query.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find role error, error is %s", err)
	}

	set := role.NewRoleSet(req.PageRequest)
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

func (s *service) DescribeRole(req *role.DescribeRoleRequest) (*role.Role, error) {
	query, err := newDescribeRoleRequest(req)
	if err != nil {
		return nil, err
	}

	ins := role.NewDefaultRole()
	if err := s.col.FindOne(context.TODO(), query.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("role %s not found", req)
		}

		return nil, exception.NewInternalServerError("find role %s error, %s", req, err)
	}

	return ins, nil
}

func (s *service) DeleteRole(id string) error {
	resp, err := s.col.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return exception.NewInternalServerError("delete role(%s) error, %s", id, err)
	}

	if resp.DeletedCount == 0 {
		return exception.NewNotFound("role(%s) not found", id)
	}

	return nil
}
