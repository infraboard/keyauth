package protocol

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/middleware/accesslog"
	"github.com/infraboard/mcube/http/middleware/cors"
	"github.com/infraboard/mcube/http/middleware/recovery"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/http/router/httprouter"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	auther "github.com/infraboard/keyauth/common/interceptor/http"
	"github.com/infraboard/keyauth/conf"
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
		ReadTimeout:       20 * time.Second,
		WriteTimeout:      25 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    1 << 20,
		Addr:              conf.C().App.HTTPAddr(),
		Handler:           r,
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

func (s *HTTPService) PathPrefix() string {
	return fmt.Sprintf("/%s/api/v1", s.c.App.Name)
}

// Start 启动服务
func (s *HTTPService) Start() error {
	// 装置子服务路由
	if err := app.LoadHttpApp(s.PathPrefix(), s.r); err != nil {
		return err
	}

	s.l.Infof("loaded http service: %v", app.LoadedHttpApp())

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
