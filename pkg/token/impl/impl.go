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

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg/application"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/namespace"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/session"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/token/issuer"
	"github.com/infraboard/keyauth/pkg/token/security"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/verifycode"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	token.UnimplementedTokenServiceServer
	col           *mongo.Collection
	log           logger.Logger
	enableCache   bool
	notifyCachPre string

	app      application.ApplicationServiceServer
	user     user.UserServiceServer
	domain   domain.DomainServiceServer
	policy   policy.PolicyServiceServer
	issuer   issuer.Issuer
	endpoint endpoint.EndpointServiceServer
	session  session.ServiceServer
	checker  security.Checker
	code     verifycode.VerifyCodeServiceServer
	ns       namespace.NamespaceServiceServer
}

func (s *service) Config() error {
	s.app = app.GetGrpcApp(application.AppName).(application.ApplicationServiceServer)
	s.user = app.GetGrpcApp(user.AppName).(user.UserServiceServer)
	s.domain = app.GetGrpcApp(domain.AppName).(domain.DomainServiceServer)
	s.policy = app.GetGrpcApp(policy.AppName).(policy.PolicyServiceServer)
	s.endpoint = app.GetGrpcApp(endpoint.AppName).(endpoint.EndpointServiceServer)
	s.session = app.GetGrpcApp(session.AppName).(session.ServiceServer)
	s.code = app.GetGrpcApp(verifycode.AppName).(verifycode.VerifyCodeServiceServer)
	s.ns = app.GetGrpcApp(namespace.AppName).(namespace.NamespaceServiceServer)

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
	token.RegisterTokenServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
