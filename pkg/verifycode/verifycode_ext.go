package verifycode

import (
	"fmt"
	"hash/fnv"
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

	c := &Code{
		Number:        GenRandomCode(6),
		Username:      req.Account(),
		IssueAt:       ftime.Now().Timestamp(),
		ExpiredMinite: 10,
	}

	c.Id = HashID(c.Username, c.Number)
	return c, nil
}

// NewDefaultCode todo
func NewDefaultCode() *Code {
	return &Code{}
}

// IsExpired todo
func (c *Code) IsExpired() bool {
	return time.Now().Sub(time.Unix(c.IssueAt/1000, 0)).Minutes() > float64(c.ExpiredMinite)
}

// ExpiredMiniteString todo
func (c *Code) ExpiredMiniteString() string {
	return fmt.Sprintf("%d", c.ExpiredMinite)
}

// HashID todo
func HashID(username, number string) string {
	hash := fnv.New32a()
	hash.Write([]byte(username))
	hash.Write([]byte(number))
	return fmt.Sprintf("%x", hash.Sum32())
}

// NewDefaultConfig todo
func NewDefaultConfig() *Config {
	return &Config{
		NotifyType:    NotifyType_MAIL,
		ExpireMinutes: 10,
		MailTemplate:  "您的动态验证码为：{1}，{2}分钟内有效！，如非本人操作，请忽略本邮件！",
	}
}

// Config todo
type Config struct {
	NotifyType    NotifyType `bson:"notify_type" json:"notify_type"`
	ExpireMinutes uint       `bson:"expire_minutes" json:"expire_minutes" validate:"required,gte=10,lte=600"` // 验证码默认过期时间
	MailTemplate  string     `bson:"mail_template" json:"mail_template"`                                      // 邮件通知时的模板
	SmsTemplateID string     `bson:"sms_template_id" json:"sms_template_id"`                                  // 短信通知时的云商模板ID
}

// RenderMailTemplate todo
func (conf *Config) RenderMailTemplate(number, expiiredMinite string) string {
	t1 := strings.ReplaceAll(conf.MailTemplate, "{1}", number)
	return strings.ReplaceAll(t1, "{2}", expiiredMinite)
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
