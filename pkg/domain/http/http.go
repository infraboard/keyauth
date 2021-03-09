package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/domain"
)

var (
	api = &handler{}
)

type handler struct {
	service domain.DomainServiceClient
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	domainRouter := router.ResourceRouter("domain")

	domainRouter.BasePath("domains")
	domainRouter.Handle("POST", "/", h.CreateDomain)
	domainRouter.Handle("GET", "/", h.ListDomains)
	domainRouter.Handle("GET", "/:name", h.GetDomain)
	domainRouter.Handle("PUT", "/:name", h.PutDomain)
	domainRouter.Handle("PATCH", "/:name", h.PatchDomain)
	domainRouter.Handle("DELETE", "/:name", h.DeleteDomain)
	domainRouter.Handle("PUT", "/:name/security", h.UpdateDomainSecurity)
}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.Domain()
	return nil
}

func init() {
	pkg.RegistryHTTPV1("domain", api)
}
