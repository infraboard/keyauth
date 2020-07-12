package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/user/types"
)

func (h *handler) CreateSubAccount(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := user.NewCreateUserRequest()
	req.WithToken(tk)
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	d, err := h.service.CreateAccount(types.SubAccount, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

func (h *handler) QuerySubAccount(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	page := request.NewPageRequestFromHTTP(r)
	req := user.NewQueryAccountRequest(page)
	req.WithToken(tk)
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	d, err := h.service.QueryAccount(types.SubAccount, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}
