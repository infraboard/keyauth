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
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	page := request.NewPageRequestFromHTTP(r)
	req := department.NewQueryDepartmentRequest(page)
	req.WithToken(tk)

	apps, err := h.service.QueryDepartment(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, apps)
	return
}

// Create 创建主账号
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := department.NewCreateDepartmentRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.WithToken(tk)

	ins, err := h.service.CreateDepartment(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	req := department.NewDescriptDepartmentRequest()
	req.ID = rctx.PS.ByName("id")
	ins, err := h.service.DescribeDepartment(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}

func (h *handler) GetSub(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)
	pid := rctx.PS.ByName("id")

	page := request.NewPageRequestFromHTTP(r)
	req := department.NewQueryDepartmentRequest(page)
	req.ParentID = &pid
	req.WithToken(tk)

	ins, err := h.service.QueryDepartment(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}

// DestroyPrimaryAccount 注销账号
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	if err := h.service.DeleteDepartment(rctx.PS.ByName("id")); err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
	return
}
