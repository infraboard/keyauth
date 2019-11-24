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

// LoadConfigFromToml 从toml中添加配置文件, 并初始化全局对象
func LoadConfigFromToml(filePath string) error {
	cfg := newConfig()
	if _, err := toml.DecodeFile(filePath, cfg); err != nil {
		return err
	}

	// 加载全局配置单例
	global = cfg

	// 加载全局数据量单例
	mclient, err := cfg.Mongo.getClient()
	if err != nil {
		return err
	}
	mgoclient = mclient

	return nil
}
