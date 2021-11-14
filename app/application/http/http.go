package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/app/application"
)

var (
	api = &handler{}
)

type handler struct {
	service application.ServiceServer
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	self := router.ResourceRouter("application")
	self.BasePath("applications")
	self.Handle("POST", "/", h.CreateApplication)
	self.Handle("GET", "/", h.QueryApplication)
	self.Handle("GET", "/:id", h.GetApplication)
	self.Handle("DELETE", "/:id", h.DestroyApplication)
}

func (h *handler) Config() error {
	h.service = app.GetGrpcApp(application.AppName).(application.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return application.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
