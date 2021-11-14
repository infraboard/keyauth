package impl

import (
	"context"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"

	"github.com/infraboard/keyauth/app/application"
	"github.com/infraboard/keyauth/app/endpoint"
	"github.com/infraboard/keyauth/app/micro"
	"github.com/infraboard/keyauth/app/policy"
	"github.com/infraboard/keyauth/app/role"
	"github.com/infraboard/keyauth/app/token"
	"github.com/infraboard/keyauth/app/user"
	"github.com/infraboard/keyauth/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	scol *mongo.Collection

	token    token.ServiceServer
	user     user.ServiceServer
	app      application.ServiceServer
	policy   policy.ServiceServer
	role     role.ServiceServer
	endpoint endpoint.ServiceServer
	log      logger.Logger

	micro.UnimplementedServiceServer
}

func (s *service) Config() error {
	if err := s.configService(); err != nil {
		return err
	}
	return nil
}

func (s *service) configService() error {
	db := conf.C().Mongo.GetDB()
	sc := db.Collection("micro")
	sIndexs := []mongo.IndexModel{
		{
			Keys:    bsonx.Doc{{Key: "name", Value: bsonx.Int32(-1)}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bsonx.Doc{{Key: "client_id", Value: bsonx.Int32(-1)}},
		},
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err := sc.Indexes().CreateMany(context.Background(), sIndexs)
	if err != nil {
		return err
	}
	s.scol = sc
	s.log = zap.L().Named("Micro")

	s.token = app.GetGrpcApp(token.AppName).(token.ServiceServer)
	s.app = app.GetGrpcApp(application.AppName).(application.ServiceServer)
	s.user = app.GetGrpcApp(user.AppName).(user.ServiceServer)
	s.policy = app.GetGrpcApp(policy.AppName).(policy.ServiceServer)
	s.role = app.GetGrpcApp(role.AppName).(role.ServiceServer)
	s.endpoint = app.GetGrpcApp(endpoint.AppName).(endpoint.ServiceServer)
	return nil
}

func (s *service) Name() string {
	return micro.AppName
}

func (s *service) Registry(server *grpc.Server) {
	micro.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
