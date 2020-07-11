package http

import (
	"errors"

	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/endpoint"
)

var (
	api = &handler{}
)

type handler struct {
	endpoint endpoint.Service
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("endpoint")
	r.BasePath("endpoints")
	r.Handle("POST", "/", h.Create).AddLabel(label.Create)
	r.Handle("GET", "/", h.List).AddLabel(label.List)
}

func (h *handler) Config() error {
	if pkg.Endpoint == nil {
		return errors.New("denpence endpoint service is nil")
	}

	h.endpoint = pkg.Endpoint
	return nil
}

func init() {
	pkg.RegistryHTTPV1("endpoint", api)
}
