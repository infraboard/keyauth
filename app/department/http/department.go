package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/app/department"
	"github.com/infraboard/keyauth/app/token"
)

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := department.NewQueryDepartmentRequestFromHTTP(r)
	req.Domain = tk.Domain

	apps, err := h.service.QueryDepartment(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, apps)
}

// Create 创建主账号
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := department.NewCreateDepartmentRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateOwner(tk)

	ins, err := h.service.CreateDepartment(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

// Create 创建主账号
func (h *handler) Put(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)
	req := department.NewPutUpdateDepartmentRequest(rctx.PS.ByName("id"))
	if err := request.GetDataFromRequest(r, req.Data); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateDepartment(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

// Create 创建主账号
func (h *handler) Patch(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)
	req := department.NewPatchUpdateDepartmentRequest(rctx.PS.ByName("id"))
	if err := request.GetDataFromRequest(r, req.Data); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateDepartment(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	rctx := context.GetContext(r)
	qs := r.URL.Query()

	req := department.NewDescribeDepartmentRequest()
	req.Domain = tk.Domain

	req.Id = rctx.PS.ByName("id")
	req.WithSubCount = qs.Get("with_sub_count") == "true"
	req.WithUserCount = qs.Get("with_user_count") == "true"
	req.WithRole = qs.Get("with_role") == "true"

	ins, err := h.service.DescribeDepartment(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) GetSub(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)
	pid := rctx.PS.ByName("id")

	req := department.NewQueryDepartmentRequestFromHTTP(r)
	req.ParentId = pid

	ins, err := h.service.QueryDepartment(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

// DestroyPrimaryAccount 注销账号
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	req := department.NewDeleteDepartmentRequestWithID(rctx.PS.ByName("id"))

	_, err := h.service.DeleteDepartment(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
}
