package http

import (
	"fmt"
	"net"
	"net/http"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/geoip"
)

func (h *handler) UpdateDBFile(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req, err := geoip.NewUploadFileRequestFromHTTP(r)
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

	bitCount := uint(len(ip) * 8)
	fmt.Println(bitCount)

	response.Success(w, "ok")
	return
}
