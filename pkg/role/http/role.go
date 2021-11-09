package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/keyauth/pkg/token"
)

// CreateApplication 创建自定义角色
func (h *handler) CreateRole(w http.ResponseWriter, r *http.Request) {
	req := role.NewCreateRoleRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	d, err := h.service.CreateRole(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}

func (h *handler) QueryRole(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := role.NewQueryRoleRequestFromHTTP(r)
	req.Domain = tk.Domain

	apps, err := h.service.QueryRole(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, apps)
}

func (h *handler) DescribeRole(w http.ResponseWriter, r *http.Request) {

	rctx := context.GetContext(r)
	pid := rctx.PS.ByName("id")
	qs := r.URL.Query()

	req := role.NewDescribeRoleRequestWithID(pid)
	req.WithPermissions = qs.Get("with_permissions") == "true"

	ins, err := h.service.DescribeRole(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)
	req := role.NewDeleteRoleWithID(rctx.PS.ByName("id"))

	_, err := h.service.DeleteRole(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
}

// ListRolePermission 创建自定义角色
func (h *handler) ListRolePermission(w http.ResponseWriter, r *http.Request) {
	req := role.NewQueryPermissionRequestFromHTTP(r)
	rctx := context.GetContext(r)
	req.RoleId = rctx.PS.ByName("id")

	d, err := h.service.QueryPermission(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

// CreateApplication 创建自定义角色
func (h *handler) AddPermissionToRole(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)
	req := role.NewAddPermissionToRoleRequest()
	req.RoleId = rctx.PS.ByName("id")

	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	d, err := h.service.AddPermissionToRole(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}

// CreateApplication 创建自定义角色
func (h *handler) RemovePermissionFromRole(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)
	req := role.NewRemovePermissionFromRoleRequest()
	req.RoleId = rctx.PS.ByName("id")

	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	d, err := h.service.RemovePermissionFromRole(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}

func (h *handler) UpdatePermission(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	// 查找出原来的domain
	req := role.NewUpdatePermissionRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.Id = rctx.PS.ByName("id")

	ins, err := h.service.UpdatePermission(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
