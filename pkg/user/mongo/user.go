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
	r, err := newQueryUserRequest(req)
	if err != nil {
		return nil, err
	}

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
		u.Domain = tk.Domain
	}

	u.Type = t
	if err := s.saveAccount(u); err != nil {
		return nil, err
	}

	u.HashedPassword = nil
	return u, nil
}

func (s *service) UpdateAccountProfile(u *user.User) error {
	req, err := newUpdateUserRequest(u)
	if err != nil {
		return err
	}

	_, err = s.col.UpdateOne(context.TODO(), bson.M{"_id": u.Account}, bson.M{"$set": req.updateData()})
	if err != nil {
		return exception.NewInternalServerError("update user(%s) error, %s", u.Account, err)
	}

	return nil
}

func (s *service) UpdateAccountPassword(req *user.UpdatePasswordRequest) (*user.Password, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("check update pass request error, %s", err)
	}

	descReq := user.NewDescriptAccountRequest()
	descReq.Account = req.GetToken().Account
	u, err := s.DescribeAccount(descReq)
	if err != nil {
		return nil, err
	}

	if err := u.ChangePassword(req.OldPass, req.NewPass); err != nil {
		return nil, err
	}

	_, err = s.col.UpdateOne(context.TODO(), bson.M{"_id": u.Account}, bson.M{"$set": bson.M{
		"password": u.HashedPassword,
	}})

	if err != nil {
		return nil, exception.NewInternalServerError("update user(%s) password error, %s", u.Account, err)
	}

	return u.HashedPassword, nil
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

func (s *service) BlockAccount(account, reason string) error {
	desc := user.NewDescriptAccountRequestWithAccount(account)
	user, err := s.DescribeAccount(desc)
	if err != nil {
		return fmt.Errorf("describe user error, %s", err)
	}

	user.Block(reason)
	return s.saveAccount(user)
}

func (s *service) DeleteAccount(account string) error {
	_, err := s.col.DeleteOne(context.TODO(), bson.M{"_id": account})
	if err != nil {
		return exception.NewInternalServerError("delete user(%s) error, %s", account, err)
	}
	return nil
}
