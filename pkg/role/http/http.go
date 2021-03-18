package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/role"
)

var (
	api = &handler{}
)

type handler struct {
	service role.RoleServiceClient
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
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.Role()
	return nil
}

func init() {
	pkg.RegistryHTTPV1("role", api)
}
