package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/app/policy"
)

var (
	api = &handler{}
)

type handler struct {
	service policy.PolicyServiceServer
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("policy")

	r.Allow()
	r.BasePath("policies")
	r.Handle("POST", "/", h.Create)
	r.Handle("GET", "/", h.List)
	r.Handle("GET", "/:id", h.Get)
	r.Handle("DELETE", "/:id", h.Delete)
}

func (h *handler) Config() error {
	h.service = app.GetGrpcApp(policy.AppName).(policy.PolicyServiceServer)
	return nil
}

func (h *handler) Name() string {
	return policy.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
