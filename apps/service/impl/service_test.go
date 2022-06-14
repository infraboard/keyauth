package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger/zap"

	// 注册所有服务
	_ "github.com/infraboard/keyauth/apps/all"
	"github.com/infraboard/keyauth/apps/service"
)

var (
	impl service.MetaServiceServer
)

func TestCreateService(t *testing.T) {
	req := service.NewCreateServiceRequest()
	req.Name = "cmdb"
	req.Description = "资源中心"
	req.Owner = "admin"
	app, err := impl.CreateService(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(app)
}

func TestQueryService(t *testing.T) {
	req := service.NewQueryServiceRequest()
	set, err := impl.QueryService(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func init() {
	zap.DevelopmentSetup()

	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}

	if err := app.InitAllApp(); err != nil {
		panic(err)
	}

	impl = app.GetGrpcApp(service.AppName).(service.MetaServiceServer)
}
