package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/app/department"
	"github.com/infraboard/keyauth/app/token"
)

// 创建部门加入申请
func (h *handler) CreateJoinApply(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := department.NewJoinDepartmentRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateOwner(tk)

	ins, err := h.service.JoinDepartment(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

// 查询部门加入申请
func (h *handler) QueryJoinApply(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req, err := department.NewQueryApplicationFormRequestFromHTTP(r)
	if err != nil {
		response.Failed(w, err)
		return
	}
	req.Domain = tk.Domain

	ins, err := h.service.QueryApplicationForm(
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
func (h *handler) GetJoinApply(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := department.NewDescribeApplicationFormRequetWithID(ctx.PS.ByName("id"))
	req.Domain = tk.Domain

	ins, err := h.service.DescribeApplicationForm(
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
func (h *handler) DealJoinApply(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	req := department.NewDefaultDealApplicationFormRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.Id = rctx.PS.ByName("id")

	ins, err := h.service.DealApplicationForm(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
