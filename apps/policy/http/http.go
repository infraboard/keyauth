package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/apps/policy"
	"github.com/infraboard/keyauth/apps/user/types"
)

var (
	api = &handler{}
)

type handler struct {
	service policy.ServiceServer
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("policy")
	r.BasePath("policies")
	r.Handle("POST", "/", h.Create).SetAllow(types.UserType_PERM_ADMIN)
	r.Handle("GET", "/", h.List).SetAllow(types.UserType_PERM_ADMIN)
	r.Handle("GET", "/:id", h.Get)
	r.Handle("DELETE", "/:id", h.Delete).SetAllow(types.UserType_PERM_ADMIN)
}

func (h *handler) Config() error {
	h.service = app.GetGrpcApp(policy.AppName).(policy.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return policy.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
