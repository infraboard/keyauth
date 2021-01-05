package pkg

import (
	"google.golang.org/grpc"
)

var (
	v1GRPCAPIs = make(map[string]GRPCAPI)
)

// LoadedGRPC 查询加载成功的GRPC API
func LoadedGRPC() []string {
	var apis []string
	for k := range v1GRPCAPIs {
		apis = append(apis, k)
	}
	return apis
}

// GRPCAPI todo
type GRPCAPI interface {
	Registry(*grpc.Server)
	Config() error
}

// RegistryGRPCV1 注册GRPC服务
func RegistryGRPCV1(name string, api GRPCAPI) {
	if _, ok := v1GRPCAPIs[name]; ok {
		panic("grpc api " + name + " has registry")
	}
	v1GRPCAPIs[name] = api
}

// InitV1GRPCAPI 初始化API服务
func InitV1GRPCAPI(server *grpc.Server) error {
	for _, api := range v1GRPCAPIs {
		if err := api.Config(); err != nil {
			return err
		}

		api.Registry(server)
	}

	return nil
}
