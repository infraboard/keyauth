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
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/role"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	col           *mongo.Collection
	perm          *mongo.Collection
	enableCache   bool
	notifyCachPre string

	policy policy.PolicyServiceServer
	log    logger.Logger
	role.UnimplementedRoleServiceServer
}

func (s *service) Config() error {
	if pkg.Policy == nil {
		return fmt.Errorf("dependence policy service is nil, please load first")
	}
	s.policy = pkg.Policy

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

// HttpEntry todo
func (s *service) HTTPEntry() *http.EntrySet {
	return role.HttpEntry()
}

func init() {
	pkg.RegistryService("role", Service)
}
