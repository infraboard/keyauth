package wxwork_test

import (
	"fmt"
	"github.com/infraboard/keyauth/app/provider/auth/wxwork"
	"testing"
)

func TestAuthCode(t *testing.T) {
	wx := wxwork.Wechat{
		AppID:     "wx8xxx",  // 企业微信app ID
		AppSecret: "84Orxxxx4", // 企业微信app secret
		AgentID:   "100xxxx",     // 企业微信 应用ID
	}
	token := wx.GetAccessToken()
	resp := wx.FromCodeGetUserInfo("", token) // code: oauth_code
	fmt.Printf("%+v\n", resp)
	fmt.Println("token: ", token)
}
