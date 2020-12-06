package mail_test

import (
	"testing"

	"github.com/infraboard/keyauth/pkg/system/notify"
	"github.com/infraboard/keyauth/pkg/system/notify/mail"
	"github.com/stretchr/testify/assert"
)

func TestSendMail(t *testing.T) {
	should := assert.New(t)
	conf, err := mail.LoadConfigFromEnv()
	if should.NoError(err) {
		sd, err := mail.NewSender(conf)
		if should.NoError(err) {
			req := notify.NewSendMailRequest()
			req.To = "719118794@qq.com"
			req.Subject = "欢迎使用"
			req.Content = "请不要回复"
			should.NoError(sd.Send(req))
		}
	}
}
