package verifycode

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewCode todo
func NewCode(req *IssueCodeRequest) (*Code, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate issue code request error, %s", err)
	}

	cr := CheckCodeRequest{
		Number:   GenRandomCode(6),
		Username: req.Account(),
	}

	c := &Code{
		ID:               cr.HashID(),
		CheckCodeRequest: cr,
		IssueAt:          ftime.Now(),
		ExpiredMinite:    10,
	}

	return c, nil
}

// NewDefaultCode todo
func NewDefaultCode() *Code {
	return &Code{}
}

// Code todo
type Code struct {
	ID               string `bson:"_id" json:"id"`
	CheckCodeRequest `bson:",inline"`
	IssueAt          ftime.Time `bson:"issue_at" json:"issue_at"`
	ExpiredMinite    uint       `bson:"expired_minite" json:"expired_minite"`
}

// IsExpired todo
func (c *Code) IsExpired() bool {
	return time.Now().Sub(c.IssueAt.T()).Minutes() > float64(c.ExpiredMinite)
}

// NewDefaultConfig todo
func NewDefaultConfig() *Config {
	return &Config{
		NotifyType:    NotifyTypeMail,
		ExpireMinutes: 10,
		MailTemplate:  "您的动态验证码为：{1}，{2}分钟内有效！，如非本人操作，请忽略本短信！",
	}
}

// Config todo
type Config struct {
	NotifyType    NotifyType `json:"notify_type"`
	ExpireMinutes uint       `json:"expire_minutes" validate:"required,gte=10,lte=600"` // 验证码默认过期时间
	MailTemplate  string     `json:"mail_template"`                                     // 邮件通知时的模板
	SmsTemplateID string     `json:"sms_template_id"`                                   // 短信通知时的云商模板ID
}

// Validate todo
func (conf *Config) Validate() error {
	return validate.Struct(conf)
}

// GenRandomCode todo
func GenRandomCode(length uint) string {
	numbers := []string{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < int(length); i++ {
		c := rand.Intn(9)
		// 第一位不能为0
		if c == 0 {
			c = 1
		}

		numbers = append(numbers, strconv.Itoa(c))
	}

	return strings.Join(numbers, "")
}
