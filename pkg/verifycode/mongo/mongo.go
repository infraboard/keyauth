package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/verifycode"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	col *mongo.Collection
}

func (s *service) Config() error {
	db := conf.C().Mongo.GetDB()
	col := db.Collection("verify_code")

	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err := col.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	s.col = col
	return nil
}

func init() {
	var _ verifycode.Service = Service
	pkg.RegistryService("verify_code", Service)
}
