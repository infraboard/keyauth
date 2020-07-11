package http

import (
	"errors"

	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/application"
)

var (
	api = &handler{}
)

type handler struct {
	service application.Service
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	appRouter := router.ResourceRouter("application")
	appRouter.BasePath("applications")
	appRouter.Permission(true)
	appRouter.Handle("POST", "/", h.CreateUserApplication).AddLabel(label.Create)
	appRouter.Handle("GET", "/", h.QueryUserApplication).AddLabel(label.List)
	appRouter.Handle("GET", "/:id", h.GetApplication).AddLabel(label.Get)
	appRouter.Handle("DELETE", "/:id", h.DestroyApplication).AddLabel(label.Delete)

}

func (h *handler) Config() error {
	if pkg.Application == nil {
		return errors.New("denpence application service is nil")
	}

	h.service = pkg.Application
	return nil
}

func init() {
	pkg.RegistryHTTPV1("application", api)
}
