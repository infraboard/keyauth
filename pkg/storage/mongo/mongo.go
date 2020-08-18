package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/storage"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	log logger.Logger
	db  *mongo.Database
}

func (s *service) Config() error {
	s.db = conf.C().Mongo.GetDB()
	s.log = zap.L().Named("Storage")
	return nil
}

func init() {
	var _ storage.Service = Service
	pkg.RegistryService("storage", Service)
}
