package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/app/session"
)

var (
	api = &handler{}
)

type handler struct {
	service session.ServiceServer
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("sessions")
	r.BasePath("sessions")
	r.Handle("GET", "/", h.List)
	r.Handle("GET", "/:id", h.Get)
}

func (h *handler) Config() error {
	h.service = app.GetGrpcApp(session.AppName).(session.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return session.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
