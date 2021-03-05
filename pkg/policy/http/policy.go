package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/policy"
)

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.GetGrpcCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := policy.NewQueryPolicyRequestFromHTTP(r)

	apps, err := h.service.QueryPolicy(ctx.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, apps)
	return
}

// CreateApplication 创建主账号
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.GetGrpcCtxFromHTTPRequest(r)
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

	d, err := h.service.CreatePolicy(ctx.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.GetGrpcCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}
	rctx := context.GetContext(r)

	req := policy.NewDescriptPolicyRequest()
	req.Id = rctx.PS.ByName("id")
	d, err := h.service.DescribePolicy(ctx.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.GetGrpcCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)

	req := policy.NewDeletePolicyRequestWithID(rctx.PS.ByName("id"))
	if _, err := h.service.DeletePolicy(ctx.Context(), req); err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
	return
}
