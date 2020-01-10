package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/infraboard/keyauth/pkg/user"
)

func (s *service) CreatePrimayAccount(req *user.CreateUserRequest) (*user.User, error) {
	u, err := user.New(req)
	if err != nil {
		return nil, err
	}

	u.Type = user.PrimaryAccount
	if err := s.saveAccount(u); err != nil {
		return nil, err
	}

	u.HashedPassword = nil
	return u, nil
}

func (s *service) DeletePrimaryAccount(id string) error {
	_, err := s.col.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return exception.NewInternalServerError("delete user(%s) error, %s", id, err)
	}
	return nil
}
