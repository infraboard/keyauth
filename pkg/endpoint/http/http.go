package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/user/types"
)

var (
	api = &handler{}
)

type handler struct {
	endpoint endpoint.EndpointServiceClient
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("endpoint")

	r.BasePath("endpoints")
	r.Handle("POST", "/", h.Create).SetAllow(types.UserType_INTERNAL)
	r.Handle("GET", "/", h.List)
	r.Handle("GET", "/:id", h.Get)

	rr := router.ResourceRouter("resource")
	rr.BasePath("resources")
	rr.Handle("GET", "/", h.ListResource)
}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.endpoint = client.Endpoint()
	return nil
}

func init() {
	pkg.RegistryHTTPV1("endpoint", api)
}
