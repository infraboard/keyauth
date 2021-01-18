package http

import (
	"context"
	"net/http"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/session"
)

func (h *handler) QueryLoginLog(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req, err := session.NewQuerySessionRequestFromHTTP(r)
	if err != nil {
		response.Failed(w, exception.NewBadRequest("validate request error, %s", err))
		return
	}

	ctx := pkg.WithTokenContext(context.Background(), tk)
	set, err := h.service.QuerySession(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
	return
}
