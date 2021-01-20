package http

import (
	"errors"

	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/domain"
)

var (
	api = &handler{}
)

type handler struct {
	service domain.DomainServiceServer
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	domainRouter := router.ResourceRouter("domain")
	domainRouter.BasePath("domains")
	domainRouter.Permission(true)
	domainRouter.Handle("POST", "/", h.CreateDomain).AddLabel(label.Create)
	domainRouter.Handle("GET", "/", h.ListDomains).AddLabel(label.List)
	domainRouter.Handle("GET", "/:name", h.GetDomain).AddLabel(label.Get)
	domainRouter.Handle("PUT", "/:name", h.PutDomain).AddLabel(label.Update)
	domainRouter.Handle("PATCH", "/:name", h.PatchDomain).AddLabel(label.Update)
	domainRouter.Handle("DELETE", "/:name", h.DeleteDomain).AddLabel(label.Delete)
	domainRouter.Handle("PUT", "/:name/security", h.UpdateDomainSecurity).AddLabel(label.Update)
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
