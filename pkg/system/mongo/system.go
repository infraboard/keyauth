package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg/system"
	"github.com/infraboard/keyauth/pkg/system/notify/mail"
	"github.com/infraboard/keyauth/pkg/system/notify/sms"
	"github.com/infraboard/keyauth/pkg/verifycode"
)

func (s *service) UpdateEmail(mailConf *mail.Config) error {
	if err := mailConf.Validate(); err != nil {
		return exception.NewBadRequest("validate mail config error, %s", err)
	}

	return s.updateEmail(mailConf)
}

func (s *service) UpdateSMS(smsConf *sms.Config) error {
	if err := smsConf.Validate(); err != nil {
		return exception.NewBadRequest("validate mail config error, %s", err)
	}

	return s.updateSMS(smsConf)
}

func (s *service) UpdateVerifyCode(vfconf *verifycode.Config) error {
	if err := vfconf.Validate(); err != nil {
		return exception.NewBadRequest("validate verify code config error, %s", err)
	}

	conf, err := s.GetConfig()
	if err != nil {
		return err
	}

	// 校验配置的通知方式是否已经配置
	switch vfconf.NotifyType {
	case verifycode.NotifyType_MAIL:
		if err := conf.Email.Validate(); err != nil {
			s.log.Infof("email setting validate error, %s", err)
			return exception.NewBadRequest("系统邮件未配置, 或配置有误")
		}
	// 校验配置的通知方式是否已经配置
	case verifycode.NotifyType_SMS:
		if err := conf.SMS.Validate(); err != nil {
			s.log.Infof("sms setting validate error, %s", err)
			return exception.NewBadRequest("系统短信未配置, 或配置有误")
		}
	}

	return s.updateVerifyCode(vfconf)
}

func (s *service) GetConfig() (*system.Config, error) {
	conf := system.NewDefaultConfig()
	if err := s.col.FindOne(context.TODO(), bson.M{"_id": conf.Version}).Decode(conf); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("version: %s system config not found, please init first", conf.Version)
		}

		return nil, exception.NewInternalServerError("find system config %s error, %s", conf.Version, err)
	}

	return conf, nil
}

func (s *service) InitConfig(conf *system.Config) error {
	return s.save(conf)
}
