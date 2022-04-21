package impl

import (
	"github.com/infraboard/keyauth/apps/otp"
	"github.com/infraboard/keyauth/apps/user"
	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	svr = &service{}
)

type service struct {
	col           *mongo.Collection
	log           logger.Logger
	enableCache   bool
	notifyCachPre string

	user user.ServiceServer
	otp.UnimplementedServiceServer
}

func (s *service) Config() error {
	db := conf.C().Mongo.GetDB()
	col := db.Collection("otp")

	s.col = col
	s.log = zap.L().Named("OTP")
	s.user = app.GetGrpcApp(user.AppName).(user.ServiceServer)
	return nil
}
func (s *service) Name() string {
	return otp.AppName
}
func (s *service) Registry(server *grpc.Server) {
	otp.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
