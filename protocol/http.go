package protocol

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/infraboard/mcube/http/middleware/accesslog"
	"github.com/infraboard/mcube/http/middleware/cors"
	"github.com/infraboard/mcube/http/middleware/recovery"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/http/router/httprouter"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/common/auther"
	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/version"
)

// NewHTTPService 构建函数
func NewHTTPService() *HTTPService {
	r := httprouter.New()
	r.Use(recovery.NewWithLogger(zap.L().Named("Recovery")))
	r.Use(accesslog.NewWithLogger(zap.L().Named("AccessLog")))
	r.Use(cors.AllowAll())
	r.EnableAPIRoot()
	r.SetAuther(auther.NewHTTPAuther())
	r.Auth(true)
	r.Permission(true)

	server := &http.Server{
		ReadHeaderTimeout: 20 * time.Second,
		// ReadTimeout:       20 * time.Second,
		// WriteTimeout:      25 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Addr:           conf.C().App.HTTPAddr(),
		Handler:        r,
	}

	return &HTTPService{
		r:      r,
		server: server,
		l:      zap.L().Named("HTTP Service"),
		c:      conf.C(),
	}
}

// HTTPService http服务
type HTTPService struct {
	r      router.Router
	l      logger.Logger
	c      *conf.Config
	server *http.Server
}

// Start 启动服务
func (s *HTTPService) Start() error {
	// 初始化GRPC客户端
	if err := s.initGRPCClient(); err != nil {
		return err
	}

	// 装置子服务路由
	if err := pkg.InitV1HTTPAPI(s.c.App.Name, s.r); err != nil {
		return err
	}

	// 启动HTTP服务
	s.l.Infof("HTTP 服务开始启动, 监听地址: %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.l.Info("service is stopped")
		}
		return fmt.Errorf("start service error, %s", err.Error())
	}
	return nil
}

// Stop 停止server
func (s *HTTPService) Stop() error {
	s.l.Info("start http graceful shutdown ...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 优雅关闭HTTP服务
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Errorf("graceful shutdown timeout, force exit")
	}

	return nil
}

// InitGRPCClient 初始化grpc客户端
func (s *HTTPService) initGRPCClient() error {
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

	cf := client.NewDefaultConfig()
	cf.SetAddress(s.c.App.GRPCAddr())
	cf.SetClientCredentials(svr.ClientId, svr.ClientSecret)
	cli, err := client.NewClient(cf)
	if err != nil {
		return err
	}
	client.SetGlobal(cli)
	return err
}
