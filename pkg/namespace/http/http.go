package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/namespace"
)

var (
	api = &handler{}
)

type handler struct {
	service namespace.NamespaceServiceClient
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("namespace")
	// 获取自己的namespace 不需做权限限制
	r.BasePath("self/namespaces")
	r.Handle("GET", "/", h.ListSelfNamespace)

	// 需要做权限限制
	r.BasePath("namespaces")
	r.Handle("POST", "/", h.Create)
	r.Handle("GET", "/", h.List)
	r.Handle("GET", "/:id", h.Get)
	r.Handle("DELETE", "/:id", h.Delete)
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
