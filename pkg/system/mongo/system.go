package mongo

import (
	"github.com/infraboard/keyauth/pkg/system"
	"github.com/infraboard/keyauth/pkg/system/notify/mail"
	"github.com/infraboard/keyauth/pkg/system/notify/sms"
)

func (s *service) UpdateEmail(*mail.Config) error {
	return nil
}

func (s *service) UpdateSMS(*sms.Config) error {
	return nil
}

func (s *service) GetConfig() (*system.Config, error) {
	return nil, nil
}
