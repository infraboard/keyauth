package protocol

import (
	"fmt"
	"net"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/infraboard/keyauth/app/micro"
	auther "github.com/infraboard/keyauth/common/interceptor/grpc"
	"github.com/infraboard/keyauth/conf"
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
		micro: app.GetGrpcApp(micro.AppName).(micro.ServiceServer),
	}
}

// GRPCService grpc服务
type GRPCService struct {
	svr *grpc.Server
	l   logger.Logger
	c   *conf.Config

	micro micro.ServiceServer
}

// Start 启动GRPC服务
func (s *GRPCService) Start() error {
	// 装载所有GRPC服务
	app.LoadGrpcApp(s.svr)

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
