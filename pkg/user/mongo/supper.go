package mongo

import (
	"context"

	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/mcube/exception"
)

func (s *service) CreateSupperAccount(req *user.CreateUserRequest) (*user.User, error) {
	u, err := user.New(req)
	if err != nil {
		return nil, err
	}

	u.Type = user.SubAccount
	if err := s.saveAccount(u); err != nil {
		return nil, err
	}

	u.HashedPassword = nil
	return u, nil
}

func (s *service) QueryAccount(req *user.QueryAccountRequest) (
	users []*user.User, totalPage int64, err error) {
	r := newPaggingQuery(req)
	resp, err := s.col.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, 0, exception.NewInternalServerError("find user error, error is %s", err)
	}

	// 循环
	for resp.Next(context.TODO()) {
		u := new(user.User)
		if err := resp.Decode(u); err != nil {
			return nil, 0, exception.NewInternalServerError("decode user error, error is %s", err)
		}

		users = append(users, u)
	}

	// count
	count, err := s.col.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, 0, exception.NewInternalServerError("get device count error, error is %s", err)
	}
	totalPage = count
	return users, totalPage, nil
}

func (s *service) saveAccount(u *user.User) error {
	if _, err := s.col.InsertOne(context.TODO(), u); err != nil {
		return exception.NewInternalServerError("inserted user(%s) document error, %s",
			u.Account, err)
	}

	return nil
}
