package conf

import (
	"github.com/BurntSushi/toml"
)

var (
	global *Config
)

// C 全局配置对象
func C() *Config {
	if global == nil {
		panic("Load Config first")
	}

	return global
}

// LoadConfigFromToml 从toml中添加配置文件
func LoadConfigFromToml(filePath string) error {
	cfg := newConfig()
	if _, err := toml.DecodeFile(filePath, cfg); err != nil {
		return err
	}

	global = cfg
	return nil
}
