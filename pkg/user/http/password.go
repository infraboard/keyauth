package http

import (
	"net/http"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 解析需要更新的数据
	req := user.NewUpdatePasswordRequest()
	req.WithToken(tk)
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.Account = tk.Account

	pass, err := h.service.UpdateAccountPassword(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	pass.Password = ""
	response.Success(w, pass)
	return
}
