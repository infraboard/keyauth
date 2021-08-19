package mongo

import (
	"context"

	"github.com/infraboard/mcube/pb/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/application"
)

var (
	// Service 服务实例
	Service = &userimpl{service: &service{}}
)

type service struct {
	col           *mongo.Collection
	enableCache   bool
	notifyCachPre string
}

func (s *service) Config() error {
	db := conf.C().Mongo.GetDB()
	ac := db.Collection("application")

	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{Key: "user_id", Value: bsonx.Int32(-1)},
				{Key: "name", Value: bsonx.Int32(-1)}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bsonx.Doc{{Key: "client_id", Value: bsonx.Int32(-1)}},
			Options: options.Index().SetUnique(true),
		},
		{
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

// HttpEntry todo
func (s *service) HTTPEntry() *http.EntrySet {
	return application.HttpEntry()
}

func init() {
	pkg.RegistryService("application", Service)
}
