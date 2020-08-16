package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
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

	req := user.NewNewQueryAccountRequestFromHTTP(r)
	req.WithToken(tk)

	d, err := h.service.QueryAccount(types.SubAccount, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

func (h *handler) DescribeSubAccount(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	req := user.NewDescriptAccountRequestWithAccount(rctx.PS.ByName("account"))
	d, err := h.service.DescribeAccount(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

// DestroySubAccount 注销账号
func (h *handler) DestroySubAccount(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	if err := h.service.DeleteAccount(rctx.PS.ByName("account")); err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
	return
}
