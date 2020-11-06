package http

import (
	"net/http"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/department"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) CreateJoinApply(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := department.NewQueryDepartmentRequestFromHTTP(r)
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
func (h *handler) QueryJoinApply(w http.ResponseWriter, r *http.Request) {
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

// Create 创建主账号
func (h *handler) DealJoinApply(w http.ResponseWriter, r *http.Request) {
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
