package sms

import (
	"fmt"

	"github.com/infraboard/keyauth/pkg/system/notify"
)

// NewSender todo
func NewSender(conf *Config) (notify.SMSSender, error) {
	switch conf.EnabledProvider {
	case ProviderTenCent:
		return newTenCentSMSSender(conf.Tencent)
	case ProviderALI:
		return nil, fmt.Errorf("not impl")
	default:
		return nil, fmt.Errorf("unknwon provier, %s", conf.EnabledProvider)
	}

}
