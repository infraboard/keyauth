package http

import (
	"github.com/infraboard/keyauth/apps/otp"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"
)

var (
	api = &handler{}
)

type handler struct {
	service otp.ServiceServer
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {

	r := router.ResourceRouter("otp")
	r.BasePath("otp")
	r.Handle("DELETE", "/:account", h.DeleteOTP)
	r.Handle("GET", "/:account", h.GetOTP)
	r.Handle("POST", "/", h.UpdateOTP)
}

func (h *handler) Config() error {
	h.service = app.GetGrpcApp(otp.AppName).(otp.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return otp.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
