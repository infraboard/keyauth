package http

import (
	"errors"

	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/micro"
)

var (
	api = &handler{}
)

type handler struct {
	service micro.Service
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("service")
	r.BasePath("services")
	r.Permission(true)
	r.Handle("GET", "/", h.QueryService).AddLabel(label.List)
	r.Handle("POST", "/", h.CreateService).AddLabel(label.Create)
	r.Handle("GET", "/:name", h.GetService).AddLabel(label.Get)
	r.Handle("DELETE", "/:name", h.DestroyService).AddLabel(label.Delete)
}

func (h *handler) Config() error {
	if pkg.Application == nil {
		return errors.New("denpence application service is nil")
	}

	h.service = pkg.Micro
	return nil
}

func init() {
	pkg.RegistryHTTPV1("service", api)
}
