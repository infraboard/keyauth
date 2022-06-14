package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/keyauth/apps/service"
	"github.com/infraboard/keyauth/conf"
)

var (
	// Service 服务实例
	svr = &impl{}
)

type impl struct {
	col *mongo.Collection
	log logger.Logger
	service.UnimplementedMetaServiceServer
}

func (i *impl) Config() error {
	db := conf.C().Mongo.GetDB()
	i.col = db.Collection(i.Name())

	i.log = zap.L().Named(i.Name())
	return nil
}

func (i *impl) Name() string {
	return service.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	service.RegisterMetaServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
