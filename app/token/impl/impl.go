package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/cache"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"

	"github.com/infraboard/keyauth/app/application"
	"github.com/infraboard/keyauth/app/domain"
	"github.com/infraboard/keyauth/app/endpoint"
	"github.com/infraboard/keyauth/app/namespace"
	"github.com/infraboard/keyauth/app/policy"
	"github.com/infraboard/keyauth/app/session"
	"github.com/infraboard/keyauth/app/token"
	"github.com/infraboard/keyauth/app/token/issuer"
	"github.com/infraboard/keyauth/app/token/security"
	"github.com/infraboard/keyauth/app/user"
	"github.com/infraboard/keyauth/app/verifycode"
	"github.com/infraboard/keyauth/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	token.UnimplementedServiceServer
	col           *mongo.Collection
	log           logger.Logger
	enableCache   bool
	notifyCachPre string

	app      application.ServiceServer
	user     user.ServiceServer
	domain   domain.ServiceServer
	policy   policy.ServiceServer
	issuer   issuer.Issuer
	endpoint endpoint.ServiceServer
	session  session.ServiceServer
	checker  security.Checker
	code     verifycode.ServiceServer
	ns       namespace.ServiceServer
}

func (s *service) Config() error {
	s.app = app.GetGrpcApp(application.AppName).(application.ServiceServer)
	s.user = app.GetGrpcApp(user.AppName).(user.ServiceServer)
	s.domain = app.GetGrpcApp(domain.AppName).(domain.ServiceServer)
	s.policy = app.GetGrpcApp(policy.AppName).(policy.ServiceServer)
	s.endpoint = app.GetGrpcApp(endpoint.AppName).(endpoint.ServiceServer)
	s.session = app.GetGrpcApp(session.AppName).(session.ServiceServer)
	s.code = app.GetGrpcApp(verifycode.AppName).(verifycode.ServiceServer)
	s.ns = app.GetGrpcApp(namespace.AppName).(namespace.ServiceServer)

	issuer, err := issuer.NewTokenIssuer()
	if err != nil {
		return err
	}
	s.issuer = issuer

	c := cache.C()
	if c == nil {
		return fmt.Errorf("denpence cache service is nil")
	}
	s.checker, err = security.NewChecker()
	if err != nil {
		return fmt.Errorf("new checker error, %s", err)
	}

	db := conf.C().Mongo.GetDB()
	col := db.Collection("token")

	indexs := []mongo.IndexModel{
		{
			Keys:    bsonx.Doc{{Key: "refresh_token", Value: bsonx.Int32(-1)}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err = col.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	s.col = col
	s.log = zap.L().Named("token")
	return nil
}

func (s *service) Name() string {
	return token.AppName
}

func (s *service) Registry(server *grpc.Server) {
	token.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
