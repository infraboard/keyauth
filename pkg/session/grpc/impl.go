package grpc

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pb/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/ip2region"
	"github.com/infraboard/keyauth/pkg/session"
	"github.com/infraboard/keyauth/pkg/token"
)

var (
	// UserService 服务实例
	UserService = &userimpl{service: &service{}}
	// AdminService todo
	AdminService = &adminimpl{service: &service{}}
)

type service struct {
	col           *mongo.Collection
	enableCache   bool
	notifyCachPre string

	ip    ip2region.Service
	token token.TokenServiceServer
	log   logger.Logger
}

func (s *service) Config() error {
	if pkg.IP2Region == nil {
		return fmt.Errorf("depence service ip2region is nil")
	}
	s.ip = pkg.IP2Region

	if pkg.Token == nil {
		return fmt.Errorf("depence service token is nil")
	}
	s.token = pkg.Token

	db := conf.C().Mongo.GetDB()
	dc := db.Collection("session")

	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "account", Value: bsonx.Int32(-1)}},
		},
		{
			Keys: bsonx.Doc{{Key: "login_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err := dc.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	s.col = dc
	s.log = zap.L().Named("Session")
	return nil
}

// HttpEntry todo
func (s *service) HTTPEntry() *http.EntrySet {
	return session.HttpEntry()
}

func init() {
	pkg.RegistryService("application_admin", AdminService)
	pkg.RegistryService("session_user", UserService)
}
