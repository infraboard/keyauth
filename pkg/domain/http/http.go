package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/domain"
)

var (
	api = &handler{}
)

type handler struct {
	service domain.Service
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	domainRouter := router.ResourceRouter("domain")
	domainRouter.BasePath("domains")
	domainRouter.AddProtected("POST", "/", h.CreateDomain)
	domainRouter.AddPublict("GET", "/", h.ListDomains)
	domainRouter.AddProtected("GET", "/:id", h.GetDomain)
	domainRouter.AddProtected("PUT", "/:id", h.UpdateDomain)
	domainRouter.AddProtected("DELETE", "/:id", h.DeleteDomain)
}

func (h *handler) Config() error {
	if pkg.Domain == nil {
		return errors.New("denpence domain service is nil")
	}

	h.service = pkg.Domain
	return nil
}

func init() {
	pkg.RegistryHTTPV1("domain", api)
}
