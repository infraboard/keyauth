package http

import (
	"net/http"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/ip2region"
)

func (h *handler) UpdateDBFile(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req, err := ip2region.NewUploadFileRequestFromHTTP(r)
	if err != nil {
		response.Failed(w, exception.NewBadRequest("init request error, %s", err))
		return
	}
	req.WithToken(tk)

	err = h.service.UpdateDBFile(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "ok")
	return
}

func (h *handler) LoopupIP(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()

	ip := qs.Get("ip")
	if ip == "" {
		response.Failed(w, exception.NewBadRequest("ip need"))
		return
	}

	rc, err := h.service.LookupIP(ip)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, rc)
	return
}
