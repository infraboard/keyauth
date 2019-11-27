package pkg

import (
	"github.com/infraboard/mcube/http/router"
)

var (
	httpAPIs = make(map[string]HTTPAPI)
)

// HTTPAPI restful 服务
type HTTPAPI interface {
	Registry(router router.SubRouter)
	Config() error
}

// RegistryHTTP 注册HTTP服务
func RegistryHTTP(name string, api HTTPAPI) {
	if _, ok := httpAPIs[name]; ok {
		panic("http api " + name + " has registry")
	}
	httpAPIs[name] = api
}

// InitHTTPAPI 初始化API服务
func InitHTTPAPI(root router.Router) error {
	for name, api := range httpAPIs {
		if err := api.Config(); err != nil {
			return err
		}

		api.Registry(root.SubRouter("/" + name))
	}

	return nil
}
