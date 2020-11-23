package sms

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator/v10"
)

const (
	// DEFAULT_TENCENT_SMS_ENDPOINT todo
	DEFAULT_TENCENT_SMS_ENDPOINT = "sms.tencentcloudapi.com"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// LoadTenCentSMSConfigFromEnv todo
func LoadTenCentSMSConfigFromEnv() (*TenCentSMS, error) {
	cfg := &TenCentSMS{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("load config from env, %s", err.Error())
	}
	return cfg, nil
}

// TenCentSMS todo
// 接口和相关文档请参考https://console.cloud.tencent.com/api/explorer?Product=sms&Version=2019-07-11&Action=SendSms&SignVersion=
type TenCentSMS struct {
	Endpoint   string `json:"endpoint" env:"K_SMS_TENCENT_ENDPOINT"`
	SecretID   string `json:"secret_id" validate:"required,lte=64" env:"K_SMS_TENCENT_SECRET_ID"`
	SecretKey  string `json:"secret_key" validate:"required,lte=64" env:"K_SMS_TENCENT_SECRET_KEY"`
	AppID      string `json:"app_id" validate:"required,lte=64" env:"K_SMS_TENCENT_APPID"`
	TemplateID string `json:"template_id" validate:"required,lte=64" env:"K_SMS_TENCENT_TEMPLATE_ID"`
	Sign       string `json:"sign" validate:"required,lte=128" env:"K_SMS_TENCENT_SIGN"`
}

// Validate todo
func (s *TenCentSMS) Validate() error {
	return validate.Struct(s)
}

// GetEndpoint todo
func (s *TenCentSMS) GetEndpoint() string {
	if s.Endpoint != "" {
		return s.Endpoint
	}

	return DEFAULT_TENCENT_SMS_ENDPOINT
}
