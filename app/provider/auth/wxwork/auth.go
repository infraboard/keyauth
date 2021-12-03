package wxwork

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/infraboard/mcube/logger/zap"
)

var (
	URLGetToken            = "https://qyapi.weixin.qq.com/cgi-bin/gettoken"
	URLFromCodeGetUserInfo = "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo"
)

type Response struct {
	ErrCode        int    `json:"errcode"`         // 返回码
	ErrMsg         string `json:"errmsg"`          // 对返回码的文本描述内容
	ExpiresIN      int64  `json:"expires_in"`      // 有效期
	AccessToken    string `json:"access_token"`    // 认证Token
	ExternalUserID string `json:"external_userid"` // 外部联系人id，当且仅当用户是企业的客户，且跟进人在应用的可见范围内时返回。
	// 如果是第三方应用调用，针对同一个客户，同一个服务商不同应用获取到的id相同
	UserInfo
}

type UserInfo struct {
	UserID   string `json:"UserId"`   // 成员UserID
	OpenID   string `json:"OpenId"`   // 非企业成员的标识，对当前企业唯一
	DeviceID string `json:"DeviceId"` // 手机设备号
}

type Wechat struct {
	AppID           string `json:"app_id" bson:"app_id"`
	AgentID         string `json:"agent_id" bson:"agent_id"`
	AppSecret       string `json:"app_secret" bson:"app_secret"`
	AccessToken     string
	ExpiresIN       int64 `json:"expires_in"`        // 有效期
	CreateTokenDate int64 `json:"create_token_date"` // 创建时时间
}

func (w *Wechat) GetAccessToken() string {
	req, err := http.NewRequest(http.MethodGet, URLGetToken, nil)
	if err != nil {
		zap.L().Errorf("GetAccessToken http.NewRequest error", err)
	}
	params := url.Values{
		"corpid":     []string{w.AppID},
		"corpsecret": []string{w.AppSecret},
	}
	req.URL.RawQuery = params.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		zap.L().Errorf("GetAccessToken http.NewRequest Do error", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, _ := ioutil.ReadAll(resp.Body)
	wr := Response{}
	if err := json.Unmarshal(body, &wr); err == nil {
		w.ExpiresIN = wr.ExpiresIN
		w.CreateTokenDate = time.Now().Unix()
		return wr.AccessToken
	}
	zap.L().Errorf("GetAccessToken empty Done")
	return ""
}

func (w Wechat) FromCodeGetUserInfo(code, accessToken string) *Response {
	// 通过扫码回调函数传回的code等信息获取用户信息
	req, err := http.NewRequest(http.MethodGet, URLFromCodeGetUserInfo, nil)
	if err != nil {
		zap.L().Errorf("FromCodeGetUserInfo http.NewRequest error", err)
	}
	params := url.Values{
		"access_token": []string{accessToken},
		"code":         []string{code},
	}
	req.URL.RawQuery = params.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		zap.L().Errorf("FromCodeGetUserInfo http.NewRequest Do error", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, _ := ioutil.ReadAll(resp.Body)
	scr := Response{}
	if err := json.Unmarshal(body, &scr); err == nil {
		return &scr
	}
	zap.L().Errorf("FromCodeGetUserInfo json.Unmarshal error", err)
	return nil
}
