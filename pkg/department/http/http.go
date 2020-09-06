package http

import (
	"errors"

	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/department"
	"github.com/infraboard/keyauth/pkg/user"
)

var (
	api = &handler{}
)

type handler struct {
	service department.Service
	user    user.Service
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	appRouter := router.ResourceRouter("department")
	appRouter.BasePath("departments")
	appRouter.Permission(true)
	appRouter.Handle("POST", "/", h.Create).AddLabel(label.Create)
	appRouter.Handle("GET", "/", h.List).AddLabel(label.List)
	appRouter.Handle("GET", "/:id", h.Get).AddLabel(label.Get)
	appRouter.Handle("PUT", "/:id", h.Put).AddLabel(label.Update)
	appRouter.Handle("PATCH", "/:id", h.Patch).AddLabel(label.Update)
	appRouter.Handle("GET", "/:id/subs", h.GetSub).AddLabel(label.Get)
	appRouter.Handle("DELETE", "/:id", h.Delete).AddLabel(label.Delete)

}

func (h *handler) Config() error {
	if pkg.Department == nil {
		return errors.New("denpence department service is nil")
	}

	if pkg.User == nil {
		return errors.New("dependence user service is nil")
	}

	h.service = pkg.Department
	h.user = pkg.User
	return nil
}

func init() {
	pkg.RegistryHTTPV1("department", api)
}
