package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/common/session"
	"github.com/infraboard/keyauth/pkg/permission"
)

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	ctx, err := session.GetTokenCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)

	req := permission.NewQueryPermissionRequest(request.NewPageRequestFromHTTP(r))
	req.NamespaceId = rctx.PS.ByName("id")

	set, err := h.service.QueryPermission(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
	return
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx, err := session.GetTokenCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)

	req := permission.NewCheckPermissionrequest()
	req.NamespaceId = rctx.PS.ByName("id")
	req.EndpointId = rctx.PS.ByName("eid")

	d, err := h.service.CheckPermission(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}
