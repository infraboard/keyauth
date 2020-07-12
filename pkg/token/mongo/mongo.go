package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/application"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/token/issuer"
	"github.com/infraboard/keyauth/pkg/user"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	col           *mongo.Collection
	enableCache   bool
	notifyCachPre string

	app      application.Service
	user     user.Service
	domain   domain.Service
	issuer   issuer.Issuer
	endpoint endpoint.Service
}

func (s *service) Config() error {
	if pkg.Application == nil {
		return errors.New("denpence application service is nil")
	}
	s.app = pkg.Application

	if pkg.User == nil {
		return errors.New("denpence user service is nil")
	}
	s.user = pkg.User

	if pkg.Domain == nil {
		return errors.New("denpence domain service is nil")
	}
	s.domain = pkg.Domain

	if pkg.Endpoint == nil {
		return errors.New("denpence endpoint service is nil")
	}
	s.endpoint = pkg.Endpoint

	issuer, err := issuer.NewTokenIssuer()
	if err != nil {
		return err
	}
	s.issuer = issuer

	db := conf.C().Mongo.GetDB()
	col := db.Collection("token")

	indexs := []mongo.IndexModel{
		{
			Keys:    bsonx.Doc{{Key: "refresh_token", Value: bsonx.Int32(-1)}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err = col.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	s.col = col
	return nil
}

func init() {
	var _ token.Service = Service
	pkg.RegistryService("token", Service)
}
