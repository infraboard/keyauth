package engine

import (
	"errors"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/permission"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/role"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	policy   policy.Service
	role     role.Service
	endpoint endpoint.Service
}

func (s *service) Config() error {
	if pkg.Policy == nil {
		return errors.New("denpence policy service is nil")
	}
	s.policy = pkg.Policy

	if pkg.Role == nil {
		return errors.New("denpence role service is nil")
	}
	s.role = pkg.Role

	if pkg.Endpoint == nil {
		return errors.New("denpence endpoint service is nil")
	}
	s.endpoint = pkg.Endpoint

	return nil
}

func init() {
	var _ permission.Service = Service
	pkg.RegistryService("permission", Service)
}
