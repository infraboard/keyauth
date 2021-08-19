package grpc

import (
	"context"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pb/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/tag"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	key         *mongo.Collection
	value       *mongo.Collection
	enableCache bool

	log logger.Logger
	tag.UnimplementedTagServiceServer
}

func (s *service) Config() error {
	db := conf.C().Mongo.GetDB()
	key := db.Collection("tag_key")

	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err := key.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	value := db.Collection("tag_value")
	vIndexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err = value.Indexes().CreateMany(context.Background(), vIndexs)
	if err != nil {
		return err
	}

	s.key = key
	s.value = value
	s.log = zap.L().Named("Tag")
	return nil
}

// HttpEntry todo
func (s *service) HTTPEntry() *http.EntrySet {
	return tag.HttpEntry()
}

func init() {
	pkg.RegistryService("tag", Service)
}
