package mongo

import (
	"context"

	"github.com/infraboard/keyauth/pkg/service"
	"github.com/infraboard/mcube/exception"
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

	set := service.NewApplicationSet(req.PageRequest)
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

func (s *microService) Registry(req *service.RegistryRequest) error {
	return nil
}
