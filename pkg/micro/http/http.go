package http

import (
	"errors"

	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/token"
)

var (
	api = &handler{}
)

type handler struct {
	service micro.MicroServiceClient
	token   token.TokenServiceClient
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("service")
	r.BasePath("services")
	r.Permission(true)
	r.Handle("GET", "/", h.QueryService).AddLabel(label.List)
	r.Handle("POST", "/", h.CreateService).AddLabel(label.Create)
	r.Handle("GET", "/:id", h.GetService).AddLabel(label.Get)
	r.Handle("DELETE", "/:id", h.DestroyService).AddLabel(label.Delete)
	r.Handle("POST", "/:id/refresh_client_secret", h.RefreshServiceClientSecret).AddLabel(label.Update)
}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.Micro()
	h.token = client.Token()
	return nil
}

func init() {
	pkg.RegistryHTTPV1("service", api)
}
