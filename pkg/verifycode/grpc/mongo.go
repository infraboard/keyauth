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
	"github.com/infraboard/keyauth/pkg/system"
	"github.com/infraboard/keyauth/pkg/token/issuer"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/verifycode"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	col    *mongo.Collection
	issuer issuer.Issuer
	system system.Service
	log    logger.Logger
	user   user.UserServiceServer

	verifycode.UnimplementedVerifyCodeServiceServer
}

func (s *service) Config() error {
	db := conf.C().Mongo.GetDB()
	col := db.Collection("verify_code")

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

	if pkg.System == nil {
		return fmt.Errorf("depence system config service is required")
	}
	s.system = pkg.System

	if pkg.User == nil {
		return fmt.Errorf("depence user service is required")
	}
	s.user = pkg.User

	is, err := issuer.NewTokenIssuer()
	if err != nil {
		return err
	}
	s.issuer = is
	s.log = zap.L().Named("Verify Code")
	return nil
}

// HttpEntry todo
func (s *service) HTTPEntry() *http.EntrySet {
	return verifycode.HttpEntry()
}

func init() {
	pkg.RegistryService("verify_code", Service)
}
