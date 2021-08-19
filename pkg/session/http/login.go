package http

import (
	"net/http"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/session"
)

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	tk, err := ctx.GetToken()
	if err != nil {
		response.Failed(w, err)
		return
	}

	req, err := session.NewQuerySessionRequestFromHTTP(r)
	if err != nil {
		response.Failed(w, exception.NewBadRequest("validate request error, %s", err))
		return
	}
	req.Domain = tk.Domain

	var header, trailer metadata.MD
	set, err := h.service.QuerySession(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
		return
	}
	response.Success(w, set)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)
	req := session.NewDescribeSessionRequestWithID(rctx.PS.ByName("id"))

	if err != nil {
		response.Failed(w, exception.NewBadRequest("validate request error, %s", err))
		return
	}

	var header, trailer metadata.MD
	set, err := h.service.DescribeSession(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
		return
	}
	response.Success(w, set)
}
