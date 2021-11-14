package impl

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/keyauth/app/endpoint"
	"github.com/infraboard/keyauth/app/permission"
	"github.com/infraboard/keyauth/app/policy"
	"github.com/infraboard/keyauth/app/role"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	log      logger.Logger
	policy   policy.ServiceServer
	role     role.ServiceServer
	endpoint endpoint.ServiceServer

	permission.UnimplementedServiceServer
}

func (s *service) Config() error {
	s.policy = app.GetGrpcApp(policy.AppName).(policy.ServiceServer)
	s.role = app.GetGrpcApp(role.AppName).(role.ServiceServer)
	s.endpoint = app.GetGrpcApp(endpoint.AppName).(endpoint.ServiceServer)
	s.log = zap.L().Named("Permission")
	return nil
}

func (s *service) Name() string {
	return permission.AppName
}

func (s *service) Registry(server *grpc.Server) {
	permission.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
