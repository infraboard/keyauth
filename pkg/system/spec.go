package system

import (
	"github.com/infraboard/keyauth/pkg/system/notify/mail"
	"github.com/infraboard/keyauth/pkg/system/notify/sms"
)

// Service 存储服务
type Service interface {
	UpdateEmail(*mail.Config) error
	UpdateSMS(*sms.Config) error
	GetConfig(version string) (*Config, error)
}
