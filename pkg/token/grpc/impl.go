package grpc

import (
	"context"
	"errors"
	"fmt"

	"github.com/infraboard/mcube/cache"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pb/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
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
	Service = &service{}
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
	session  session.UserServiceServer
	checker  security.Checker
	code     verifycode.VerifyCodeServiceServer
	ns       namespace.NamespaceServiceServer
}

func (s *service) Config() error {
	if pkg.Application == nil {
		return errors.New("denpence application service is nil")
	}
	s.app = pkg.Application

	if pkg.User == nil {
		return errors.New("denpence user service is nil")
	}
	s.user = pkg.User

	if pkg.Domain == nil {
		return errors.New("denpence domain service is nil")
	}
	s.domain = pkg.Domain

	if pkg.Policy == nil {
		return errors.New("denpence policy service is nil")
	}
	s.policy = pkg.Policy

	if pkg.Endpoint == nil {
		return errors.New("denpence endpoint service is nil")
	}
	s.endpoint = pkg.Endpoint

	if pkg.SessionUser == nil {
		return errors.New("denpence session service is nil")
	}
	s.session = pkg.SessionUser

	if pkg.VerifyCode == nil {
		return errors.New("denpence verify code service is nil")
	}
	s.code = pkg.VerifyCode

	if pkg.Namespace == nil {
		return errors.New("denpence namespace service is nil")
	}
	s.ns = pkg.Namespace

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

// HttpEntry todo
func (s *service) HTTPEntry() *http.EntrySet {
	return token.HttpEntry()
}

func init() {
	pkg.RegistryService("token", Service)
}
