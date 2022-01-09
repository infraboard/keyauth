package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/apps/department"
	"github.com/infraboard/keyauth/apps/user"
	"github.com/infraboard/keyauth/apps/user/types"
)

var (
	api = &handler{}
)

type handler struct {
	service department.ServiceServer
	user    user.ServiceServer
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("department")
	r.BasePath("join_apply")
	r.Handle("POST", "/", h.CreateJoinApply)
	r.Handle("GET", "/", h.QueryJoinApply)
	r.Handle("GET", "/:id", h.GetJoinApply)
	r.Handle("PATCH", "/:id", h.DealJoinApply)

	r.BasePath("departments")
	r.Handle("POST", "/", h.Create).SetAllow(types.UserType_ORG_ADMIN)
	r.Handle("GET", "/", h.List)
	r.Handle("GET", "/:id", h.Get)
	r.Handle("PUT", "/:id", h.Put)
	r.Handle("PATCH", "/:id", h.Patch)
	r.Handle("GET", "/:id/subs", h.GetSub)
	r.Handle("DELETE", "/:id", h.Delete).SetAllow(types.UserType_ORG_ADMIN)
}

func (h *handler) Config() error {
	h.service = app.GetGrpcApp(department.AppName).(department.ServiceServer)
	h.user = app.GetGrpcApp(user.AppName).(user.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return department.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
