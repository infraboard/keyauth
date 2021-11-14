package impl

import (
	"context"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"

	"github.com/infraboard/keyauth/app/tag"
	"github.com/infraboard/keyauth/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	key         *mongo.Collection
	value       *mongo.Collection
	enableCache bool

	log logger.Logger
	tag.UnimplementedServiceServer
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

func (s *service) Name() string {
	return tag.AppName
}

func (s *service) Registry(server *grpc.Server) {
	tag.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
