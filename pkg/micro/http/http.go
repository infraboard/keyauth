package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
)

var (
	api = &handler{}
)

type handler struct {
	service micro.MicroServiceServer
	token   token.TokenServiceServer
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("service")
	r.Allow(types.UserType_INTERNAL)
	r.BasePath("services")

	r.Handle("GET", "/", h.QueryService)
	r.Handle("POST", "/", h.CreateService)
	r.Handle("GET", "/:id", h.GetService)
	r.Handle("DELETE", "/:id", h.DestroyService)
	r.Handle("POST", "/:id/refresh_client_secret", h.RefreshServiceClientSecret)
}

func (h *handler) Config() error {
	h.service = app.GetGrpcApp(micro.AppName).(micro.MicroServiceServer)
	h.token = app.GetGrpcApp(token.AppName).(token.TokenServiceServer)
	return nil
}

func (h *handler) Name() string {
	return micro.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
