package mongo

import (
	"errors"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/ip2region"
	"github.com/infraboard/keyauth/pkg/storage"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	storage storage.Service
	log     logger.Logger
}

func (s *service) Config() error {
	if pkg.Storage == nil {
		return errors.New("denpence Storage service is nil")
	}
	s.storage = pkg.Storage

	s.log = zap.L().Named("IP2Region")
	return nil
}

func init() {
	var _ ip2region.Service = Service
	pkg.RegistryService("geoip", Service)
}
