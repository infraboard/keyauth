package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"

	"github.com/infraboard/keyauth/pkg/endpoint"
)

func (s *service) QueryEndpoints(req *endpoint.QueryEndpointRequest) (
	*endpoint.Set, error) {
	r := newQueryEndpointRequest(req)
	resp, err := s.col.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find endpoint error, error is %s", err)
	}

	set := endpoint.NewEndpointSet(req.PageRequest)
	// 循环
	for resp.Next(context.TODO()) {
		app := endpoint.NewDefaultEndpoint()
		if err := resp.Decode(app); err != nil {
			return nil, exception.NewInternalServerError("decode domain error, error is %s", err)
		}

		set.Add(app)
	}

	// count
	count, err := s.col.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get device count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (s *service) Registry(req *endpoint.RegistryRequest) error {
	if err := req.Validate(); err != nil {
		return exception.NewBadRequest(err.Error())
	}

	endpoints := req.Endpoints()
	many := make([]interface{}, 0, len(endpoints))
	for i := range endpoints {
		many = append(many, endpoints[i])
	}

	if _, err := s.col.InsertMany(context.TODO(), many); err != nil {
		return exception.NewInternalServerError("inserted a service document error, %s", err)
	}

	return nil
}
