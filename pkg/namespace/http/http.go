package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg/namespace"
	"github.com/infraboard/keyauth/pkg/user/types"
)

var (
	api = &handler{}
)

type handler struct {
	service namespace.NamespaceServiceServer
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
	h.service = app.GetGrpcApp(namespace.AppName).(namespace.NamespaceServiceServer)
	return nil
}

func (h *handler) Name() string {
	return namespace.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
