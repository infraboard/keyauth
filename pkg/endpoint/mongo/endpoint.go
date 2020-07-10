package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"

	"github.com/infraboard/keyauth/pkg/endpoint"
)

func (s *service) QueryEndpoints(req *endpoint.QueryEndpointRequest) (
	*endpoint.Set, error) {
	return nil, nil
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
