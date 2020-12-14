package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/verifycode"
)

func (h *handler) IssueCodeByPass(w http.ResponseWriter, r *http.Request) {
	req := verifycode.NewIssueCodeRequestByPass()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	code, err := h.service.IssueCode(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, code)
	return
}

func (h *handler) IssueCodeByToken(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := verifycode.NewIssueCodeRequestByToken()
	req.WithToken(tk)
	code, err := h.service.IssueCode(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, code)
	return
}
