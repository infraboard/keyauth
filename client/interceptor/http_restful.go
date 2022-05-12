package interceptor

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/keyauth/apps/token"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/pb/http"
)

func GetEntryFromRouteReader(r restful.RouteReader) http.Entry {
	m := label.Meta(r.Metadata())
	return http.Entry{
		FunctionName:     r.Operation(),
		Path:             r.Path(),
		Resource:         m.Resource(),
		AuthEnable:       m.AuthEnable(),
		PermissionEnable: m.PermissionEnable(),
		Allow:            m.Allow(),
		AuditLog:         m.AuditEnable(),
		Labels: map[string]string{
			label.ActionLableKey: m.Action(),
		},
	}
}

func (a *HTTPAuther) RestfulAuthHandlerFunc(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	authInfo, err := a.Auth(req.Request, GetEntryFromRouteReader(req.SelectedRoute()))
	if err != nil {
		return
	}
	tk, ok := authInfo.(*token.Token)
	if ok {
		response.Failed(resp.ResponseWriter, exception.NewInternalServerError("auth info not *token.Token"))
		return
	}

	req.SetAttribute("token", tk)
	chain.ProcessFilter(req, resp)
}
