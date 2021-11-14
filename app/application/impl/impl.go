package impl

import (
	"context"

	"github.com/infraboard/mcube/app"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"

	"github.com/infraboard/keyauth/app/application"
	"github.com/infraboard/keyauth/conf"
)

var (
	// Service 服务实例
	svr = &userimpl{service: &service{}}
)

type service struct {
	col *mongo.Collection
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

func (s *service) Name() string {
	return application.AppName
}

func (s *service) Registry(server *grpc.Server) {
	application.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
