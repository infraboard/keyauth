package impl

import (
	"context"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"

	"github.com/infraboard/keyauth/apps/ip2region"
	"github.com/infraboard/keyauth/apps/session"
	"github.com/infraboard/keyauth/apps/token"
	"github.com/infraboard/keyauth/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col *mongo.Collection

	ip    ip2region.Service
	token token.ServiceServer
	log   logger.Logger
	session.UnimplementedServiceServer
}

func (s *service) Config() error {
	s.ip = app.GetInternalApp(ip2region.AppName).(ip2region.Service)
	s.token = app.GetGrpcApp(token.AppName).(token.ServiceServer)

	db := conf.C().Mongo.GetDB()
	dc := db.Collection("session")

	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "account", Value: bsonx.Int32(-1)}},
		},
		{
			Keys: bsonx.Doc{{Key: "login_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err := dc.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	s.col = dc
	s.log = zap.L().Named("Session")
	return nil
}

func (s *service) Name() string {
	return session.AppName
}

func (s *service) Registry(server *grpc.Server) {
	session.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
