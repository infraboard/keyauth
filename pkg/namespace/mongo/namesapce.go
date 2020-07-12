package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg/namespace"
)

func (s *service) CreateNamespace(req *namespace.CreateNamespaceRequest) (
	*namespace.Namespace, error) {
	ins, err := namespace.NewNamespace(req)
	if err != nil {
		return nil, err
	}

	if _, err := s.col.InsertOne(context.TODO(), ins); err != nil {
		return nil, exception.NewInternalServerError("inserted namespace(%s) document error, %s",
			ins.Name, err)
	}

	return ins, nil
}

func (s *service) QueryNamespace(req *namespace.QueryNamespaceRequest) (
	*namespace.Set, error) {
	r := newPaggingQuery(req)
	resp, err := s.col.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find namespace error, error is %s", err)
	}

	set := namespace.NewNamespaceSet(req.PageRequest)
	// 循环
	for resp.Next(context.TODO()) {
		ins := namespace.NewDefaultNamespace()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode namespace error, error is %s", err)
		}

		set.Add(ins)
	}

	// count
	count, err := s.col.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get namespace count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (s *service) DescribeNamespace(req *namespace.DescriptNamespaceRequest) (
	*namespace.Namespace, error) {
	r, err := newDescribeQuery(req)
	if err != nil {
		return nil, err
	}

	ins := namespace.NewDefaultNamespace()
	if err := s.col.FindOne(context.TODO(), r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("namespace %s not found", req)
		}

		return nil, exception.NewInternalServerError("find namespace %s error, %s", req.ID, err)
	}

	return ins, nil
}

func (s *service) DeleteNamespace(req *namespace.DeleteNamespaceRequest) error {
	return nil
}
