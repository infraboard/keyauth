package http

import (
	"net/http"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/endpoint"
)

// CreateApplication 创建自定义角色
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	cid, cs := pkg.GetClientCredentialsFromHTTPRequest(r)
	if cid == "" && cs == "" {
		response.Failed(w, exception.NewBadRequest("service client credentials in header missed"))
		return
	}

	ctx, err := pkg.NewGrpcInCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	ctx.SetClientCredentials(cid, cs)
	req := endpoint.NewDefaultRegistryRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	_, err = pkg.Endpoint.Registry(
		ctx.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, req)
	return
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := endpoint.NewQueryEndpointRequestFromHTTP(r)

	var header, trailer metadata.MD
	set, err := h.endpoint.QueryEndpoints(
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
	return
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}
	rctx := context.GetContext(r)

	id := rctx.PS.ByName("id")
	req := endpoint.NewDescribeEndpointRequestWithID(id)

	var header, trailer metadata.MD
	d, err := h.endpoint.DescribeEndpoint(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
		return
	}

	response.Success(w, d)
	return
}
