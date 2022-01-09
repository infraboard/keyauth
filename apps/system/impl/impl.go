package impl

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/apps/system"
	"github.com/infraboard/keyauth/conf"
)

var (
	// Service 服务实例
	svr = &service{}
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

func (s *service) Name() string {
	return system.AppName
}

func init() {
	app.RegistryInternalApp(svr)
}
