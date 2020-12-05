package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg/system"
	"github.com/infraboard/keyauth/pkg/system/notify/mail"
	"github.com/infraboard/keyauth/pkg/system/notify/sms"
	"github.com/infraboard/mcube/exception"
)

func (s *service) UpdateEmail(*mail.Config) error {
	return nil
}

func (s *service) UpdateSMS(*sms.Config) error {
	return nil
}

func (s *service) GetConfig(version string) (*system.Config, error) {
	conf := system.NewDefaultConfig()
	if err := s.col.FindOne(context.TODO(), bson.M{"_id": version}).Decode(conf); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("version: %s system config %s not found", version)
		}

		return nil, exception.NewInternalServerError("find system config %s error, %s", version, err)
	}

	return conf, nil
}
