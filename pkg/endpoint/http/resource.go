package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

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

	var header, trailer metadata.MD
	set, err := h.endpoint.QueryResources(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
		return
	}

	if len(set.Items) == 0 {
		set.Items = []*endpoint.Resource{}
	}

	response.Success(w, set)
	return
}
