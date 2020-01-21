package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg/token"
)

// IssueToken 颁发资源访问令牌
func (h *handler) IssueToken(w http.ResponseWriter, r *http.Request) {
	req := token.NewIssueTokenRequest()

	// 从Header中获取client凭证, 如果有
	req.ClientID, req.ClientSecret, _ = r.BasicAuth()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	d, err := h.service.IssueToken(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

// IssueToken 颁发资源访问令牌
func (h *handler) ValidateToken(w http.ResponseWriter, r *http.Request) {
	req := token.NewValidateTokenRequest()

	_, _, _ = r.BasicAuth()
	req.AccessToken = r.Header.Get("X-OAUTH-TOKEN")
	req.Endpoint = r.URL.Query().Get("endpoint")

	d, err := h.service.ValidateToken(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

// RevolkToken 撤销资源访问令牌
func (h *handler) RevolkToken(w http.ResponseWriter, r *http.Request) {
	req := token.NewRevolkTokenRequest("", "")
	req.AccessToken = r.Header.Get("X-OAUTH-TOKEN")

	if err := h.service.RevolkToken(req); err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "revolk ok")
	return
}

// QueryApplicationToken 获取应用访问凭证
func (h *handler) QueryApplicationToken(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	page := request.NewPageRequestFromHTTP(r)
	req := token.NewQueryTokenRequest(page)
	req.ApplicationID = rctx.PS.ByName("id")

	tkSet, err := h.service.QueryToken(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, tkSet)
	return
}
