package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/policy"
)

var (
	api = &handler{}
)

type handler struct {
	service policy.PolicyServiceClient
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("policy")

	r.Allow()
	r.BasePath("policies")
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

	h.service = client.Policy()
	return nil
}

func init() {
	pkg.RegistryHTTPV1("policy", api)
}
