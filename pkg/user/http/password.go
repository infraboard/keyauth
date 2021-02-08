package http

import (
	"context"
	"net/http"

	"github.com/infraboard/keyauth/common/session"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	tk, err := session.GetTokenFromHTTPRequest(r)
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

	ctx := session.WithTokenContext(context.Background(), tk)
	pass, err := h.service.UpdateAccountPassword(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	pass.Password = ""
	response.Success(w, pass)
	return
}

func (h *handler) GeneratePassword(w http.ResponseWriter, r *http.Request) {
	tk, err := session.GetTokenFromHTTPRequest(r)
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

	ctx := session.WithTokenContext(context.Background(), tk)
	pass, err := h.service.GeneratePassword(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, pass)
	return
}
