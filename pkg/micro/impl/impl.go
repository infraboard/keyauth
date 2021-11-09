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

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg/application"
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	scol *mongo.Collection

	token    token.TokenServiceServer
	user     user.UserServiceServer
	app      application.ApplicationServiceServer
	policy   policy.PolicyServiceServer
	role     role.RoleServiceServer
	endpoint endpoint.EndpointServiceServer
	log      logger.Logger

	micro.UnimplementedMicroServiceServer
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

	s.token = app.GetGrpcApp(token.AppName).(token.TokenServiceServer)
	s.app = app.GetGrpcApp(application.AppName).(application.ApplicationServiceServer)
	s.user = app.GetGrpcApp(user.AppName).(user.UserServiceServer)
	s.policy = app.GetGrpcApp(policy.AppName).(policy.PolicyServiceServer)
	s.role = app.GetGrpcApp(role.AppName).(role.RoleServiceServer)
	s.endpoint = app.GetGrpcApp(endpoint.AppName).(endpoint.EndpointServiceServer)
	return nil
}

func (s *service) Name() string {
	return micro.AppName
}

func (s *service) Registry(server *grpc.Server) {
	micro.RegisterMicroServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
