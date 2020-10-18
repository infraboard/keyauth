package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/micro"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	col           *mongo.Collection
	enableCache   bool
	notifyCachPre string

	micro micro.Service
}

func (s *service) Config() error {
	db := conf.C().Mongo.GetDB()
	col := db.Collection("endpoint")

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

	if pkg.Micro == nil {
		return fmt.Errorf("dependence micro service is nil, please load first")
	}
	s.micro = pkg.Micro
	return nil
}

func init() {
	var _ endpoint.Service = Service
	pkg.RegistryService("endpoint", Service)
}
