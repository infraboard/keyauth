package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/common/session"
	"github.com/infraboard/keyauth/pkg/department"
)

// 创建部门加入申请
func (h *handler) CreateJoinApply(w http.ResponseWriter, r *http.Request) {
	ctx, err := session.GetTokenCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := department.NewJoinDepartmentRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.JoinDepartment(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}

// 查询部门加入申请
func (h *handler) QueryJoinApply(w http.ResponseWriter, r *http.Request) {
	ctx, err := session.GetTokenCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req, err := department.NewQueryApplicationFormRequestFromHTTP(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.QueryApplicationForm(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

// Create 创建主账号
func (h *handler) GetJoinApply(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)
	ctx, err := session.GetTokenCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := department.NewDescribeApplicationFormRequetWithID(rctx.PS.ByName("id"))
	ins, err := h.service.DescribeApplicationForm(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}

// Create 创建主账号
func (h *handler) DealJoinApply(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)
	ctx, err := session.GetTokenCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := department.NewDefaultDealApplicationFormRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.Id = rctx.PS.ByName("id")

	ins, err := h.service.DealApplicationForm(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}
