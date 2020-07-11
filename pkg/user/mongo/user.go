package mongo

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/user/types"
)

func (s *service) QueryAccount(t types.Type, req *user.QueryAccountRequest) (*user.Set, error) {
	r := newPaggingQuery(req)
	r.userType = t
	return s.queryAccount(r)
}

func (s *service) CreateAccount(t types.Type, req *user.CreateUserRequest) (*user.User, error) {
	u, err := user.New(req)
	if err != nil {
		return nil, err
	}

	tk := req.GetToken()
	if tk != nil {
		u.DomainID = tk.DomainID
	}

	u.Type = t
	if err := s.saveAccount(u); err != nil {
		return nil, err
	}

	u.HashedPassword = nil
	return u, nil
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

func (s *service) DeleteAccount(id string) error {
	_, err := s.col.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return exception.NewInternalServerError("delete user(%s) error, %s", id, err)
	}
	return nil
}
