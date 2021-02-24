package http

import (
	"errors"

	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/token"
)

var (
	api = &handler{}
)

type handler struct {
	service micro.MicroServiceServer
	token   token.TokenServiceServer
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
	if pkg.ApplicationAdmin == nil {
		return errors.New("denpence application service is nil")
	}
	if pkg.Token == nil {
		return errors.New("denpence token service is nil")
	}

	h.service = pkg.Micro
	h.token = pkg.Token
	return nil
}

func init() {
	pkg.RegistryHTTPV1("service", api)
}
