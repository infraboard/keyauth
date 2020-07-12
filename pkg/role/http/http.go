package http

import (
	"errors"

	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/role"
)

var (
	api = &handler{}
)

type handler struct {
	service role.Service
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("role")
	r.BasePath("roles")
	r.Handle("POST", "/", h.CreateRole).AddLabel(label.Create)
	r.Handle("GET", "/", h.QueryRole).AddLabel(label.List)
}

func (h *handler) Config() error {
	if pkg.Role == nil {
		return errors.New("denpence application service is nil")
	}

	h.service = pkg.Role
	return nil
}

func init() {
	pkg.RegistryHTTPV1("role", api)
}
