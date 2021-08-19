package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/namespace"
	"github.com/infraboard/keyauth/pkg/user/types"
)

var (
	api = &handler{}
)

type handler struct {
	service namespace.NamespaceServiceClient
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	// 需要做权限限制
	r := router.ResourceRouter("namespace")
	r.BasePath("namespaces")
	r.Handle("POST", "/", h.Create).SetAllow(types.UserType_PERM_ADMIN)
	r.Handle("GET", "/", h.List).SetAllow(types.UserType_PERM_ADMIN)
	r.Handle("DELETE", "/:id", h.Delete).SetAllow(types.UserType_PERM_ADMIN)
	r.Handle("GET", "/:id", h.Get)

	// 获取自己的namespace 不需做权限限制
	self := router.ResourceRouter("namespace")
	self.BasePath("self/namespaces")
	self.Handle("GET", "/", h.ListSelfNamespace)
}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.Namespace()
	return nil
}

func init() {
	pkg.RegistryHTTPV1("namespace", api)
}
