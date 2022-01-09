package wxwork

import "fmt"

// NewDefaultConfig 初始化默认配置
func NewDefaultConfig() *Config {
	return &Config{
		CorpID:     "",
		AgentID:    "",
		CorpSecret: "",
		State:      "",
	}
}

// Config 企业微信相关认证配置
type Config struct {
	CorpID     string `bson:"corp_id" json:"corp_id"`
	AgentID    string `bson:"agent_id" json:"agent_id"`
	CorpSecret string `bson:"corp_secret" json:"corp_secret"`
	State      string `bson:"state" json:"state"`
}

func (c *Config) Validate() error {
	if c.CorpID == "" || c.CorpSecret == "" {
		return fmt.Errorf("corpID and CorpSecret required")
	}
	return nil
}
