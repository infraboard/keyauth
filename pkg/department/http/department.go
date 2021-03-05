package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/department"
)

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.GetGrpcCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := department.NewQueryDepartmentRequestFromHTTP(r)
	apps, err := h.service.QueryDepartment(ctx.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, apps)
	return
}

// Create 创建主账号
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.GetGrpcCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := department.NewCreateDepartmentRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.CreateDepartment(ctx.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}

// Create 创建主账号
func (h *handler) Put(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)
	ctx, err := pkg.GetGrpcCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := department.NewPutUpdateDepartmentRequest(rctx.PS.ByName("id"))
	if err := request.GetDataFromRequest(r, req.Data); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateDepartment(ctx.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}

// Create 创建主账号
func (h *handler) Patch(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	ctx, err := pkg.GetGrpcCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := department.NewPatchUpdateDepartmentRequest(rctx.PS.ByName("id"))
	if err := request.GetDataFromRequest(r, req.Data); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateDepartment(ctx.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.GetGrpcCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)
	qs := r.URL.Query()

	req := department.NewDescribeDepartmentRequest()

	req.Id = rctx.PS.ByName("id")
	req.WithSubCount = qs.Get("with_sub_count") == "true"
	req.WithUserCount = qs.Get("with_user_count") == "true"
	req.WithRole = qs.Get("with_role") == "true"
	ins, err := h.service.DescribeDepartment(ctx.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}

func (h *handler) GetSub(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.GetGrpcCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)
	pid := rctx.PS.ByName("id")

	req := department.NewQueryDepartmentRequestFromHTTP(r)
	req.ParentId = pid

	ins, err := h.service.QueryDepartment(ctx.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}

// DestroyPrimaryAccount 注销账号
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.GetGrpcCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)

	req := department.NewDeleteDepartmentRequestWithID(rctx.PS.ByName("id"))
	if _, err := h.service.DeleteDepartment(ctx.Context(), req); err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
	return
}
