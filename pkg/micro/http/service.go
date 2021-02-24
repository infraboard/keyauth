package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/common/session"
	"github.com/infraboard/keyauth/pkg/micro"
)

func (h *handler) QueryService(w http.ResponseWriter, r *http.Request) {
	page := request.NewPageRequestFromHTTP(r)
	req := micro.NewQueryMicroRequest(page)

	ctx, err := session.GetTokenCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	apps, err := h.service.QueryService(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, apps)
	return
}

func (h *handler) CreateService(w http.ResponseWriter, r *http.Request) {
	ctx, err := session.GetTokenCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := micro.NewCreateMicroRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	d, err := h.service.CreateService(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}

func (h *handler) GetService(w http.ResponseWriter, r *http.Request) {
	ctx, err := session.GetTokenCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}
	rctx := context.GetContext(r)

	req := micro.NewDescribeServiceRequest()
	req.Id = rctx.PS.ByName("id")

	d, err := h.service.DescribeService(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}

// DestroyService 销毁服务
func (h *handler) DestroyService(w http.ResponseWriter, r *http.Request) {
	ctx, err := session.GetTokenCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)
	req := micro.NewDeleteMicroRequestWithID(rctx.PS.ByName("id"))
	if _, err := h.service.DeleteService(ctx, req); err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
	return
}

func (h *handler) RefreshServiceClientCredential(w http.ResponseWriter, r *http.Request) {
	ctx, err := session.GetTokenCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}
	rctx := context.GetContext(r)

	req := micro.NewDescribeServiceRequest()
	req.Id = rctx.PS.ByName("id")

	d, err := h.service.RefreshServiceClientSecret(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}
