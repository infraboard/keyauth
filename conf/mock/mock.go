package mock

import (
	"github.com/infraboard/keyauth/conf"
)

// Load 加载mock配置文件
func Load() {
	if err := conf.LoadConfigFromToml("E:\\Projects\\Golang\\keyauth\\etc\\keyauth.toml"); err != nil {
		panic(err)
	}
}
