package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	dc            *mongo.Collection
	enableCache   bool
	notifyCachPre string
}

func (s *service) Config() error {
	db := conf.C().Mongo.GetDB()
	s.dc = db.Collection("domain")
	return nil
}

func init() {
	pkg.RegistryService("domain", Service)
}
