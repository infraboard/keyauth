package mongo

import (
	"context"

	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *service) update(ins *micro.Micro) error {
	ins.UpdateAt = ftime.Now()
	_, err := s.scol.UpdateOne(context.TODO(), bson.M{"_id": ins.ID}, bson.M{"$set": ins})
	if err != nil {
		return exception.NewInternalServerError("update service(%s) error, %s", ins.Name, err)
	}

	return nil
}