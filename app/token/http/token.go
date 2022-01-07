package http

import (
	"fmt"
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/app/token"
)

const (
	// CodeHeaderKeyName 认证码
	CodeHeaderKeyName = "X-Verify-Code"
)

// IssueToken 颁发资源访问令牌
func (h *handler) IssueToken(w http.ResponseWriter, r *http.Request) {
	req := token.NewIssueTokenRequest()
	req.WithUserAgent(r.UserAgent())
	req.WithRemoteIPFromHTTP(r)

	// 从Header中获取client凭证, 如果有
	req.ClientId, req.ClientSecret, _ = r.BasicAuth()
	req.VerifyCode = r.Header.Get(CodeHeaderKeyName)
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	d, err := h.service.IssueToken(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	if req.Service != "" {
		http.Redirect(w, r, fmt.Sprintf("%s?token=%s", req.Service, d.AccessToken), http.StatusFound)
		return
	}

	response.Success(w, d)
	return
}

// IssueToken 颁发资源访问令牌
func (h *handler) ValidateToken(w http.ResponseWriter, r *http.Request) {
	req := token.NewValidateTokenRequest()
	qs := r.URL.Query()

	req.AccessToken = r.Header.Get("X-OAUTH-TOKEN")
	req.EndpointId = qs.Get("endpoint_id")
	req.NamespaceId = qs.Get("namespace_id")

	d, err := h.service.ValidateToken(
		r.Context(),
		req,
	)
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
	req.ClientId, req.ClientSecret, _ = r.BasicAuth()

	_, err := h.service.RevolkToken(
		r.Context(),
		req,
	)

	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "revolk ok")
	return
}

// QueryToken 获取应用访问凭证
func (h *handler) QueryToken(w http.ResponseWriter, r *http.Request) {
	req, err := token.NewQueryTokenRequestFromHTTP(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	tkSet, err := h.service.QueryToken(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, tkSet)
	return
}

// QueryToken 获取应用访问凭证
func (h *handler) ChangeNamespace(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := token.NewChangeNamespaceRequest()
	req.Token = tk.AccessToken

	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	tkSet, err := h.service.ChangeNamespace(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, tkSet)
	return
}

// RevolkToken 撤销资源访问令牌
func (h *handler) DeleteToken(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := token.NewDeleteTokenRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.Domain = tk.Domain
	req.Account = tk.Account

	resp, err := h.service.DeleteToken(
		r.Context(),
		req,
	)

	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, resp)
	return
}
