package grpc

import (
	"context"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"

	"github.com/infraboard/keyauth/apps/department"
	"github.com/infraboard/keyauth/apps/namespace"
	"github.com/infraboard/keyauth/apps/policy"
	"github.com/infraboard/keyauth/apps/role"
	"github.com/infraboard/keyauth/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col    *mongo.Collection
	depart department.ServiceServer
	policy policy.ServiceServer
	role   role.ServiceServer
	log    logger.Logger

	namespace.UnimplementedServiceServer
}

func (s *service) Config() error {
	s.depart = app.GetGrpcApp(department.AppName).(department.ServiceServer)
	s.policy = app.GetGrpcApp(policy.AppName).(policy.ServiceServer)
	s.role = app.GetGrpcApp(role.AppName).(role.ServiceServer)

	db := conf.C().Mongo.GetDB()
	ac := db.Collection("namespace")

	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{Key: "domain", Value: bsonx.Int32(-1)},
				{Key: "name", Value: bsonx.Int32(-1)},
			},
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
	s.log = zap.L().Named("Namespace")
	return nil
}

func (s *service) Name() string {
	return namespace.AppName
}

func (s *service) Registry(server *grpc.Server) {
	namespace.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
