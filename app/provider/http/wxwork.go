package http

import (
	"fmt"
	"github.com/infraboard/keyauth/app/token"
	"github.com/infraboard/keyauth/app/token/issuer"
	"github.com/infraboard/mcube/http/response"
	"net/http"
)

// WechatCheck todo 待添加企业微信配置数据后进行更新
// 扫码授权文档: https://work.weixin.qq.com/api/doc/90000/90135/91019
func (h *handler) WechatCheck(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	service := values.Get("service")
	state := values.Get("state")
	i, err := issuer.NewTokenIssuer()
	if err != nil {
		response.Failed(w, err)
	}
	nt, err := i.IssueToken(r.Context(), &token.IssueTokenRequest{
		GrantType: token.GrantType_WECHAT_WORK,
		ClientId: "gmEZgiNk7t11sH0pcuCibi1S",
		ClientSecret: "CilKesb1UvL0MJ5PAL1Kxm1VhejGFEp7",
		State: state,
		Service: service,
	})
	if err != nil {
		response.Failed(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("%s?token=%s",service, nt.AccessToken), http.StatusFound)
	return
}
