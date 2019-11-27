package pkg

import (
	"errors"
	"fmt"

	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/keyauth/pkg/domain"
)

var (
	log = zap.L().Named("PKG")
)

var (
	// Domain 服务
	Domain domain.Service
)

// Service 注册上的服务必须实现的方法
type Service interface {
	Config() error
}

// Registry 服务实例注册
func Registry(name string, svr Service) error {
	switch value := svr.(type) {
	case domain.Service:
		if Domain != nil {
			return errors.New("service has registried")
		}
		Domain = value
		log.Info("domain service %s registry success", name)
	default:
		return fmt.Errorf("unknown service type: %v", value)
	}

	return nil
}
