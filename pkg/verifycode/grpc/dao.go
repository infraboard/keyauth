package grpc

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/infraboard/keyauth/pkg/verifycode"
)

func (s *service) delete(ins *verifycode.Code) error {
	ck := verifycode.NewCheckCodeRequest(ins.Username, ins.Number)
	result, err := s.col.DeleteOne(context.TODO(), bson.M{"_id": ck.HashID()})
	if err != nil {
		return exception.NewInternalServerError("delete verify code(%s) error, %s", ins.Number, err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("delete verify code %s not found", ins.Number)
	}

	return nil
}
