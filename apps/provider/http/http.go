package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/apps/provider"
)

var (
	api = &handler{}
)

type handler struct {
	service provider.LDAP
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("ldap")
	r.BasePath("settings/ldap")
	r.Permission(true)
	r.Handle("POST", "/", h.Create).AddLabel(label.Create)
	r.Handle("POST", "/conn_check", h.Check).AddLabel(label.Create)
	r.Handle("GET", "/", h.Get).AddLabel(label.List)

}

func (h *handler) Config() error {
	h.service = app.GetInternalApp(provider.AppName).(provider.LDAP)
	return nil
}

func (h *handler) Name() string {
	return provider.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
