package mongo

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/infraboard/keyauth/pkg/verifycode"
)

func (s *service) delete(ins *verifycode.Code) error {
	result, err := s.col.DeleteOne(context.TODO(), bson.M{"_id": ins.HashID()})
	if err != nil {
		return exception.NewInternalServerError("delete verify code(%s) error, %s", ins.Number, err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("delete verify code %s not found", ins.Number)
	}

	return nil
}
