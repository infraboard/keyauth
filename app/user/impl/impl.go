package impl

import (
	"context"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"

	"github.com/infraboard/keyauth/app/department"
	"github.com/infraboard/keyauth/app/domain"
	"github.com/infraboard/keyauth/app/policy"
	"github.com/infraboard/keyauth/app/user"
	"github.com/infraboard/keyauth/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	log    logger.Logger
	col    *mongo.Collection
	policy policy.PolicyServiceServer
	depart department.DepartmentServiceServer
	domain domain.DomainServiceServer

	user.UnimplementedUserServiceServer
}

func (s *service) Config() error {
	s.policy = app.GetGrpcApp(policy.AppName).(policy.PolicyServiceServer)
	s.depart = app.GetGrpcApp(department.AppName).(department.DepartmentServiceServer)
	s.domain = app.GetGrpcApp(domain.AppName).(domain.DomainServiceServer)

	db := conf.C().Mongo.GetDB()
	uc := db.Collection("user")

	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
		{
			Keys: bsonx.Doc{{Key: "department_id", Value: bsonx.Int32(-1)}},
		},
	}

	_, err := uc.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	s.col = uc
	s.log = zap.L().Named("User")
	return nil
}

func (s *service) Name() string {
	return user.AppName
}

func (s *service) Registry(server *grpc.Server) {
	user.RegisterUserServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
