package wxwork_test

import (
	"fmt"
	"github.com/infraboard/keyauth/app/provider/auth/wxwork"
	"testing"
)

func TestAuthCode(t *testing.T) {
	wx := wxwork.Wechat{
		AppID:     "wx8918a4299cc1b440",                          // 企业微信app ID
		AppSecret: "84OrKL_QXM1WCW_N0pAhtLLenbmFcsUetrDqymqAIH4", // 企业微信app secret
		AgentID:   "1000071",                                     // 企业微信 应用ID
	}
	token := wx.GetAccessToken()
	resp := wx.FromCodeGetUserInfo("", token) // code: oauth_code
	fmt.Printf("%+v\n", resp)
	fmt.Println("token: ", token)
}
