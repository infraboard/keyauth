package mongo

import (
	"context"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pb/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
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
			Keys: bsonx.Doc{
				{Key: "start", Value: bsonx.Int32(-1)},
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bsonx.Doc{
				{Key: "end", Value: bsonx.Int32(-1)},
			},
			Options: options.Index().SetUnique(true),
		},
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

// HttpEntry todo
func (s *service) HTTPEntry() *http.EntrySet {
	return http.NewEntrySet()
}

func init() {
	pkg.RegistryService("geoip", Service)
}
