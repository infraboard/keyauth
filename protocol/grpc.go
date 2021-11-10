package protocol

import (
	"fmt"
	"net"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	auther "github.com/infraboard/keyauth/common/interceptor/grpc"
	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/mcube/grpc/middleware/recovery"
)

// NewGRPCService todo
func NewGRPCService() *GRPCService {
	log := zap.L().Named("GRPC Service")

	rc := recovery.NewInterceptor(recovery.NewZapRecoveryHandler())
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		rc.UnaryServerInterceptor(),
		auther.GrpcAuthUnaryServerInterceptor(),
	)))

	return &GRPCService{
		svr:   grpcServer,
		l:     log,
		c:     conf.C(),
		micro: app.GetGrpcApp(micro.AppName).(micro.MicroServiceServer),
	}
}

// GRPCService grpc服务
type GRPCService struct {
	svr *grpc.Server
	l   logger.Logger
	c   *conf.Config

	micro micro.MicroServiceServer
}

// Start 启动GRPC服务
func (s *GRPCService) Start() error {
	// 装载所有GRPC服务
	if err := app.LoadGrpcApp(s.svr); err != nil {
		return err
	}

	// 加载内部服务
	if err := app.LoadInternalApp(); err != nil {
		return err
	}

	s.l.Infof("loaded grpc service: %v", app.LoadedGrpcApp())
	s.l.Infof("loaded internal service: %v", app.LoadedInternalApp())

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

// Stop 停止GRPC服务
func (s *GRPCService) Stop() error {
	s.l.Info("start grpc graceful shutdown ...")
	// 优雅关闭HTTP服务
	s.svr.GracefulStop()
	return nil
}
