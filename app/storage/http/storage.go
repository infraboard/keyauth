package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/app/storage"
	"github.com/infraboard/keyauth/app/token"
)

func (h *handler) UploadGEOIPDBFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := storage.NewUploadFileRequestFromHTTP(r)
	req.BucketName = ctx.PS.ByName("name")
	req.WithToken(tk)

	err := h.service.UploadFile(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "ok")
	return
}
