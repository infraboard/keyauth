package api

import (
	"fmt"
	"log"
	"net"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
)

// NewGRPCService todo
func NewGRPCService() *GRPCService {
	grpcServer := grpc.NewServer()

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
	if err := pkg.InitV1GRPCAPI(s.svr); err != nil {
		return err
	}

	// GRPC 服务无效注册Endpoint

	// 启动HTTP服务
	s.l.Infof("GRPC 开始启动, 监听地址: %s ...", s.c.App.GRPCAddr())
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

// Stop 停止GRPC服务
func (s *GRPCService) Stop() error {
	s.l.Info("start graceful shutdown")

	// 优雅关闭HTTP服务
	s.svr.GracefulStop()

	return nil
}
