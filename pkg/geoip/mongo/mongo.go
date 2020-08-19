package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/geoip"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	ip       *mongo.Collection
	location *mongo.Collection
	log      logger.Logger
}

func (s *service) Config() error {
	db := conf.C().Mongo.GetDB()
	s.ip = db.Collection("geoip_blocks")
	s.location = db.Collection("geoip_locations")
	s.log = zap.L().Named("GeoIP")

	// 添加ip表的索引
	ipIndexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "geoname_id", Value: bsonx.Int32(-1)}},
		},
	}

	_, err := s.ip.Indexes().CreateMany(context.Background(), ipIndexs)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	var _ geoip.Service = Service
	pkg.RegistryService("geoip", Service)
}
