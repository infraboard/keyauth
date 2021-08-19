package grpc

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pb/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
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
	Service = &service{}
)

type service struct {
	scol          *mongo.Collection
	fcol          *mongo.Collection
	enableCache   bool
	notifyCachPre string

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

	if pkg.Token == nil {
		return fmt.Errorf("dependence token service is nil, please load first")
	}
	s.token = pkg.Token

	if pkg.Application == nil {
		return fmt.Errorf("dependence application service is nil, please load first")
	}
	s.app = pkg.Application

	if pkg.User == nil {
		return fmt.Errorf("dependence user service is nil, please load first")
	}
	s.user = pkg.User

	if pkg.Policy == nil {
		return fmt.Errorf("dependence policy service is nil, please load first")
	}
	s.policy = pkg.Policy

	if pkg.Role == nil {
		return fmt.Errorf("dependence role service is nil, please load first")
	}
	s.role = pkg.Role

	if pkg.Endpoint == nil {
		return fmt.Errorf("dependence endpoint service is nil, please load first")
	}
	s.endpoint = pkg.Endpoint
	return nil
}

// HttpEntry todo
func (s *service) HTTPEntry() *http.EntrySet {
	return micro.HttpEntry()
}

func init() {
	pkg.RegistryService("micro", Service)
}
