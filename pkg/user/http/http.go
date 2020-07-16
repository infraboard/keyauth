package http

import (
	"errors"

	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/user"
)

var (
	api = &handler{}
)

type handler struct {
	service user.Service
	domain  domain.Service
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	prmaryRouter := router.ResourceRouter("primary_account")
	prmaryRouter.BasePath("users")
	prmaryRouter.Permission(true)
	prmaryRouter.Handle("POST", "/", h.CreatePrimayAccount).AddLabel(label.Create)
	prmaryRouter.Handle("DELETE", "/", h.DestroyPrimaryAccount).AddLabel(label.Delete)

	ramRouter := router.ResourceRouter("ram_account")
	ramRouter.Permission(true)
	ramRouter.BasePath("sub_users")
	ramRouter.Handle("POST", "/", h.CreateSubAccount).AddLabel(label.Create)
	ramRouter.Handle("GET", "/", h.QuerySubAccount).AddLabel(label.List)

	portalRouter := router.ResourceRouter("profile")
	portalRouter.BasePath("profile")
	portalRouter.Handle("GET", "/", h.QueryProfile).AddLabel(label.Get)
	portalRouter.Handle("GET", "/domain", h.QueryDomain).AddLabel(label.Get)
	portalRouter.Handle("PUT", "/", h.UpdateProfile).AddLabel(label.Create)
	portalRouter.Permission(true)
	portalRouter.Handle("PUT", "/domain", h.UpdateDomain).AddLabel(label.Create)

	passRouter := router.ResourceRouter("password")
	passRouter.BasePath("password")
	passRouter.Handle("PUT", "/", h.UpdatePassword).AddLabel(label.Update)
}

func (h *handler) Config() error {
	if pkg.User == nil {
		return errors.New("denpence user service is nil")
	}

	if pkg.Domain == nil {
		return errors.New("denpence domain service is nil")
	}

	h.service = pkg.User
	h.domain = pkg.Domain
	return nil
}

func init() {
	pkg.RegistryHTTPV1("user", api)
}
