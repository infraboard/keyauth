package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg/permission"
)

func (h *handler) ListRole(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)
	req := permission.NewQueryRoleRequest(rctx.PS.ByName("id"))

	set, err := h.service.QueryRole(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) ListPermission(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	req := permission.NewQueryPermissionRequest(request.NewPageRequestFromHTTP(r))
	req.NamespaceId = rctx.PS.ByName("id")

	set, err := h.service.QueryPermission(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) CheckPermission(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	req := permission.NewCheckPermissionRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.NamespaceId = rctx.PS.ByName("id")

	d, err := h.service.CheckPermission(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}
