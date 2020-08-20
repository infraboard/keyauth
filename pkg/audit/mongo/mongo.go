package mongo

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/audit"
	"github.com/infraboard/keyauth/pkg/ip2region"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	login         *mongo.Collection
	enableCache   bool
	notifyCachPre string
	ip            ip2region.Service
	log           logger.Logger
}

func (s *service) Config() error {
	if pkg.IP2Region == nil {
		return fmt.Errorf("depence service ip2region is nil")
	}
	s.ip = pkg.IP2Region

	db := conf.C().Mongo.GetDB()
	dc := db.Collection("login")

	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "account", Value: bsonx.Int32(-1)}},
		},
		{
			Keys: bsonx.Doc{{Key: "login_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err := dc.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	s.login = dc
	s.log = zap.L().Named("Audit")
	return nil
}

func init() {
	var _ audit.Service = Service
	pkg.RegistryService("audit", Service)
}
