package http

import (
	"errors"

	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/system"
)

var (
	api = &handler{}
)

type handler struct {
	service system.Service
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("system_config")
	r.BasePath("system_config")
	r.Handle("POST", "/email/test", h.TestEmailSend).AddLabel(label.Create)
	r.Handle("POST", "/sms/test", h.TestSMSSend).AddLabel(label.Create)
}

func (h *handler) Config() error {
	if pkg.System == nil {
		return errors.New("denpence system service is nil")
	}

	h.service = pkg.System
	return nil
}

func init() {
	pkg.RegistryHTTPV1("system_config", api)
}
