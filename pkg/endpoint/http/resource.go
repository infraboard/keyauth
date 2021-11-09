package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/endpoint"
)

func (h *handler) ListResource(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := endpoint.NewQueryResourceRequestFromHTTP(r)

	set, err := h.endpoint.QueryResources(
		ctx.Context(),
		req,
	)

	if len(set.Items) == 0 {
		set.Items = []*endpoint.Resource{}
	}

	response.Success(w, set)
	return
}
