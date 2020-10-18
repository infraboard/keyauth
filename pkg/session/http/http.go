package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/session"
)

var (
	api = &handler{}
)

type handler struct {
	service session.Service
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("sessions")
	r.BasePath("sessions")
	r.Permission(true)
	r.Handle("GET", "/", h.QueryLoginLog)
}

func (h *handler) Config() error {
	if pkg.Domain == nil {
		return errors.New("denpence domain service is nil")
	}

	h.service = pkg.Session
	return nil
}

func init() {
	pkg.RegistryHTTPV1("session", api)
}
