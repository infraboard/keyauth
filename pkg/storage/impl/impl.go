package impl

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg/storage"
)

var (
	// Service 服务实例
	svr = &service{}
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

func (s *service) Name() string {
	return storage.AppName
}

func init() {
	app.RegistryInternalApp(svr)
}
