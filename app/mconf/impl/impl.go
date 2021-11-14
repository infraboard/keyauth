package impl

import (
	"context"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"

	"github.com/infraboard/keyauth/app/mconf"
	"github.com/infraboard/keyauth/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	group *mongo.Collection
	item  *mongo.Collection
	log   logger.Logger

	mconf.UnimplementedConfigServiceServer
}

func (s *service) Config() error {
	if err := s.configService(); err != nil {
		return err
	}
	return nil
}

func (s *service) configService() error {
	s.log = zap.L().Named("Micro Config")

	db := conf.C().Mongo.GetDB()
	gc := db.Collection("config_group")
	gIndexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}
	_, err := gc.Indexes().CreateMany(context.Background(), gIndexs)
	if err != nil {
		return err
	}
	s.group = gc

	ic := db.Collection("config_item")
	iIndexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}
	_, err = ic.Indexes().CreateMany(context.Background(), iIndexs)
	if err != nil {
		return err
	}
	s.item = ic
	return nil
}

func (s *service) Name() string {
	return mconf.AppName
}

func (s *service) Registry(server *grpc.Server) {
	mconf.RegisterConfigServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
