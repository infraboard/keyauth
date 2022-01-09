package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/apps/domain"
	"github.com/infraboard/keyauth/apps/user"
	"github.com/infraboard/keyauth/apps/user/types"
)

var (
	api = &handler{}
)

type handler struct {
	service user.ServiceServer
	domain  domain.ServiceServer
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	prmary := router.ResourceRouter("primary_account")
	prmary.Allow(types.UserType_SUPPER)
	prmary.BasePath("users")
	prmary.Handle("POST", "/", h.CreatePrimayAccount)
	prmary.Handle("DELETE", "/", h.DestroyPrimaryAccount)

	ram := router.ResourceRouter("ram_account")
	ram.Allow(types.UserType_PRIMARY, types.UserType_DOMAIN_ADMIN, types.UserType_ORG_ADMIN)
	ram.BasePath("sub_users")
	ram.Handle("POST", "/", h.CreateSubAccount)
	ram.Handle("GET", "/", h.QuerySubAccount)
	ram.Handle("GET", "/:account", h.DescribeSubAccount)
	ram.Handle("PATCH", "/:account", h.PatchSubAccount)
	ram.Handle("DELETE", "/:account", h.DestroySubAccount)
	ram.BasePath("manage")
	ram.Handle("POST", "/block", h.BlockSubAccount)

	portal := router.ResourceRouter("profile")
	portal.BasePath("profile")
	portal.Handle("GET", "/", h.QueryProfile)
	portal.Handle("GET", "/domain", h.QueryDomain)
	portal.Handle("PUT", "/", h.PutProfile)
	portal.Handle("PATCH", "/", h.PatchProfile)

	dom := router.ResourceRouter("domain")
	dom.Allow(types.UserType_DOMAIN_ADMIN)
	dom.BasePath("settings/domain")
	dom.Handle("PUT", "/info", h.UpdateDomainInfo)
	dom.Handle("PUT", "/security", h.UpdateDomainSecurity)

	pass := router.ResourceRouter("password")
	pass.BasePath("password")
	pass.Handle("POST", "/", h.GeneratePassword)
	pass.Handle("PUT", "/", h.UpdatePassword)
}

func (h *handler) Config() error {
	h.service = app.GetGrpcApp(user.AppName).(user.ServiceServer)
	h.domain = app.GetGrpcApp(domain.AppName).(domain.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return user.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
