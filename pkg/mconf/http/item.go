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

func (h *handler) QueryItem(w http.ResponseWriter, r *http.Request) {
	page := request.NewPageRequestFromHTTP(r)
	req := mconf.NewQueryItemRequest(page)

	rctx := context.GetContext(r)
	req.GroupName = rctx.PS.ByName("name")

	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	var header, trailer metadata.MD
	apps, err := h.service.QueryItem(
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

func (h *handler) AddItemToGroup(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := mconf.NewAddItemToGroupRequest()
	rctx := context.GetContext(r)
	req.GroupName = rctx.PS.ByName("name")
	if err := request.GetDataFromRequest(r, &req.Items); err != nil {
		response.Failed(w, err)
		return
	}

	var header, trailer metadata.MD
	d, err := h.service.AddItemToGroup(
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
