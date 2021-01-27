package http

import (
	"context"
	"net/http"

	httpctx "github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/common/session"
	"github.com/infraboard/keyauth/pkg/application"
)

func (h *handler) QueryUserApplication(w http.ResponseWriter, r *http.Request) {
	tk, err := session.GetTokenFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	page := request.NewPageRequestFromHTTP(r)
	req := application.NewQueryApplicationRequest(page)
	req.Account = tk.Account

	ctx := session.WithTokenContext(context.Background(), tk)
	apps, err := h.service.QueryApplication(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, apps)
	return
}

// CreateApplication 创建主账号
func (h *handler) CreateUserApplication(w http.ResponseWriter, r *http.Request) {
	tk, err := session.GetTokenFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := application.NewCreateApplicatonRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	ctx := session.WithTokenContext(context.Background(), tk)
	d, err := h.service.CreateUserApplication(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

func (h *handler) GetApplication(w http.ResponseWriter, r *http.Request) {
	tk, err := session.GetTokenFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := httpctx.GetContext(r)

	req := application.NewDescriptApplicationRequest()
	req.Id = rctx.PS.ByName("id")

	ctx := session.WithTokenContext(context.Background(), tk)

	d, err := h.service.DescribeApplication(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

// DestroyPrimaryAccount 注销账号
func (h *handler) DestroyApplication(w http.ResponseWriter, r *http.Request) {
	tk, err := session.GetTokenFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := httpctx.GetContext(r)

	req := application.NewDeleteApplicationRequestWithID(rctx.PS.ByName("id"))
	ctx := session.WithTokenContext(context.Background(), tk)

	if _, err := h.service.DeleteApplication(ctx, req); err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
	return
}
