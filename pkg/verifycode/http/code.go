package http

import (
	"context"
	"net/http"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/common/session"
	"github.com/infraboard/keyauth/pkg/verifycode"
)

func (h *handler) IssueCodeByPass(w http.ResponseWriter, r *http.Request) {
	req := verifycode.NewIssueCodeRequestByPass()
	// 从Header中获取client凭证, 如果有
	req.ClientId, req.ClientSecret, _ = r.BasicAuth()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	req.IssueType = verifycode.IssueType_PASS
	code, err := h.service.IssueCode(context.Background(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, code)
	return
}

func (h *handler) IssueCodeByToken(w http.ResponseWriter, r *http.Request) {
	ctx, err := session.GetTokenCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := verifycode.NewIssueCodeRequestByToken()
	code, err := h.service.IssueCode(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, code)
	return
}
