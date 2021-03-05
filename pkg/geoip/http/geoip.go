package http

import (
	"net"
	"net/http"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg/geoip"
)

func (h *handler) UpdateDBFile(w http.ResponseWriter, r *http.Request) {
	req, err := geoip.NewUploadFileRequestFromHTTP(r)
	if err != nil {
		response.Failed(w, exception.NewBadRequest("init request error, %s", err))
		return
	}

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

	ipStr := qs.Get("ip")
	if ipStr == "" {
		response.Failed(w, exception.NewBadRequest("ip need"))
		return
	}

	ip := net.ParseIP(ipStr)
	if ip == nil {
		response.Failed(w, exception.NewBadRequest("ip not validate"))
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
