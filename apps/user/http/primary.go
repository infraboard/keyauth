package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/apps/user"
	"github.com/infraboard/keyauth/apps/user/types"
)

// CreatePrimayAccount 创建主账号
func (h *handler) CreatePrimayAccount(w http.ResponseWriter, r *http.Request) {

	req := user.NewCreateUserRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	d, err := h.service.CreateAccount(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	req.UserType = types.UserType_PRIMARY

	response.Success(w, d)
}

// DestroyPrimaryAccount 注销账号
func (h *handler) DestroyPrimaryAccount(w http.ResponseWriter, r *http.Request) {
	_, err := h.service.DeleteAccount(
		r.Context(),
		nil,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
}
