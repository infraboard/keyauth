package sms

// TenCentSMS todo
// 接口和相关文档请参考https://console.cloud.tencent.com/api/explorer?Product=sms&Version=2019-07-11&Action=SendSms&SignVersion=
type TenCentSMS struct {
	SecretID   string `json:"secret_id"`
	SecretKey  string `json:"secret_key"`
	AppID      string `json:"app_id"`
	TemplateID string `json:"template_id"`
	Sign       string `json:"sign"`
}

// Send todo
func (s *TenCentSMS) Send(to string, msg []byte)
