package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/user"
)

var (
	api = &handler{}
)

type handler struct {
	service user.UserServiceClient
	domain  domain.DomainServiceClient
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
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.User()
	h.domain = client.Domain()
	return nil
}

func init() {
	pkg.RegistryHTTPV1("user", api)
}
