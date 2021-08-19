package http

import (
	"net/http"

	"github.com/infraboard/mcube/exception"
	httpctx "github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/application"
)

func (h *handler) QueryApplication(w http.ResponseWriter, r *http.Request) {
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

	page := request.NewPageRequestFromHTTP(r)
	req := application.NewQueryApplicationRequest(page)
	req.Account = tk.Account

	apps, err := h.service.QueryApplication(ctx.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, apps)
}

// CreateApplication 创建主账号
func (h *handler) CreateApplication(w http.ResponseWriter, r *http.Request) {
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

	req := application.NewCreateApplicatonRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateOwner(tk)

	var header, trailer metadata.MD
	d, err := h.service.CreateApplication(
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

func (h *handler) GetApplication(w http.ResponseWriter, r *http.Request) {
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

	rctx := httpctx.GetContext(r)
	req := application.NewDescriptApplicationRequest()
	req.Id = rctx.PS.ByName("id")

	var header, trailer metadata.MD
	ins, err := h.service.DescribeApplication(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)

	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
		return
	}

	if !ins.IsOwner(tk.Account) {
		response.Failed(w, exception.NewPermissionDeny("this application is not yours"))
		return
	}

	response.Success(w, ins)
}

// DestroyPrimaryAccount 注销账号
func (h *handler) DestroyApplication(w http.ResponseWriter, r *http.Request) {
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

	rctx := httpctx.GetContext(r)
	descReq := application.NewDescriptApplicationRequest()
	descReq.Id = rctx.PS.ByName("id")

	var header, trailer metadata.MD
	ins, err := h.service.DescribeApplication(
		ctx.Context(),
		descReq,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)

	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
		return
	}

	if !ins.IsOwner(tk.Account) {
		response.Failed(w, exception.NewPermissionDeny("this application is not yours"))
		return
	}

	rctx = httpctx.GetContext(r)
	req := application.NewDeleteApplicationRequestWithID(rctx.PS.ByName("id"))

	_, err = h.service.DeleteApplication(
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
}
