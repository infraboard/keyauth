package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/app/policy"
	"github.com/infraboard/keyauth/app/token"
)

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := policy.NewQueryPolicyRequestFromHTTP(r)
	req.Domain = tk.Domain

	apps, err := h.service.QueryPolicy(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, apps)
}

// CreateApplication 创建主账号
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	// 用户添加的策略都是自定义策略
	req := policy.NewCreatePolicyRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.Type = policy.PolicyType_CUSTOM

	d, err := h.service.CreatePolicy(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	req := policy.NewDescriptPolicyRequest()
	req.Id = rctx.PS.ByName("id")

	d, err := h.service.DescribePolicy(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := policy.NewDeletePolicyRequestWithID(ctx.PS.ByName("id"))
	req.Domain = tk.Domain

	_, err := h.service.DeletePolicy(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
}
