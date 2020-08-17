package mongo

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/counter"
	"github.com/infraboard/keyauth/pkg/geoip"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	bucket *gridfs.Bucket
}

func (s *service) Config() error {
	db := conf.C().Mongo.GetDB()

	opts := options.GridFSBucket()
	opts.SetName("geoip_db")

	bucket, err := gridfs.NewBucket(db, opts)
	if err != nil {
		return fmt.Errorf("new bucket error, %s", err)
	}
	s.bucket = bucket
	return nil
}

func (s *service) GetNextSequenceValue(sequenceName string) (*counter.Count, error) {
	return nil, nil
}

func init() {
	var _ geoip.Service = Service
	pkg.RegistryService("geoip", Service)
}
