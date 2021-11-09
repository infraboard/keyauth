package impl

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/permission"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/role"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	log      logger.Logger
	policy   policy.PolicyServiceServer
	role     role.RoleServiceServer
	endpoint endpoint.EndpointServiceServer

	permission.UnimplementedPermissionServiceServer
}

func (s *service) Config() error {
	s.policy = app.GetGrpcApp(policy.AppName).(policy.PolicyServiceServer)
	s.role = app.GetGrpcApp(role.AppName).(role.RoleServiceServer)
	s.endpoint = app.GetGrpcApp(endpoint.AppName).(endpoint.EndpointServiceServer)
	s.log = zap.L().Named("Permission")
	return nil
}

func (s *service) Name() string {
	return permission.AppName
}

func (s *service) Registry(server *grpc.Server) {
	permission.RegisterPermissionServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
