package mongo

import (
	"context"

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
	uc            *mongo.Collection
	enableCache   bool
	notifyCachPre string
}

func (s *service) Config() error {
	db := conf.C().Mongo.GetDB()
	uc := db.Collection("user")

	indexs := []mongo.IndexModel{
		mongo.IndexModel{
			Keys:    bsonx.Doc{{Key: "account", Value: bsonx.Int32(-1)}},
			Options: options.Index().SetUnique(true),
		},
		mongo.IndexModel{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err := uc.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	s.uc = uc
	return nil
}

func init() {
	pkg.RegistryService("user", Service)
}