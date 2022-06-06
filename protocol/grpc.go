package protocol

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/client"
	"github.com/infraboard/mcenter/client/lifecycle"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/keyauth/apps/micro"
	auther "github.com/infraboard/keyauth/common/interceptor/grpc"
	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/mcube/grpc/middleware/recovery"
)

// NewGRPCService todo
func NewGRPCService() *GRPCService {
	log := zap.L().Named("GRPC Service")

	rc := recovery.NewInterceptor(recovery.NewZapRecoveryHandler())
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		rc.UnaryServerInterceptor(),
		auther.GrpcAuthUnaryServerInterceptor(),
	))

	ctx, cancel := context.WithCancel(context.Background())
	return &GRPCService{
		ctx:    ctx,
		cancel: cancel,
		svr:    grpcServer,
		l:      log,
		c:      conf.C(),
		micro:  app.GetGrpcApp(micro.AppName).(micro.ServiceServer),
	}
}

// GRPCService grpc服务
type GRPCService struct {
	svr    *grpc.Server
	l      logger.Logger
	c      *conf.Config
	ctx    context.Context
	cancel context.CancelFunc
	lf     lifecycle.Lifecycler

	micro micro.ServiceServer
}

// 注册
func (s *GRPCService) registry() {
	// 提前加载好 mcenter客户端
	err := client.LoadClientFromConfig(s.c.Mcenter)
	if err != nil {
		s.l.Errorf("load mcenter client from config error, %s", err)
		return
	}

	// 注册服务实例
	req := instance.NewRegistryRequest()
	req.Address = s.c.App.GRPCAddr()
	lf, err := client.C().Registry(s.ctx, req)
	if err != nil {
		s.l.Errorf("registry to mcenter error, %s", err)
		return
	}
	s.lf = lf

	// 上报实例心跳
	lf.Heartbeat(s.ctx)
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

	time.AfterFunc(1*time.Second, s.registry)

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
	// 取消
	s.cancel()

	s.l.Info("start grpc graceful shutdown ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 注销服务实例
	if s.lf != nil {
		if err := s.lf.UnRegistry(ctx); err != nil {
			s.l.Errorf("unregistry error, %s", err)
		}
	}

	// 优雅关闭HTTP服务
	s.svr.GracefulStop()
	return nil
}
