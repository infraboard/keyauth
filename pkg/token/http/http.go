package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/token"
)

var (
	api = &handler{}
)

type handler struct {
	service token.TokenServiceClient
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("token")

	r.BasePath("/oauth2/tokens")
	r.Handle("POST", "/", h.IssueToken).DisableAuth()
	r.Handle("GET", "/", h.ValidateToken)
	r.Handle("DELETE", "/", h.RevolkToken)

	r.BasePath("/self/tokens")
	r.Handle("GET", "/", h.QueryToken)
	r.Handle("POST", "/", h.ChangeNamespace)
	r.Handle("DELETE", "/", h.DeleteToken)
}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.Token()
	return nil
}

func init() {
	pkg.RegistryHTTPV1("token", api)
}
