package http

import (
	"errors"

	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/department"
	"github.com/infraboard/keyauth/pkg/user"
)

var (
	api = &handler{}
)

type handler struct {
	service department.DepartmentServiceClient
	user    user.UserServiceClient
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	appRouter := router.ResourceRouter("department")
	appRouter.BasePath("join_apply")
	appRouter.Handle("POST", "/", h.CreateJoinApply).AddLabel(label.Create)
	appRouter.Handle("GET", "/", h.QueryJoinApply).AddLabel(label.List)
	appRouter.Handle("GET", "/:id", h.GetJoinApply).AddLabel(label.Get)
	appRouter.Handle("PATCH", "/:id", h.DealJoinApply).AddLabel(label.Update)

	appRouter.BasePath("departments")
	appRouter.Permission(true)
	appRouter.Handle("POST", "/", h.Create).AddLabel(label.Create)
	appRouter.Handle("GET", "/", h.List).AddLabel(label.List)
	appRouter.Handle("GET", "/:id", h.Get).AddLabel(label.Get)
	appRouter.Handle("PUT", "/:id", h.Put).AddLabel(label.Update)
	appRouter.Handle("PATCH", "/:id", h.Patch).AddLabel(label.Update)
	appRouter.Handle("GET", "/:id/subs", h.GetSub).AddLabel(label.Get)
	appRouter.Handle("DELETE", "/:id", h.Delete).AddLabel(label.Delete)
}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.Department()
	h.user = client.User()
	return nil
}

func init() {
	pkg.RegistryHTTPV1("department", api)
}
