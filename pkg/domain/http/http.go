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
	domainRouter.Permission(true)
	domainRouter.Handle("POST", "/", h.CreateDomain)
	domainRouter.Handle("GET", "/", h.ListDomains)
	domainRouter.Handle("GET", "/:id", h.GetDomain)
	domainRouter.Handle("PUT", "/:id", h.UpdateDomain)
	domainRouter.Handle("DELETE", "/:id", h.DeleteDomain)
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
