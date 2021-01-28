package http

import (
	"errors"

	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/permission"
)

var (
	api = &handler{}
)

type handler struct {
	service permission.PermissionServiceServer
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("permission")
	r.BasePath("namespaces")
	r.Handle("GET", "/:id/permissions", h.List).AddLabel(label.List)
	r.Handle("GET", "/:id/permissions/endpoints/:eid", h.Get).AddLabel(label.Get)
}

func (h *handler) Config() error {
	if pkg.Namespace == nil {
		return errors.New("denpence namespace service is nil")
	}

	h.service = pkg.Permission
	return nil
}

func init() {
	pkg.RegistryHTTPV1("permission", api)
}
