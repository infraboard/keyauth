package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/service"
)

var (
	api = &handler{}
)

type handler struct {
	service service.Service
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("service")
	r.BasePath("services")
	r.AddProtected("GET", "/", h.QueryService)
	r.AddProtected("POST", "/", h.CreateService)
	r.AddProtected("GET", "/:name", h.GetService)
	r.AddProtected("DELETE", "/:name", h.DestroyService)
}

func (h *handler) Config() error {
	if pkg.Application == nil {
		return errors.New("denpence application service is nil")
	}

	h.service = pkg.MicroService
	return nil
}

func init() {
	pkg.RegistryHTTPV1("service", api)
}
