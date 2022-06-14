package rpc

import (
	"net/url"

	"github.com/infraboard/keyauth/apps/instance"
	"github.com/infraboard/keyauth/client/rpc/auth"
)

// NewDefaultConfig todo
func NewDefaultConfig() *Config {
	return &Config{
		Address:  "localhost:18010",
		Resolver: NewDefaultResolver(),
	}
}

// Config 客户端配置
type Config struct {
	Address      string    `json:"adress" toml:"adress" yaml:"adress" env:"MCENTER_ADDRESS"`
	ClientID     string    `json:"client_id" toml:"client_id" yaml:"client_id" env:"MCENTER_CLINET_ID"`
	ClientSecret string    `json:"client_secret" toml:"client_secret" yaml:"client_secret" env:"MCENTER_CLIENT_SECRET"`
	Resolver     *Resolver `json:"resolver" toml:"resolver" yaml:"resolver"`
}

func (c *Config) Credentials() *auth.Authentication {
	return auth.NewAuthentication(c.ClientID, c.ClientSecret)
}

func NewDefaultResolver() *Resolver {
	return &Resolver{
		Region:      instance.DefaultRegion,
		Environment: instance.DefaultEnvironment,
		Group:       instance.DefaultGroup,
	}
}

type Resolver struct {
	// 实例所属地域, 默认default
	Region string `json:"region" toml:"region" yaml:"region" env:"MCENTER_REGION" validate:"required"`
	// 实例所属环境, 默认default
	Environment string `json:"environment" toml:"environment" yaml:"environment" env:"MCENTER_ENV" validate:"required"`
	// 实例所属分组,默认default
	Group string `json:"group" toml:"group" yaml:"group" env:"MCENTER_GROUP" validate:"required"`
	// 实例标签, 可以根据标签快速过滤实例, 格式k=v,k=v
	Tags string `json:"tags" toml:"tags" yaml:"tags" env:"MCENTER_TAGS"`
}

func (r *Resolver) ToQueryString() string {
	m := make(url.Values)
	m.Add("region", r.Region)
	m.Add("environment", r.Environment)
	m.Add("group", r.Group)
	m.Add("tags", r.Tags)
	return m.Encode()
}
