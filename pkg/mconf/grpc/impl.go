package grpc

import (
	"context"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pb/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/mconf"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	group *mongo.Collection
	item  *mongo.Collection
	log   logger.Logger

	mconf.UnimplementedMicroConfigServiceServer
}

func (s *service) Config() error {
	if err := s.configService(); err != nil {
		return err
	}
	return nil
}

func (s *service) configService() error {
	s.log = zap.L().Named("Micro Config")

	db := conf.C().Mongo.GetDB()
	gc := db.Collection("config_group")
	gIndexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}
	_, err := gc.Indexes().CreateMany(context.Background(), gIndexs)
	if err != nil {
		return err
	}
	s.group = gc

	ic := db.Collection("config_item")
	iIndexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}
	_, err = ic.Indexes().CreateMany(context.Background(), iIndexs)
	if err != nil {
		return err
	}
	s.item = ic
	return nil
}

// HttpEntry todo
func (s *service) HTTPEntry() *http.EntrySet {
	return mconf.HttpEntry()
}

func init() {
	pkg.RegistryService("mconf", Service)
}
