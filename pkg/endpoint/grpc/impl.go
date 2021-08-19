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
	micro         micro.MicroServiceServer
	log           logger.Logger

	endpoint.UnimplementedEndpointServiceServer
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
	s.log = zap.L().Named("Endpoint")
	return nil
}

// HttpEntry todo
func (s *service) HTTPEntry() *http.EntrySet {
	return endpoint.HttpEntry()
}

func init() {
	pkg.RegistryService("endpoint", Service)
}
