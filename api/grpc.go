package api

import (
	"fmt"
	"log"
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
)

// NewGRPCService todo
func NewGRPCService() *GRPCService {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		pkg.AuthUnaryServerInterceptor(),
	)))

	return &GRPCService{
		svr: grpcServer,
		l:   zap.L().Named("GRPC Service"),
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

	// 注册服务
	s.l.Info("start registry endpoints ...")
	if err := s.RegistryEndpoints(); err != nil {
		s.l.Warnf("registry endpoints error, %s", err)
	}
	s.l.Debug("service endpoints registry success")

	// 启动HTTP服务
	s.l.Infof("GRPC 开始启动, 监听地址: %s", s.c.App.GRPCAddr())
	lis, err := net.Listen("tcp", s.c.App.GRPCAddr())
	if err != nil {
		log.Fatal(err)
	}
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
	req.ClientId = svr.ClientId
	req.ClientSecret = svr.ClientSecret
	_, err = pkg.Endpoint.Registry(ctx.Context(), req)
	return err
}

// Stop 停止GRPC服务
func (s *GRPCService) Stop() error {
	s.l.Info("start graceful shutdown")

	// 优雅关闭HTTP服务
	s.svr.GracefulStop()

	return nil
}
