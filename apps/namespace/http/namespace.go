package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/apps/namespace"
	"github.com/infraboard/keyauth/apps/token"
)

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	req := namespace.NewQueryNamespaceRequestFromHTTP(r)

	apps, err := h.service.QueryNamespace(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, apps)
}

func (h *handler) ListSelfNamespace(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := namespace.NewQueryNamespaceRequestFromHTTP(r)
	req.UpdateOwner(tk)

	apps, err := h.service.QueryNamespace(
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
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := namespace.NewCreateNamespaceRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateOwner(tk)

	d, err := h.service.CreateNamespace(
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

	qs := r.URL.Query()
	req := namespace.NewDescriptNamespaceRequest()
	req.Id = rctx.PS.ByName("id")
	req.WithDepartment = qs.Get("with_department") == "true"

	d, err := h.service.DescribeNamespace(
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
	rctx := context.GetContext(r)
	req := namespace.NewDeleteNamespaceRequestWithID(rctx.PS.ByName("id"))

	_, err := h.service.DeleteNamespace(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
}
