package sms

import (
	"fmt"

	"github.com/infraboard/keyauth/pkg/system/notify"
)

// NewSender todo
func NewSender(conf *Config) (notify.SMSSender, error) {
	switch conf.Provider {
	case ProviderTenCent:
		return newTenCentSMSSender(conf.TencentSMS)
	case ProviderALI:
		return nil, fmt.Errorf("not impl")
	default:
		return nil, fmt.Errorf("unknwon provier, %s", conf.Provider)
	}

}
