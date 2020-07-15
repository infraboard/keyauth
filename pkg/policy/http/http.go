package http

import (
	"errors"

	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/policy"
)

var (
	api = &handler{}
)

type handler struct {
	service policy.Service
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("policy")
	r.BasePath("policies")
	r.Permission(true)
	r.Handle("POST", "/", h.Create).AddLabel(label.Create)
	r.Handle("GET", "/", h.List).AddLabel(label.List)
	r.Handle("GET", "/:id", h.Get).AddLabel(label.Get)
}

func (h *handler) Config() error {
	if pkg.Policy == nil {
		return errors.New("denpence policy service is nil")
	}

	h.service = pkg.Policy
	return nil
}

func init() {
	pkg.RegistryHTTPV1("policy", api)
}
