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
	"github.com/infraboard/keyauth/pkg/department"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/user"
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
	policy        policy.PolicyServiceServer
	depart        department.DepartmentServiceServer
	domain        domain.DomainServiceServer

	user.UnimplementedUserServiceServer
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

	if pkg.Domain == nil {
		return fmt.Errorf("dependence domain service is nil")
	}
	s.domain = pkg.Domain

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

// HttpEntry todo
func (s *service) HTTPEntry() *http.EntrySet {
	return user.HttpEntry()
}

func init() {
	pkg.RegistryService("user", Service)
}
