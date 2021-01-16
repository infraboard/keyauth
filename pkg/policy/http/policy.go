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
	tk, err := pkg.GetTokenFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := policy.NewQueryPolicyRequestFromHTTP(r)
	req.WithToken(tk)

	apps, err := h.service.QueryPolicy(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, apps)
	return
}

// CreateApplication 创建主账号
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromHTTPRequest(r)
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
	req.Type = policy.CustomPolicy
	req.WithToken(tk)

	d, err := h.service.CreatePolicy(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	req := policy.NewDescriptPolicyRequest()
	req.ID = rctx.PS.ByName("id")
	d, err := h.service.DescribePolicy(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)

	req := policy.NewDeletePolicyRequestWithID(rctx.PS.ByName("id"))
	req.WithToken(tk)
	if err := h.service.DeletePolicy(req); err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
	return
}
