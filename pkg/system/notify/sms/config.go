package sms

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/keyauth/pkg/system/notify"
)

const (
	// DEFAULT_TENCENT_SMS_ENDPOINT todo
	DEFAULT_TENCENT_SMS_ENDPOINT = "sms.tencentcloudapi.com"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// LoadConfigFromEnv todo
func LoadConfigFromEnv() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("load config from env, %s", err.Error())
	}
	return cfg, nil
}

// NewDefaultConfig todo
func NewDefaultConfig() *Config {
	return &Config{
		Provider:   ProviderTenCent,
		TencentSMS: &TenCentSMS{},
		ALISMS:     &ALISMS{},
	}
}

// Config 短信配置
type Config struct {
	Provider   Provider    `bson:"provider" json:"provider"`
	TencentSMS *TenCentSMS `bson:"tencent_sms" json:"tencent_sms"`
	ALISMS     *ALISMS     `bson:"ali_sms" json:"ali_sms"`
}

// Validate todo
func (c *Config) Validate() error {
	switch c.Provider {
	case ProviderTenCent:
		return c.TencentSMS.Validate()
	case ProviderALI:
		return fmt.Errorf("not impl")
	default:
		return fmt.Errorf("unknown provider type: %s", c.Provider)
	}
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

// NewDeaultTestSendRequest todo
func NewDeaultTestSendRequest() *TestSendRequest {
	return &TestSendRequest{
		SendSMSRequest: notify.NewSendSMSRequest(),
		Config:         NewDefaultConfig(),
	}
}

// TestSendRequest todo
type TestSendRequest struct {
	*notify.SendSMSRequest
	*Config
}

// Send todo
func (req *TestSendRequest) Send() error {
	sd, err := NewSender(req.Config)
	if err != nil {
		return err
	}

	return sd.Send(req.SendSMSRequest)
}
