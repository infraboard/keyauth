package impl

import (
	"sync"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/keyauth/apps/ip2region"
	"github.com/infraboard/keyauth/apps/ip2region/reader"
	"github.com/infraboard/keyauth/apps/storage"
)

var (
	// Service 服务实例
	svr = &service{
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
	s.storage = app.GetInternalApp(storage.AppName).(storage.Service)

	s.log = zap.L().Named("IP2Region")
	return nil
}

func (s *service) Name() string {
	return ip2region.AppName
}

func init() {
	app.RegistryInternalApp(svr)
}
