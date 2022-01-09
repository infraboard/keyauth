package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/apps/token"
)

var (
	api = &handler{}
)

type handler struct {
	service token.ServiceServer
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("token")

	r.BasePath("/oauth2/tokens")
	r.Handle("POST", "/", h.IssueToken).DisableAuth()
	r.Handle("GET", "/", h.ValidateToken)
	r.Handle("DELETE", "/", h.RevolkToken)

	r.BasePath("/self/tokens")
	r.Handle("GET", "/", h.QueryToken)
	r.Handle("POST", "/", h.ChangeNamespace)
	r.Handle("DELETE", "/", h.DeleteToken)
}

func (h *handler) Config() error {
	h.service = app.GetGrpcApp(token.AppName).(token.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return token.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
