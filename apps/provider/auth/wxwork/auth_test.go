package wxwork_test

import (
	"fmt"
	"testing"

	"github.com/infraboard/keyauth/apps/provider/auth/wxwork"
)

func TestAuthCode(t *testing.T) {
	wx := wxwork.Wechat{
		AppID:     "wx891xxxxx",  // 企业微信app ID
		AppSecret: "84OrKLxIxxx", // 企业微信app secret
		AgentID:   "1000xxx",     // 企业微信 应用ID
	}
	token := wx.GetAccessToken()
	resp := wx.FromCodeGetUserInfo("", token) // code: oauth_code
	fmt.Printf("%+v\n", resp)
	fmt.Println("token: ", token)
}
