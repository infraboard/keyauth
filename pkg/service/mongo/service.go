package mongo

import (
	"context"

	"github.com/infraboard/keyauth/pkg/service"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *microService) CreateService(req *service.CreateServiceRequest) (
	*service.MicroService, error) {
	ins, err := service.New(req)
	if err != nil {
		return nil, err
	}

	if _, err := s.col.InsertOne(context.TODO(), ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a service document error, %s", err)
	}
	return ins, nil
}

func (s *microService) QueryService(req *service.QueryServiceRequest) (*service.MicroServiceSet, error) {
	r := newPaggingQuery(req)
	resp, err := s.col.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find service error, error is %s", err)
	}

	set := service.NewMicroServiceSet(req.PageRequest)
	// 循环
	for resp.Next(context.TODO()) {
		ins := new(service.MicroService)
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode service error, error is %s", err)
		}

		set.Add(ins)
	}

	// count
	count, err := s.col.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get device count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (s *microService) DescribeService(req *service.DescriptServiceRequest) (
	*service.MicroService, error) {
	r, err := newDescribeQuery(req)
	if err != nil {
		return nil, err
	}

	ins := new(service.MicroService)
	if err := s.col.FindOne(context.TODO(), r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("service %s not found", req)
		}

		return nil, exception.NewInternalServerError("find service %s error, %s", req, err)
	}
	return ins, nil
}

func (s *microService) DeleteService(name string) error {
	describeReq := service.NewDescriptServiceRequest()
	describeReq.Name = name
	if _, err := s.DescribeService(describeReq); err != nil {
		return err
	}

	_, err := s.col.DeleteOne(context.TODO(), bson.M{"_id": name})
	if err != nil {
		return exception.NewInternalServerError("delete service(%s) error, %s", name, err)
	}
	return nil
}

func (s *microService) Registry(req *service.RegistryRequest) error {
	descReq := service.NewDescriptServiceRequest()
	descReq.ServiceID = req.ServiceID
	svr, err := s.DescribeService(descReq)
	if err != nil {
		return err
	}
	if !svr.CheckKey(req.ServiceKey) {
		return exception.NewUnauthorized("服务凭证不正确")
	}
	return nil
}
