package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/user/types"
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
	domainRouter.Handle("POST", "/", h.CreateDomain).SetAllow(types.UserType_SUPPER)
	domainRouter.Handle("GET", "/", h.ListDomains).SetAllow(types.UserType_SUPPER)
	domainRouter.Handle("GET", "/:name", h.GetDomain).SetAllow(types.UserType_DOMAIN_ADMIN)
	domainRouter.Handle("PUT", "/:name", h.PutDomain).SetAllow(types.UserType_DOMAIN_ADMIN)
	domainRouter.Handle("PATCH", "/:name", h.PatchDomain).SetAllow(types.UserType_DOMAIN_ADMIN)
	domainRouter.Handle("DELETE", "/:name", h.DeleteDomain).SetAllow(types.UserType_SUPPER)
	domainRouter.Handle("PUT", "/:name/security", h.UpdateDomainSecurity).SetAllow(types.UserType_DOMAIN_ADMIN)
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
