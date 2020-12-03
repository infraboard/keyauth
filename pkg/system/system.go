package system

import (
	"github.com/infraboard/keyauth/pkg/system/notify/mail"
	"github.com/infraboard/keyauth/pkg/system/notify/sms"
)

// Config 系统配置
type Config struct {
	Email *mail.Config `bson:"email" json:"email"`
	SMS   *sms.Config  `bson:"sms" json:"sms"`
}
