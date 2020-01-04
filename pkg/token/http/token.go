package http

import (
	"net/http"

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
