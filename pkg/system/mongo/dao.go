package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/infraboard/keyauth/pkg/system"
	"github.com/infraboard/keyauth/pkg/system/notify/mail"
	"github.com/infraboard/keyauth/pkg/system/notify/sms"
	"github.com/infraboard/keyauth/pkg/verifycode"
)

func (s *service) save(conf *system.Config) error {
	if _, err := s.col.InsertOne(context.TODO(), conf); err != nil {
		return exception.NewInternalServerError("save config document error, %s", err)
	}
	return nil
}

func (s *service) updateEmail(conf *mail.Config) error {
	_, err := s.col.UpdateOne(context.TODO(), bson.M{"_id": system.DEFAULT_CONFIG_VERSION}, bson.M{"$set": bson.M{
		"email": conf,
	}})
	if err != nil {
		return exception.NewInternalServerError("update email config document error, %s", err)
	}

	return nil
}

func (s *service) updateSMS(conf *sms.Config) error {
	_, err := s.col.UpdateOne(context.TODO(), bson.M{"_id": system.DEFAULT_CONFIG_VERSION}, bson.M{"$set": bson.M{
		"sms": conf,
	}})
	if err != nil {
		return exception.NewInternalServerError("update sms config document error, %s", err)
	}

	return nil
}

func (s *service) updateVerifyCode(conf *verifycode.Config) error {
	_, err := s.col.UpdateOne(context.TODO(), bson.M{"_id": system.DEFAULT_CONFIG_VERSION}, bson.M{"$set": bson.M{
		"verify_code": conf,
	}})
	if err != nil {
		return exception.NewInternalServerError("update verify code config document error, %s", err)
	}

	return nil
}
