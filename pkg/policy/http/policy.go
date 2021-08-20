package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/policy"
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

	req := policy.NewQueryPolicyRequestFromHTTP(r)
	req.Domain = tk.Domain

	var header, trailer metadata.MD
	apps, err := h.service.QueryPolicy(
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
}

// CreateApplication 创建主账号
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 用户添加的策略都是自定义策略
	req := policy.NewCreatePolicyRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.Type = policy.PolicyType_CUSTOM

	var header, trailer metadata.MD
	d, err := h.service.CreatePolicy(
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

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}
	rctx := context.GetContext(r)

	req := policy.NewDescriptPolicyRequest()
	req.Id = rctx.PS.ByName("id")

	var header, trailer metadata.MD
	d, err := h.service.DescribePolicy(
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

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)
	tk, err := ctx.GetToken()
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := policy.NewDeletePolicyRequestWithID(rctx.PS.ByName("id"))
	req.Domain = tk.Domain

	var header, trailer metadata.MD
	_, err = h.service.DeletePolicy(
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
