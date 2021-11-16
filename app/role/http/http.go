package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/app/role"
	"github.com/infraboard/keyauth/app/user/types"
)

var (
	api = &handler{}
)

type handler struct {
	service role.ServiceServer
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("role")

	r.BasePath("roles")
	r.Handle("POST", "/", h.CreateRole).SetAllow(types.UserType_PERM_ADMIN)
	r.Handle("GET", "/", h.QueryRole)
	r.Handle("GET", "/:id", h.DescribeRole)
	r.Handle("DELETE", "/:id", h.DeleteRole).SetAllow(types.UserType_PERM_ADMIN)
	r.Handle("GET", "/:id/permissions", h.ListRolePermission)
	r.Handle("POST", "/:id/permissions", h.AddPermissionToRole).SetAllow(types.UserType_PERM_ADMIN)
	r.Handle("DELETE", "/:id/permissions", h.RemovePermissionFromRole).SetAllow(types.UserType_PERM_ADMIN)
	r.BasePath("permissions")
	r.Handle("PUT", "/:id", h.UpdatePermission).SetAllow(types.UserType_PERM_ADMIN)
}

func (h *handler) Config() error {
	h.service = app.GetGrpcApp(role.AppName).(role.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return role.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
