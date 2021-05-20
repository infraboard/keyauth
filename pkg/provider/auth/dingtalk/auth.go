package dingtalk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/infraboard/keyauth/pkg/provider/auth"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	client = &http.Client{}
)

type Response struct {
	ErrCode  int64    `json:"errcode"`
	ErrMsg   string   `json:"errmsg"`
	UserInfo UserInfo `json:"user_info"`
}

type UserInfo struct {
	Nick    string `json:"nick"`
	OpenID  string `json:"openid"`
	UnionID string `json:"unionid"`
}

type Dingtalk struct {
	AppID     string `json:"app_id" bson:"app_id"`
	AppSecret string `json:"app_secret" bson:"app_secret"`
}

func (a *Dingtalk) getSignature(msg []byte) string {
	hmac := hmac.New(sha256.New, []byte(a.AppSecret))
	_, err := hmac.Write(msg)
	if err != nil {
		zap.L().Errorf("GetSignature hmac.Write error", err)
	}
	digest := hmac.Sum(nil)
	return url.QueryEscape(base64.StdEncoding.EncodeToString(digest))
}

func (a *Dingtalk) accessTokenURL() string {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	signature := a.getSignature([]byte(timestamp))
	accessTokenURL := fmt.Sprintf("https://oapi.dingtalk.com/sns/getuserinfo_bycode?accessKey=%s&timestamp=%s&signature=%s",
		a.AppID,
		timestamp,
		signature)
	return accessTokenURL
}

// https://ding-doc.dingtalk.com/doc#/serverapi3/mrugr3
// Step 1: To https://oapi.dingtalk.com/connect/qrconnect?appid=APPID&response_type=code&scope=snsapi_login&state=STATE&redirect_uri=REDIRECT_URI
// Step 2.2: Within Callback, get user_info.nick
// POST HTTPS with body { "tmp_auth_code": "23152698ea18304da4d0ce1xxxxx" }  == code ?
// https://oapi.dingtalk.com/sns/getuserinfo_bycode?accessKey=xxx&timestamp=xxx&signature=xxx
// accessKey=appid
// https://ding-doc.dingtalk.com/doc#/serverapi2/kymkv6
func (a *Dingtalk) CodeAuth(req *auth.AuthCodeRequest) error {
	body := fmt.Sprintf(`{"tmp_auth_code": "%s"}`, req.Code)
	request, _ := http.NewRequest("POST", a.accessTokenURL(), bytes.NewReader([]byte(body)))
	request.Header.Set("Content-Type", "application/json")
	// 发起请求
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 处理响应
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	zap.L().Debugf("dingding oauthcode request, req: %s [%s], response: %s", request, string(b))

	if (resp.StatusCode / 100) != 2 {
		return fmt.Errorf("status code: %d, %s", resp.StatusCode, string(b))
	}

	ins := Response{}
	err = json.Unmarshal(b, &ins)
	if err != nil {
		return err
	}

	return nil
}
