package sms_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/infraboard/keyauth/pkg/system/notify"
	"github.com/infraboard/keyauth/pkg/system/notify/sms"
)

func TestSendMail(t *testing.T) {
	should := assert.New(t)
	conf, err := sms.LoadConfigFromEnv()
	if should.NoError(err) {
		sd, err := sms.NewSender(conf)
		if should.NoError(err) {
			req := notify.NewSendSMSRequest()
			req.TemplateID = conf.Tencent.TemplateID
			req.ParamSet = []string{"409933", "10"}
			req.PhoneNumberSet = []string{"+8618108053819"}
			should.NoError(sd.Send(req))
		}
	}
}
