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

func (s *service) UpdateEmail(mailConf *mail.Config) error {
	if err := mailConf.Validate(); err != nil {
		return exception.NewBadRequest("validate mail config error, %s", err)
	}

	_, err := s.GetConfig()
	if exception.IsNotFoundError(err) {
		conf := system.NewDefaultConfig()
		*conf.Email = *mailConf
		if err := s.save(conf); err != nil {
			return err
		}
	}

	return s.updateEmail(mailConf)
}

func (s *service) UpdateSMS(smsConf *sms.Config) error {
	if err := smsConf.Validate(); err != nil {
		return exception.NewBadRequest("validate mail config error, %s", err)
	}

	_, err := s.GetConfig()
	if exception.IsNotFoundError(err) {
		conf := system.NewDefaultConfig()
		*conf.SMS = *smsConf
		if err := s.save(conf); err != nil {
			return err
		}
	}
	return s.updateSMS(smsConf)
}

func (s *service) GetConfig() (*system.Config, error) {
	conf := system.NewDefaultConfig()
	if err := s.col.FindOne(context.TODO(), bson.M{"_id": conf.Version}).Decode(conf); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("version: %s system config not found", conf.Version)
		}

		return nil, exception.NewInternalServerError("find system config %s error, %s", conf.Version, err)
	}

	return conf, nil
}
