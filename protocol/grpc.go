package protocol

import (
	"fmt"
	"net"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/version"
	"github.com/infraboard/mcube/grpc/middleware/recovery"
)

// NewGRPCService todo
func NewGRPCService() *GRPCService {
	log := zap.L().Named("GRPC Service")

	rc := recovery.NewInterceptor(recovery.NewZapRecoveryHandler())
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		rc.UnaryServerInterceptor(),
		pkg.AuthUnaryServerInterceptor(),
	)))

	return &GRPCService{
		svr: grpcServer,
		l:   log,
		c:   conf.C(),
	}
}

// GRPCService grpc服务
type GRPCService struct {
	svr *grpc.Server
	l   logger.Logger
	c   *conf.Config
}

// Start 启动GRPC服务
func (s *GRPCService) Start() error {
	// 装载所有GRPC服务
	pkg.InitV1GRPCAPI(s.svr)

	// 启动HTTP服务
	lis, err := net.Listen("tcp", s.c.App.GRPCAddr())
	if err != nil {
		return err
	}

	s.l.Infof("GRPC 服务监听地址: %s", s.c.App.GRPCAddr())
	if err := s.svr.Serve(lis); err != nil {
		if err == grpc.ErrServerStopped {
			s.l.Info("service is stopped")
		}

		return fmt.Errorf("start service error, %s", err.Error())
	}

	return nil
}

// RegistryEndpoints 注册条目
func (s *GRPCService) RegistryEndpoints() error {
	if pkg.Micro == nil {
		return fmt.Errorf("dependence micro service is nil")
	}

	ctx := pkg.NewInternalMockGrpcCtx("internal")

	desc := micro.NewDescribeServiceRequest()
	desc.Name = version.ServiceName
	svr, err := pkg.Micro.DescribeService(ctx.Context(), desc)
	if err != nil {
		return err
	}

	if pkg.Endpoint == nil {
		return fmt.Errorf("dependence endpoint service is nil")
	}

	req := endpoint.NewRegistryRequest(version.Short(), pkg.HTTPEntry().Items)
	ctx.SetClientCredentials(svr.ClientId, svr.ClientSecret)
	_, err = pkg.Endpoint.Registry(ctx.Context(), req)
	return err
}

// Stop 停止GRPC服务
func (s *GRPCService) Stop() error {
	s.l.Info("start grpc graceful shutdown ...")
	// 优雅关闭HTTP服务
	s.svr.GracefulStop()
	return nil
}
