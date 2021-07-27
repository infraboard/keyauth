package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/permission"
)

var (
	api = &handler{}
)

type handler struct {
	service permission.PermissionServiceClient
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
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.Permission()
	return nil
}

func init() {
	pkg.RegistryHTTPV1("permission", api)
}
