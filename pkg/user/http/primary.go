package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/user/types"
)

// CreatePrimayAccount 创建主账号
func (h *handler) CreatePrimayAccount(w http.ResponseWriter, r *http.Request) {
	req := user.NewCreateUserRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	d, err := h.service.CreateAccount(types.UserType_PRIMARY, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

// DestroyPrimaryAccount 注销账号
func (h *handler) DestroyPrimaryAccount(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	if err := h.service.DeleteAccount(tk.Account); err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
	return
}
