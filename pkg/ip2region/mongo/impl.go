package mongo

import (
	"errors"
	"sync"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pb/http"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/ip2region/reader"
	"github.com/infraboard/keyauth/pkg/storage"
)

var (
	// Service 服务实例
	Service = &service{
		bucketName: "ip2region",
		dbFileName: "ip2region.db",
	}
)

type service struct {
	storage    storage.Service
	log        logger.Logger
	bucketName string
	dbFileName string
	dbReader   *reader.IPReader
	sync.Mutex
}

func (s *service) Config() error {
	if pkg.Storage == nil {
		return errors.New("denpence Storage service is nil")
	}
	s.storage = pkg.Storage

	s.log = zap.L().Named("IP2Region")
	return nil
}

// HttpEntry todo
func (s *service) HTTPEntry() *http.EntrySet {
	return http.NewEntrySet()
}

func init() {
	pkg.RegistryService("ip2region", Service)
}
