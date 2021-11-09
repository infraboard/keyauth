package http

import (
	"net/http"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg/session"
	"github.com/infraboard/keyauth/pkg/token"
)

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req, err := session.NewQuerySessionRequestFromHTTP(r)
	if err != nil {
		response.Failed(w, exception.NewBadRequest("validate request error, %s", err))
		return
	}
	req.Domain = tk.Domain

	set, err := h.service.QuerySession(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)
	req := session.NewDescribeSessionRequestWithID(rctx.PS.ByName("id"))

	set, err := h.service.DescribeSession(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
