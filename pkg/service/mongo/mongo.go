package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/service"
)

var (
	// Service 服务实例
	Service = &microService{}
)

type microService struct {
	col           *mongo.Collection
	enableCache   bool
	notifyCachPre string
}

func (s *microService) Config() error {
	db := conf.C().Mongo.GetDB()
	ac := db.Collection("service")

	indexs := []mongo.IndexModel{
		mongo.IndexModel{
			Keys:    bsonx.Doc{{Key: "client_id", Value: bsonx.Int32(-1)}},
			Options: options.Index().SetUnique(true),
		},
		mongo.IndexModel{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err := ac.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	s.col = ac
	return nil
}

func init() {
	var _ service.Service = Service
	pkg.RegistryService("service", Service)
}
