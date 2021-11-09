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

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg/department"
	"github.com/infraboard/keyauth/pkg/namespace"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/role"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col    *mongo.Collection
	depart department.DepartmentServiceServer
	policy policy.PolicyServiceServer
	role   role.RoleServiceServer
	log    logger.Logger

	namespace.UnimplementedNamespaceServiceServer
}

func (s *service) Config() error {
	s.depart = app.GetGrpcApp(department.AppName).(department.DepartmentServiceServer)
	s.policy = app.GetGrpcApp(policy.AppName).(policy.PolicyServiceServer)
	s.role = app.GetGrpcApp(role.AppName).(role.RoleServiceServer)

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
	namespace.RegisterNamespaceServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
