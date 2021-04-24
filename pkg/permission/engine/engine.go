package engine

import (
	"errors"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pb/http"

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
	log      logger.Logger
	policy   policy.PolicyServiceServer
	role     role.RoleServiceServer
	endpoint endpoint.EndpointServiceServer

	permission.UnimplementedPermissionServiceServer
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
	s.log = zap.L().Named("Permission")
	return nil
}

// HttpEntry todo
func (s *service) HTTPEntry() *http.EntrySet {
	return permission.HttpEntry()
}

func init() {
	pkg.RegistryService("permission", Service)
}
