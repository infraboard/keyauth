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
	tk, err := session.GetTokenFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)

	req := permission.NewQueryPermissionRequest(request.NewPageRequestFromHTTP(r))
	req.NamespaceID = rctx.PS.ByName("id")
	req.WithToken(tk)

	set, err := h.service.QueryPermission(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
	return
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	tk, err := session.GetTokenFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)

	req := permission.NewCheckPermissionrequest()
	req.NamespaceID = rctx.PS.ByName("id")
	req.EnpointID = rctx.PS.ByName("eid")
	req.WithToken(tk)

	d, err := h.service.CheckPermission(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}
