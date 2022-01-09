package http

import (
	"fmt"
	"net/http"

	"github.com/infraboard/keyauth/apps/domain"
	"github.com/infraboard/keyauth/apps/token"
	"github.com/infraboard/keyauth/apps/token/issuer"
	"github.com/infraboard/keyauth/apps/user/types"
	"github.com/infraboard/keyauth/apps/wxwork"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
)

// WechatWorkCheck todo 待添加企业微信配置数据后进行更新
// 扫码授权文档: https://work.weixin.qq.com/api/doc/90000/90135/91019
func (h *handler) WechatWorkCheck(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	service := values.Get("service")
	state := values.Get("state")
	clientId := values.Get("client_id")
	clientSecret := values.Get("client_secret")
	d := values.Get("domain")
	if d == "" {
		d = domain.AdminDomainName
	}

	i, err := issuer.NewTokenIssuer()
	if err != nil {
		response.Failed(w, err)
	}
	nt, err := i.IssueToken(r.Context(), &token.IssueTokenRequest{
		GrantType:    token.GrantType_WECHAT_WORK,
		ClientId:     clientId,
		ClientSecret: clientSecret,
		State:        state,
		Service:      service,
		Username:     d,
	})
	if err != nil {
		response.Failed(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("%s?token=%s", service, nt.AccessToken), http.StatusFound)
	return
}

// CreateConf 创建企业微信配置
func (h *handler) CreateConf(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := wxwork.NewSaveConfRequest()
	req.WithToken(tk)
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

// GetConf 获取企业微信配置
func (h *handler) GetConf(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := wxwork.NewDescribeConfWithDomain(tk.Domain)
	d, err := h.service.DescribeConfig(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	//d.Desensitize()
	response.Success(w, d)
}

func (h *handler) ListConf(w http.ResponseWriter, r *http.Request) {
	page := request.NewPageRequestFromHTTP(r)
	req := wxwork.NewQueryConfigRequest(page)

	apps, err := h.service.QueryConfig(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, apps)
}

func (h *handler) DestroyConfig(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)
	req := wxwork.NewDescribeConfWithDomain(tk.Domain)
	err := h.service.DeleteConfig(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
}
