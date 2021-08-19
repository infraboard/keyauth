package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/department"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/user/types"
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
	appRouter.Handle("POST", "/", h.CreateJoinApply)
	appRouter.Handle("GET", "/", h.QueryJoinApply)
	appRouter.Handle("GET", "/:id", h.GetJoinApply)
	appRouter.Handle("PATCH", "/:id", h.DealJoinApply)

	appRouter.BasePath("departments")
	appRouter.Handle("POST", "/", h.Create).SetAllow(types.UserType_ORG_ADMIN)
	appRouter.Handle("GET", "/", h.List)
	appRouter.Handle("GET", "/:id", h.Get)
	appRouter.Handle("PUT", "/:id", h.Put)
	appRouter.Handle("PATCH", "/:id", h.Patch)
	appRouter.Handle("GET", "/:id/subs", h.GetSub)
	appRouter.Handle("DELETE", "/:id", h.Delete).SetAllow(types.UserType_ORG_ADMIN)
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
