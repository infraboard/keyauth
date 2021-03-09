package http

import (
	"net/http"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	tk, err := ctx.GetToken()
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 解析需要更新的数据
	req := user.NewUpdatePasswordRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.Account = tk.Account

	pass, err := h.service.UpdateAccountPassword(ctx.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	pass.Password = ""
	response.Success(w, pass)
	return
}

func (h *handler) GeneratePassword(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 解析需要更新的数据
	req := user.NewGeneratePasswordRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	pass, err := h.service.GeneratePassword(ctx.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, pass)
	return
}
