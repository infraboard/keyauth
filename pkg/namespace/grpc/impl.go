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
	"github.com/infraboard/keyauth/pkg/department"
	"github.com/infraboard/keyauth/pkg/namespace"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/role"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	col           *mongo.Collection
	enableCache   bool
	notifyCachPre string
	depart        department.DepartmentServiceServer
	policy        policy.PolicyServiceServer
	role          role.RoleServiceServer
	log           logger.Logger

	namespace.UnimplementedNamespaceServiceServer
}

func (s *service) Config() error {
	if pkg.Department == nil {
		return fmt.Errorf("depence department service is nil")
	}
	s.depart = pkg.Department

	if pkg.Policy == nil {
		return fmt.Errorf("depence policy service is nil")
	}
	s.policy = pkg.Policy

	if pkg.Role == nil {
		return fmt.Errorf("depence role service is nil")
	}
	s.role = pkg.Role

	db := conf.C().Mongo.GetDB()
	ac := db.Collection("namespace")

	indexs := []mongo.IndexModel{
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

	_, err := ac.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	s.col = ac
	s.log = zap.L().Named("Namespace")
	return nil
}

// HttpEntry todo
func (s *service) HTTPEntry() *http.EntrySet {
	return namespace.HttpEntry()
}

func init() {
	pkg.RegistryService("namespace", Service)
}
