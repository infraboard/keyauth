package impl

import (
	"context"

	micro "github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/client/rpc"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"

	"github.com/infraboard/keyauth/apps/endpoint"
	"github.com/infraboard/keyauth/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col   *mongo.Collection
	log   logger.Logger
	micro micro.MetaServiceClient

	endpoint.UnimplementedServiceServer
}

func (s *service) Config() error {
	db := conf.C().Mongo.GetDB()
	col := db.Collection("endpoint")

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

	s.micro = rpc.C().Service()
	s.log = zap.L().Named("Endpoint")
	return nil
}

func (s *service) Name() string {
	return endpoint.AppName
}

func (s *service) Registry(server *grpc.Server) {
	endpoint.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
