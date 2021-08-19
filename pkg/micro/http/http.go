package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
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
	r.Allow(types.UserType_INTERNAL)
	r.BasePath("services")

	r.Handle("GET", "/", h.QueryService)
	r.Handle("POST", "/", h.CreateService)
	r.Handle("GET", "/:id", h.GetService)
	r.Handle("DELETE", "/:id", h.DestroyService)
	r.Handle("POST", "/:id/refresh_client_secret", h.RefreshServiceClientSecret)
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
