package http

import (
	"net/http"

	"github.com/infraboard/keyauth/common/session"
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) ListResource(w http.ResponseWriter, r *http.Request) {
	ctx, err := session.GetTokenCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := endpoint.NewQueryResourceRequestFromHTTP(r)
	set, err := h.endpoint.QueryResources(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
	return
}
