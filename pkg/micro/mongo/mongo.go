package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/micro"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	scol          *mongo.Collection
	fcol          *mongo.Collection
	enableCache   bool
	notifyCachPre string
}

func (s *service) Config() error {
	if err := s.configService(); err != nil {
		return err
	}
	if err := s.configFeature(); err != nil {
		return err
	}
	return nil
}

func (s *service) configService() error {
	db := conf.C().Mongo.GetDB()
	sc := db.Collection("service")
	sIndexs := []mongo.IndexModel{
		mongo.IndexModel{
			Keys:    bsonx.Doc{{Key: "client_id", Value: bsonx.Int32(-1)}},
			Options: options.Index().SetUnique(true),
		},
		mongo.IndexModel{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err := sc.Indexes().CreateMany(context.Background(), sIndexs)
	if err != nil {
		return err
	}
	s.scol = sc
	return nil
}

func (s *service) configFeature() error {
	db := conf.C().Mongo.GetDB()
	fc := db.Collection("feature")
	fIndexs := []mongo.IndexModel{
		mongo.IndexModel{
			Keys: bsonx.Doc{
				{Key: "service_name", Value: bsonx.Int32(-1)},
				{Key: "path", Value: bsonx.Int32(-1)},
			},
			Options: options.Index().SetUnique(true),
		},
		mongo.IndexModel{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}
	_, err := fc.Indexes().CreateMany(context.Background(), fIndexs)
	if err != nil {
		return err
	}
	s.fcol = fc
	return nil
}

func init() {
	var _ micro.Service = Service
	pkg.RegistryService("service", Service)
}
