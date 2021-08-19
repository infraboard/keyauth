package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/application"
)

var (
	api = &handler{}
)

type handler struct {
	service application.ApplicationServiceClient
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	self := router.ResourceRouter("application")
	self.BasePath("applications")
	self.Handle("POST", "/", h.CreateApplication)
	self.Handle("GET", "/", h.QueryApplication)
	self.Handle("GET", "/:id", h.GetApplication)
	self.Handle("DELETE", "/:id", h.DestroyApplication)
}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.Application()
	return nil
}

func init() {
	pkg.RegistryHTTPV1("application", api)
}
