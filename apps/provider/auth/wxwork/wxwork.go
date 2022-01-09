package wxwork

import (
	"errors"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// NewProvider todo
func NewProvider(conf *Config) *Provider {
	return &Provider{
		conf: conf,
		log:  zap.L().Named("WechatWork"),
	}
}

// Provider todo
type Provider struct {
	conf *Config
	log  logger.Logger
}

func (p *Provider) Check() error {
	wx := Wechat{
		AppID:     p.conf.CorpID,     // 企业微信app ID
		AppSecret: p.conf.CorpSecret, // 企业微信app secret
		AgentID:   p.conf.AgentID,    // 企业微信 应用ID
	}
	token := wx.GetAccessToken()
	if token == "" {
		return errors.New("Authentication failed ")
	}
	return nil
}
