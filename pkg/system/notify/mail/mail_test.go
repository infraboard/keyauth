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
			req.Content = "发送的邮件内容包含了未被许可的信息，或被系统识别为垃圾邮件。请检查是否有用户发送病毒或者垃圾邮件"

			should.NoError(sd.Send(req))
		}
	}
}
