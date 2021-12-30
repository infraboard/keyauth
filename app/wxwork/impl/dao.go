package impl

import (
	"context"
	"github.com/infraboard/keyauth/app/wxwork"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *service) save(ins *wxwork.WechatWorkConfig) error {
	if _, err := s.col.InsertOne(context.TODO(), ins); err != nil {
		return exception.NewInternalServerError("inserted wechat work config(%s) document error, %s",
			ins.Domain, err)
	}
	return nil
}

func (s *service) update(ins *wxwork.WechatWorkConfig) error {
	ins.UpdateAt = ftime.Now()
	_, err := s.col.UpdateOne(context.TODO(), bson.M{"_id": ins.Domain}, bson.M{"$set": ins})
	if err != nil {
		return exception.NewInternalServerError("update domain(%s) error, %s", ins.Domain, err)
	}

	return nil
}

