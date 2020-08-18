package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/geoip"
)

func (h *handler) UploadGEOIPDBFile(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := geoip.NewUploadFileRequestFromHTTP(r)
	req.WithToken(tk)

	err = h.service.UploadDBFile(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "ok")
	return
}

func (h *handler) LoopupIP(w http.ResponseWriter, r *http.Request) {
	response.Success(w, "ok")
	return
}
