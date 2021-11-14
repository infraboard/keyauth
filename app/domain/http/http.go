package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/app/domain"
	"github.com/infraboard/keyauth/app/user/types"
)

var (
	api = &handler{}
)

type handler struct {
	service domain.ServiceServer
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
	h.service = app.GetGrpcApp(domain.AppName).(domain.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return domain.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
