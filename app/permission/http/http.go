package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/app/permission"
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
	r.Handle("GET", "/:id/permissions", h.ListPermission)
	r.Handle("GET", "/:id/roles", h.ListRole)
	r.Handle("POST", "/:id/permissions", h.CheckPermission)
}

func (h *handler) Config() error {
	h.service = app.GetGrpcApp(permission.AppName).(permission.PermissionServiceServer)
	return nil
}

func (h *handler) Name() string {
	return permission.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
