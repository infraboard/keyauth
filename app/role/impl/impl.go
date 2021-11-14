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

	"github.com/infraboard/keyauth/app/policy"
	"github.com/infraboard/keyauth/app/role"
	"github.com/infraboard/keyauth/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col  *mongo.Collection
	perm *mongo.Collection

	policy policy.ServiceServer
	log    logger.Logger
	role.UnimplementedServiceServer
}

func (s *service) Config() error {
	s.policy = app.GetGrpcApp(policy.AppName).(policy.ServiceServer)

	db := conf.C().Mongo.GetDB()
	col := db.Collection("role")

	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{Key: "name", Value: bsonx.Int32(-1)},
				{Key: "domain", Value: bsonx.Int32(-1)},
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err := col.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	perm := db.Collection("permission")
	permIndexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err = perm.Indexes().CreateMany(context.Background(), permIndexs)
	if err != nil {
		return err
	}

	s.col = col
	s.perm = perm
	s.log = zap.L().Named("Role")
	return nil
}

func (s *service) Name() string {
	return role.AppName
}

func (s *service) Registry(server *grpc.Server) {
	role.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
