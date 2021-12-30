package impl

import (
	"context"
	"github.com/infraboard/keyauth/app/wxwork"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/mcube/app"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col *mongo.Collection
}

func (s *service) Config() error {
	db := conf.C().Mongo.GetDB()
	ac := db.Collection("wxwork")

	indexs := []mongo.IndexModel{
		{
			Keys:    bsonx.Doc{{Key: "corp_id", Value: bsonx.Int32(-1)}},
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

func (s *service) Name() string {
	return wxwork.AppName
}

func init() {
	app.RegistryInternalApp(svr)
}
