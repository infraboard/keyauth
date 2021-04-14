package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/mconf"
)

func (h *handler) QueryGroup(w http.ResponseWriter, r *http.Request) {
	page := request.NewPageRequestFromHTTP(r)
	req := mconf.NewQueryGroupRequest(page)

	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	var header, trailer metadata.MD
	apps, err := h.service.QueryGroup(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
		return
	}

	response.Success(w, apps)
	return
}

func (h *handler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := mconf.NewCreateGroupRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	var header, trailer metadata.MD
	d, err := h.service.CreateGroup(
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
}

// DestroyService 销毁服务
func (h *handler) DestroyGroup(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)
	req := mconf.NewDeleteGroupRequestWithName(rctx.PS.ByName("name"))

	var header, trailer metadata.MD
	_, err = h.service.DeleteGroup(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
		return
	}

	response.Success(w, "delete ok")
	return
}
