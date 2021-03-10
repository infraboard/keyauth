package mongo

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pb/http"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
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

// HttpEntry todo
func (s *service) HTTPEntry() *http.EntrySet {
	return http.NewEntrySet()
}

func init() {
	pkg.RegistryService("storage", Service)
}
