package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg/ip2region"
)

var (
	api = &handler{}
)

type handler struct {
	service ip2region.Service
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	geoipRouter := router.ResourceRouter("IP")
	geoipRouter.BasePath("ip2region")
	geoipRouter.Handle("GET", "/query", h.LoopupIP).AddLabel(label.Get)

	geoipRouter.Permission(true)
	geoipRouter.Handle("POST", "/dbfile", h.UpdateDBFile).AddLabel(label.Create)

}

func (h *handler) Config() error {
	h.service = app.GetInternalApp(ip2region.AppName).(ip2region.Service)
	return nil
}

func (h *handler) Name() string {
	return ip2region.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
