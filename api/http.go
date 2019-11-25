package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/mcube/http/middleware/accesslog"
	"github.com/infraboard/mcube/http/middleware/recovery"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/http/router/httprouter"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// NewHTTPService 构建函数
func NewHTTPService() *HTTPService {
	r := httprouter.New()
	r.Use(recovery.NewWithLogger(zap.L().Named("Recovery")))
	r.Use(accesslog.NewWithLogger(zap.L().Named("AccessLog")))

	server := &http.Server{
		ReadHeaderTimeout: 20 * time.Second,
		ReadTimeout:       20 * time.Second,
		WriteTimeout:      25 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    1 << 20,
		Addr:              "0.0.0.0:8848",
		Handler:           r,
	}

	return &HTTPService{
		r:      r,
		server: server,
		l:      zap.L(),
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
func (s *HTTPService) Start() {
	// 启动HTTP服务
	s.l.Infof("服务启动成功, 监听地址: %s", s.http.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.l.Info("service is stopped")
		}
		return fmt.Errorf("start service error, %s", err.Error())
	}
	return nil
}

// Endpoints 服务所有的路由条目, 供服务注册时使用
func (s *HTTPService) Endpoints() []router.Entry {
	return s.r.GetEndpoints()
}
