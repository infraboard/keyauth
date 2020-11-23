package system

import (
	"github.com/infraboard/keyauth/pkg/system/notify/mail"
	"github.com/infraboard/keyauth/pkg/system/notify/sms"
)

// Config 系统配置
type Config struct {
	Email      *mail.Config    `bson:"email" json:"email"`
	TencentSMS *sms.TenCentSMS `bson:"tencent_sms" json:"tencent_sms"`
	ALISMS     *sms.ALISMS     `bson:"ali_sms" json:"ali_sms"`
}
