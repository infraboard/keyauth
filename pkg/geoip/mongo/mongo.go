package mongo

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/geoip"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	// Service 服务实例
	Service = &service{
		dbFileName: "GeoLite2-City.mmdb",
	}
)

type service struct {
	bucket     *gridfs.Bucket
	dbFileName string
	log        logger.Logger
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
	s.log = zap.L().Named("GeoIP")
	return nil
}

func init() {
	var _ geoip.Service = Service
	pkg.RegistryService("geoip", Service)
}
