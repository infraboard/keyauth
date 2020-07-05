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

	u.Type = user.SupperAccount
	if err := s.saveAccount(u); err != nil {
		return nil, err
	}

	u.HashedPassword = nil
	return u, nil
}

func (s *service) QuerySupperAccount(req *user.QueryAccountRequest) (*user.Set, error) {
	r := newPaggingQuery(req)
	r.userType = user.SupperAccount
	return s.queryAccount(r)
}

func (s *service) saveAccount(u *user.User) error {
	if _, err := s.col.InsertOne(context.TODO(), u); err != nil {
		return exception.NewInternalServerError("inserted user(%s) document error, %s",
			u.Account, err)
	}

	return nil
}
