package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/counter"
	"github.com/infraboard/mcube/pb/http"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	col *mongo.Collection
}

func (s *service) Config() error {
	db := conf.C().Mongo.GetDB()
	dc := db.Collection("counter")

	s.col = dc
	return nil
}

func (s *service) GetNextSequenceValue(sequenceName string) (*counter.Count, error) {
	result := s.col.FindOneAndUpdate(
		context.TODO(),
		bson.M{"_id": sequenceName},
		bson.M{"$inc": bson.M{"value": 1}},
		options.FindOneAndUpdate().SetUpsert(true),
	)

	count := counter.NewCount()
	err := result.Decode(count)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, fmt.Errorf("counter decode error, %s", err)
	}

	count.Value++
	return count, nil
}

// HttpEntry todo
func (s *service) HTTPEntry() *http.EntrySet {
	return http.NewEntrySet()
}

func init() {
	pkg.RegistryService("counter", Service)
}
