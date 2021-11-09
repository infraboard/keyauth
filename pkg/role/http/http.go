package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg/role"
)

var (
	api = &handler{}
)

type handler struct {
	service role.RoleServiceServer
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("role")

	r.BasePath("roles")
	r.Handle("POST", "/", h.CreateRole)
	r.Handle("GET", "/", h.QueryRole)
	r.Handle("GET", "/:id", h.DescribeRole)
	r.Handle("DELETE", "/:id", h.DeleteRole)
	r.Handle("GET", "/:id/permissions", h.ListRolePermission)
	r.Handle("POST", "/:id/permissions", h.AddPermissionToRole)
	r.Handle("DELETE", "/:id/permissions", h.RemovePermissionFromRole)
	r.BasePath("permissions")
	r.Handle("PUT", "/:id", h.UpdatePermission)
}

func (h *handler) Config() error {
	h.service = app.GetGrpcApp(role.AppName).(role.RoleServiceServer)
	return nil
}

func (h *handler) Name() string {
	return role.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
