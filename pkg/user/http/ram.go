package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg/user"
)

func (h *handler) CreateRAMAccount(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)
	did := rctx.PS.ByName("did")

	req := user.NewCreateUserRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	d, err := h.service.CreateRAMAccount(did, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}
