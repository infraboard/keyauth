package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	dc            *mongo.Collection
	enableCache   bool
	notifyCachPre string
}

func (s *service) Config() error {
	db := conf.C().Mongo.GetDB()
	dc := db.Collection("domain")

	nameIndex := mongo.IndexModel{
		Keys:    bsonx.Doc{{Key: "name", Value: bsonx.Int32(-1)}},
		Options: options.Index().SetUnique(true),
	}
	_, err := dc.Indexes().CreateOne(context.Background(), nameIndex)
	if err != nil {
		return err
	}

	s.dc = dc

	return nil
}

func init() {
	pkg.RegistryService("domain", Service)
}
