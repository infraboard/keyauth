package http

import (
	"errors"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg/session"
)

var (
	api = &handler{}
)

type handler struct {
	service session.ServiceClient
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("sessions")
	r.BasePath("sessions")
	r.Handle("GET", "/", h.List)
	r.Handle("GET", "/:id", h.Get)
}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.Session()
	return nil
}

func (h *handler) Name() string {
	return session.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
