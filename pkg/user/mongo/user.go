package mongo

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg/user"
)

func (s *service) queryAccount(req *queryRequest) (*user.Set, error) {
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

func (s *service) UpdateAccountPassword(userName, oldPass, newPass string) error {
	return nil
}

func (s *service) DescribeAccount(req *user.DescriptAccountRequest) (*user.User, error) {
	r, err := newDescribeRequest(req)
	if err != nil {
		return nil, err
	}
	user := user.NewDefaultUser()

	if err := s.col.FindOne(context.TODO(), r.FindFilter()).Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("user %s not found", req)
		}

		return nil, exception.NewInternalServerError("find user %s error, %s", req, err)
	}

	return user, nil
}

func (s *service) BlockAccount(id, reason string) error {
	desc := user.NewDescriptAccountRequestWithID(id)
	user, err := s.DescribeAccount(desc)
	if err != nil {
		return fmt.Errorf("describe user error, %s", err)
	}

	user.Block(reason)
	return s.saveAccount(user)
}
