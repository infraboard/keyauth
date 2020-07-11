package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"

	"github.com/infraboard/keyauth/pkg/user"
)

func (s *service) saveAccount(u *user.User) error {
	if _, err := s.col.InsertOne(context.TODO(), u); err != nil {
		return exception.NewInternalServerError("inserted user(%s) document error, %s",
			u.Account, err)
	}

	return nil
}

func (s *service) queryAccount(req *queryUserRequest) (*user.Set, error) {
	resp, err := s.col.Find(context.TODO(), req.FindFilter(), req.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find user error, error is %s", err)
	}

	userSet := user.NewUserSet(req.PageRequest)
	// 循环
	for resp.Next(context.TODO()) {
		u := new(user.User)
		if err := resp.Decode(u); err != nil {
			return nil, exception.NewInternalServerError("decode user error, error is %s", err)
		}
		userSet.Add(u)
	}

	// count
	count, err := s.col.CountDocuments(context.TODO(), req.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get device count error, error is %s", err)
	}
	userSet.Total = count

	return userSet, nil
}
