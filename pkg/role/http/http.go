package http

import (
	"errors"

	"github.com/infraboard/mcube/http/label"
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
	r.Handle("POST", "/", h.CreateRole).AddLabel(label.Create)
	r.Handle("GET", "/", h.QueryRole).AddLabel(label.List)
	r.Handle("GET", "/:name", h.DescribeRole).AddLabel(label.Get)
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
