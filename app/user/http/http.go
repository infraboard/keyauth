package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/app/domain"
	"github.com/infraboard/keyauth/app/user"
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
	prmaryRouter := router.ResourceRouter("primary_account")
	prmaryRouter.BasePath("users")
	prmaryRouter.Handle("POST", "/", h.CreatePrimayAccount)
	prmaryRouter.Handle("DELETE", "/", h.DestroyPrimaryAccount)

	ramRouter := router.ResourceRouter("ram_account")
	ramRouter.BasePath("sub_users")
	ramRouter.Handle("POST", "/", h.CreateSubAccount)
	ramRouter.Handle("GET", "/", h.QuerySubAccount)
	ramRouter.Handle("GET", "/:account", h.DescribeSubAccount)
	ramRouter.Handle("PATCH", "/:account", h.PatchSubAccount)
	ramRouter.Handle("DELETE", "/:account", h.DestroySubAccount)
	ramRouter.BasePath("manage")
	ramRouter.Handle("POST", "/block", h.BlockSubAccount)

	portalRouter := router.ResourceRouter("profile")
	portalRouter.BasePath("profile")
	portalRouter.Handle("GET", "/", h.QueryProfile)
	portalRouter.Handle("GET", "/domain", h.QueryDomain)
	portalRouter.Handle("PUT", "/", h.PutProfile)
	portalRouter.Handle("PATCH", "/", h.PatchProfile)

	domRouter := router.ResourceRouter("domain")
	domRouter.BasePath("settings/domain")
	domRouter.Handle("PUT", "/info", h.UpdateDomainInfo)
	domRouter.Handle("PUT", "/security", h.UpdateDomainSecurity)

	passRouter := router.ResourceRouter("password")
	passRouter.BasePath("password")
	passRouter.Handle("POST", "/", h.GeneratePassword)
	passRouter.Handle("PUT", "/", h.UpdatePassword)
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
