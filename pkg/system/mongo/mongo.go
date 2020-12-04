package mongo

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/system"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	log logger.Logger
	col *mongo.Collection
}

func (s *service) Config() error {
	s.log = zap.L().Named("System Config")
	db := conf.C().Mongo.GetDB()
	ac := db.Collection("system_config")

	s.col = ac
	return nil
}

func init() {
	var _ system.Service = Service
	pkg.RegistryService("system_config", Service)
}
