package mongo

import (
	"context"

	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/mcube/http/request"
)

const (
	// MaxQueryEndpoints todo
	MaxQueryEndpoints = 500
)

func (s *service) QueryResources(ctx context.Context, req *endpoint.QueryResourceRequest) (
	*endpoint.ResourceSet, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	rs := endpoint.NewResourceSet()
	queryE := endpoint.NewQueryEndpointRequest(request.NewPageRequest(MaxQueryEndpoints, 1))
	for _, id := range req.ServiceIds {
		queryE.ServiceId = id
		eps, err := s.QueryEndpoints(ctx, queryE)
		if err != nil {
			return nil, err
		}
		if eps.Total > MaxQueryEndpoints {
			s.log.Warnf("service %s total endpoints > %d", id, eps.Total)
		}

		rs.AddEndpointSet(eps)
	}

	return rs, nil
}
