package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/apps/application"
)

var (
	api = &handler{}
)

type handler struct {
	service application.ServiceServer
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("application")
	r.BasePath("applications")
	r.Handle("POST", "/", h.CreateApplication)
	r.Handle("GET", "/", h.QueryApplication)
	r.Handle("GET", "/:id", h.GetApplication)
	r.Handle("DELETE", "/:id", h.DestroyApplication)
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
