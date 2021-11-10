package system

import (
	"github.com/infraboard/keyauth/app/system/notify/mail"
	"github.com/infraboard/keyauth/app/system/notify/sms"
	"github.com/infraboard/keyauth/app/verifycode"
)

// Service 存储服务
type Service interface {
	UpdateEmail(*mail.Config) error
	UpdateSMS(*sms.Config) error
	UpdateVerifyCode(*verifycode.Config) error
	GetConfig() (*Config, error)
	InitConfig(*Config) error
}
