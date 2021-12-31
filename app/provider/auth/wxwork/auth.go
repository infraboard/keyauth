package wxwork

import (
	"encoding/json"
	"errors"
	"github.com/infraboard/mcube/logger/zap"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var (
	URLGetToken                   = "https://qyapi.weixin.qq.com/cgi-bin/gettoken"
	URLFromCodeGetUserInfo        = "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo"
	URLGetUserInfo         string = "https://qyapi.weixin.qq.com/cgi-bin/user/get"
)

type ScanCodeRequest struct {
	Service string `form:"service" binding:"required"`
	Code    string `form:"code" binding:"required"`
	State   string `form:"state" binding:"required"`
	AppID   string `form:"appid" binding:"required"`
}

type Response struct {
	ErrCode        int    `json:"errcode"`         // 返回码
	ErrMsg         string `json:"errmsg"`          // 对返回码的文本描述内容
	ExpiresIN      int64  `json:"expires_in"`      // 有效期
	AccessToken    string `json:"access_token"`    // 认证Token
	ExternalUserID string `json:"external_userid"` // 外部联系人id，当且仅当用户是企业的客户，且跟进人在应用的可见范围内时返回。
	// 如果是第三方应用调用，针对同一个客户，同一个服务商不同应用获取到的id相同
	UserInfo
}

type WechatUser struct {
	ErrCode      int    `json:"errcode"`    // 返回码
	ErrMsg       string `json:"errmsg"`     // 对返回码的文本描述内容
	UserID       string `json:"userid"`     // 用户ID
	Avatar       string `json:"avatar"`     // 头像
	Position     string `json:"position"`   // 职称
	Name         string `json:"name"`       // 别名
	Email        string `json:"email"`      // 邮箱
	Mobile       string `json:"mobile"`     // 电话
	DepartmentID []int  `json:"department"` // 部门IDs
	IsLeader     int    `json:"isleader"`   //是否为负责人
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

func NewAuth(appID, appSecret, agentID string) *Wechat {
	w := &Wechat{
		AppID:     appID,     // 企业微信app ID
		AppSecret: appSecret, // 企业微信app secret
		AgentID:   agentID,   // 企业微信 应用ID
	}
	w.AccessToken = w.GetAccessToken()
	return w
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
		if wr.ErrCode != 0 {
			zap.L().Errorf("GetAccessToken code: %v, msg: %v\n", wr.ErrCode, wr.ErrMsg)
			return ""
		}
		w.ExpiresIN = wr.ExpiresIN
		w.CreateTokenDate = time.Now().Unix()
		return wr.AccessToken
	}
	zap.L().Errorf("GetAccessToken empty Done")
	return ""
}

func (w Wechat) FromCodeGetUserInfo(code, accessToken string) *Response {
	if code == "" {
		zap.L().Errorf("args code error! ")
		return nil
	}
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

func (w *Wechat) CheckCallBack(param *ScanCodeRequest) (string, error) {
	token := w.GetAccessToken()
	if token == "" {
		return "", errors.New("get token error!")
	}

	scr := w.FromCodeGetUserInfo(param.Code, token)
	if scr == nil {
		return "", errors.New("get user info error from token!")
	}

	if scr.OpenID == "OPENID" {
		return "", errors.New("get openid is wrong! this not is  ID!")
	}
	return scr.UserID, nil
}

// GetUserInfo 说明文档: https://work.weixin.qq.com/api/doc/90000/90135/90196
func (w *Wechat) GetUserInfo(userID string) *WechatUser {
	req, err := http.NewRequest(http.MethodGet, URLGetUserInfo, nil)
	if err != nil {
		zap.L().Errorf("GetUserInfo http.NewRequest error", err)
		return nil
	}
	if w.AccessToken == "" {
		zap.L().Errorf("GetUserInfo token is empty!")
		return nil
	}
	params := url.Values{
		"access_token": []string{w.AccessToken},
		"userid":       []string{userID},
	}
	req.URL.RawQuery = params.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		zap.L().Errorf("GetUserInfo http.NewRequest Do error", err)
		return nil
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, _ := ioutil.ReadAll(resp.Body)
	scr := WechatUser{}
	if err := json.Unmarshal(body, &scr); err == nil {
		if scr.ErrCode != 0 {
			zap.L().Errorf("GetUserInfo error! code: %v, msg: %v", scr.ErrCode, scr.ErrMsg)
			return nil
		}
		return &scr
	}
	zap.L().Errorf("GetUserInfo json Unmarshal error", err)
	return nil
}
