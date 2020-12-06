package mail

import (
	"errors"
	"fmt"
	"strings"

	"github.com/caarlos0/env/v6"
	"github.com/infraboard/keyauth/common/tls"
	"github.com/infraboard/keyauth/pkg/system/notify"
)

// LoadConfigFromEnv todo
func LoadConfigFromEnv() (*Config, error) {
	cfg := &Config{TLSConfig: &tls.Config{}}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("load config from env, %s", err.Error())
	}
	return cfg, nil
}

// NewEmailConfig todo
func NewEmailConfig(host, user, pass string) *Config {
	return &Config{
		Host:         host,
		AuthUserName: user,
		AuthPassword: pass,
		TLSConfig:    &tls.Config{},
	}
}

// NewDefaultConfig todo
func NewDefaultConfig() *Config {
	return &Config{
		TLSConfig: &tls.Config{},
	}
}

// Config todo
type Config struct {
	Host         string      `bson:"host" json:"host" env:"K_EMAIL_HOST"`
	AuthUserName string      `bson:"username" json:"username" env:"K_EMAIL_USERNAME"`
	AuthPassword string      `bson:"password" json:"password,omitempty" env:"K_EMAIL_PASSWORD"`
	AuthSecret   string      `bson:"secret" json:"secret,omitempty" env:"K_EMAIL_SECRET"`
	AuthIdentity string      `bson:"identity" json:"identity,omitempty" env:"K_EMAIL_IDENTITY"`
	Hello        string      `bson:"hello" json:"hello,omitempty" env:"K_EMAIL_HELLO"`
	From         string      `bson:"from" json:"from,omitempty" env:"K_EMAIL_FROM"`
	SkipAuth     bool        `bson:"skip_auth" json:"skip_auth" env:"K_EMAIL_SKIP_AUTH"`
	RequireTLS   bool        `bson:"require_tls" json:"require_tls" env:"K_EMAIL_REQUIRE_TLS"`
	TLSConfig    *tls.Config `bson:"tls_config" json:"tls_config"`
}

// Desensitize 脱敏
func (c *Config) Desensitize() {
	c.AuthPassword = ""
	c.AuthSecret = ""
}

// Validate todo
func (c *Config) Validate() error {
	if c.Host == "" {
		return errors.New("host, 邮件客户端服务器地址未配置")
	}

	if c.AuthUserName == "" {
		return errors.New("username, 邮件发送用户未配置")
	}

	if !c.SkipAuth {
		if c.AuthUserName == "" || c.AuthPassword == "" {
			return errors.New("启用认证后, 需要配置用户名和密码(usernme, password)")
		}
	}

	if c.From == "" {
		c.From = fmt.Sprintf("%s<%s>", strings.Split(c.AuthUserName, "@")[0], c.AuthUserName)
	}

	return nil
}

// NewDeaultTestSendRequest todo
func NewDeaultTestSendRequest() *TestSendRequest {
	return &TestSendRequest{
		SendMailRequest: notify.NewSendMailRequest(),
		Config:          NewDefaultConfig(),
	}
}

// TestSendRequest todo
type TestSendRequest struct {
	*notify.SendMailRequest
	*Config
}

// Send todo
func (req *TestSendRequest) Send() error {
	sd, err := NewSender(req.Config)
	if err != nil {
		return err
	}

	return sd.Send(req.SendMailRequest)
}
