package http

import (
	"errors"

	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/provider"
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
	r.BasePath("ldaps")
	r.Permission(true)
	r.Handle("POST", "/", h.Create).AddLabel(label.Create)
	r.Handle("GET", "/", h.List).AddLabel(label.List)
	r.Handle("GET", "/:id", h.Get).AddLabel(label.Get)
}

func (h *handler) Config() error {
	if pkg.LDAP == nil {
		return errors.New("denpence namespace service is nil")
	}

	h.service = pkg.LDAP
	return nil
}

func init() {
	pkg.RegistryHTTPV1("ldap", api)
}
