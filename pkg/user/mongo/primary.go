package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/infraboard/keyauth/pkg/user"
)

func (s *service) CreatePrimayAccount(req *user.CreateUserRequest) (*user.User, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	user := user.NewUser(req)
	user.Primary = true
	_, err := s.uc.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, exception.NewInternalServerError("inserted user(%s) document error, %s",
			req.Account, err)
	}

	return user, nil
}

func (s *service) DeletePrimaryAccount(id string) error {
	_, err := s.uc.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return exception.NewInternalServerError("delete user(%s) error, %s", id, err)
	}
	return nil
}
