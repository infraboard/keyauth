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
	service application.UserServiceClient
	admin   application.AdminServiceClient
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	appRouter := router.ResourceRouter("application")
	appRouter.BasePath("applications")
	appRouter.Handle("POST", "/", h.CreateUserApplication)
	appRouter.Handle("GET", "/", h.QueryUserApplication)
	appRouter.Handle("GET", "/:id", h.GetApplication)
	appRouter.Handle("DELETE", "/:id", h.DestroyApplication)
}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.ApplicationUser()
	h.admin = client.ApplicationAdmin()
	return nil
}

func init() {
	pkg.RegistryHTTPV1("application", api)
}
