package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/token"
)

var (
	api = &handler{}
)

type handler struct {
	service token.Service
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("token")
	r.BasePath("/oauth2/tokens/")
	r.AddPublict("POST", "/", h.IssueToken)

}

func (h *handler) Config() error {
	if pkg.Token == nil {
		return errors.New("denpence token service is nil")
	}

	h.service = pkg.Token
	return nil
}

func init() {
	pkg.RegistryHTTPV1("token", api)
}
