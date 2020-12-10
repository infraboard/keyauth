package verifycode

import (
	"github.com/go-playground/validator/v10"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// Code todo
type Code struct {
	Number string `bson:"number" json:"number"`
}

// NewDefaultConfig todo
func NewDefaultConfig() *Config {
	return &Config{
		ExpireMinutes: 10,
		MailTemplate:  "您的动态验证码为：{1}，{2}分钟内有效！，如非本人操作，请忽略本短信！",
	}
}

// Config todo
type Config struct {
	ExpireMinutes uint   `json:"expire_minutes" validate:"required,gte=10,lte=600"` // 验证码默认过期时间
	MailTemplate  string `json:"mail_template"`                                     // 邮件通知时的模板
	SmsTemplateID string `json:"sms_template_id"`                                   // 短信通知时的云商模板ID
}

// Validate todo
func (conf *Config) Validate() error {
	return validate.Struct(conf)
}
