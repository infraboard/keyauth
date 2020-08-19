package http

import (
	"errors"

	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
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
	if pkg.IP2Region == nil {
		return errors.New("denpence IP2Region service is nil")
	}

	h.service = pkg.IP2Region
	return nil
}

func init() {
	pkg.RegistryHTTPV1("ip2region", api)
}
