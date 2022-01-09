package impl

import (
	"context"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"

	"github.com/infraboard/keyauth/apps/namespace"
	"github.com/infraboard/keyauth/apps/policy"
	"github.com/infraboard/keyauth/apps/role"
	"github.com/infraboard/keyauth/apps/user"
	"github.com/infraboard/keyauth/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	log           logger.Logger
	col           *mongo.Collection
	enableCache   bool
	notifyCachPre string

	namespace namespace.ServiceServer
	user      user.ServiceServer
	role      role.ServiceServer

	policy.UnimplementedServiceServer
}

func (s *service) Config() error {
	s.namespace = app.GetGrpcApp(namespace.AppName).(namespace.ServiceServer)
	s.user = app.GetGrpcApp(user.AppName).(user.ServiceServer)
	s.role = app.GetGrpcApp(role.AppName).(role.ServiceServer)

	db := conf.C().Mongo.GetDB()
	col := db.Collection("policy")

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
	s.log = zap.L().Named("Policy")
	return nil
}

func (s *service) Name() string {
	return policy.AppName
}

func (s *service) Registry(server *grpc.Server) {
	policy.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
