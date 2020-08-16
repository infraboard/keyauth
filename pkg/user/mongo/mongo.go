package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/department"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	log           logger.Logger
	col           *mongo.Collection
	enableCache   bool
	notifyCachPre string
	policy        policy.Service
	depart        department.Service
}

func (s *service) Config() error {
	if pkg.Namespace == nil {
		return fmt.Errorf("dependence namespace service is nil")
	}
	s.policy = pkg.Policy

	if pkg.Department == nil {
		return fmt.Errorf("dependence department service is nil")
	}
	s.depart = pkg.Department

	db := conf.C().Mongo.GetDB()
	uc := db.Collection("user")

	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
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

func init() {
	var _ user.Service = Service
	pkg.RegistryService("user", Service)
}
