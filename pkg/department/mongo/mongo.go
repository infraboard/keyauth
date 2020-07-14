package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/counter"
	"github.com/infraboard/keyauth/pkg/department"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	col           *mongo.Collection
	enableCache   bool
	notifyCachPre string
	counter       counter.Service
}

func (s *service) Config() error {
	if pkg.Counter == nil {
		return fmt.Errorf("dependence counter service is nil")
	}
	s.counter = pkg.Counter

	db := conf.C().Mongo.GetDB()
	dc := db.Collection("department")

	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{Key: "domain_id", Value: bsonx.Int32(-1)},
				{Key: "name", Value: bsonx.Int32(-1)},
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err := dc.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	s.col = dc
	return nil
}

func init() {
	var _ department.Service = Service
	pkg.RegistryService("department", Service)
}
