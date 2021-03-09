package http

import (
	"context"
	"net/http"

	httpCtx "github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/token"
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

	var header, trailer metadata.MD
	d, err := h.service.IssueToken(
		pkg.NewGrpcOutCtx().Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
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

	var header, trailer metadata.MD
	d, err := h.service.ValidateToken(
		context.Background(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
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

	var header, trailer metadata.MD
	_, err := h.service.RevolkToken(
		context.Background(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)

	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
		return
	}

	response.Success(w, "revolk ok")
	return
}

// QueryApplicationToken 获取应用访问凭证
func (h *handler) QueryApplicationToken(w http.ResponseWriter, r *http.Request) {
	rctx := httpCtx.GetContext(r)

	page := request.NewPageRequestFromHTTP(r)
	req := token.NewQueryTokenRequest(&page.PageRequest)
	req.ApplicationId = rctx.PS.ByName("id")

	var header, trailer metadata.MD
	tkSet, err := h.service.QueryToken(
		context.Background(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
		return
	}

	response.Success(w, tkSet)
	return
}
