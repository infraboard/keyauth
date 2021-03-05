package http

import (
	"net/http"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) ListResource(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.GetGrpcCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := endpoint.NewQueryResourceRequestFromHTTP(r)
	set, err := h.endpoint.QueryResources(ctx.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
	return
}
