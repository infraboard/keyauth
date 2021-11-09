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
	"github.com/infraboard/keyauth/pkg/counter"
	"github.com/infraboard/keyauth/pkg/department"
	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/keyauth/pkg/user"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	dc            *mongo.Collection
	ac            *mongo.Collection
	enableCache   bool
	notifyCachPre string
	counter       counter.Service
	user          user.UserServiceServer
	role          role.RoleServiceServer
	log           logger.Logger

	department.UnimplementedDepartmentServiceServer
}

func (s *service) Config() error {
	s.counter = app.GetInternalApp(counter.AppName).(counter.Service)
	s.user = app.GetGrpcApp(user.AppName).(user.UserServiceServer)
	s.role = app.GetGrpcApp(role.AppName).(role.RoleServiceServer)

	db := conf.C().Mongo.GetDB()

	dc := db.Collection("department")
	dcIndexs := []mongo.IndexModel{
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
	_, err := dc.Indexes().CreateMany(context.Background(), dcIndexs)
	if err != nil {
		return err
	}

	ac := db.Collection("join_apply")
	acIndexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}
	_, err = ac.Indexes().CreateMany(context.Background(), acIndexs)
	if err != nil {
		return err
	}

	s.dc = dc
	s.ac = ac
	s.log = zap.L().Named("Department")
	return nil
}

func (s *service) Name() string {
	return department.AppName
}

func (s *service) Registry(server *grpc.Server) {
	department.RegisterDepartmentServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
