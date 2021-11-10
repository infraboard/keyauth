package http

import (
	"net/http"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/app/provider"
	"github.com/infraboard/keyauth/app/token"
	"github.com/infraboard/keyauth/app/user/types"
)

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	page := request.NewPageRequestFromHTTP(r)
	req := provider.NewQueryLDAPConfigRequest(page)

	apps, err := h.service.QueryConfig(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, apps)
}

// CreateApplication 创建主账号
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := provider.NewSaveLDAPConfigRequest()
	req.WithToken(tk)
	req.GetDryRunParamFromHTTP(r)
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	if !tk.UserType.IsIn(types.UserType_SUPPER, types.UserType_PRIMARY) {
		response.Failed(w, exception.NewPermissionDeny("只有域管理员可以设置域的LDAP"))
		return
	}

	d, err := h.service.SaveConfig(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := provider.NewDescribeLDAPConfigWithDomain(tk.Domain)
	d, err := h.service.DescribeConfig(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	d.Desensitize()
	response.Success(w, d)
}

func (h *handler) Check(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := provider.NewDescribeLDAPConfigWithDomain(tk.Domain)
	if err := h.service.CheckConnect(req); err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "passed")
}
