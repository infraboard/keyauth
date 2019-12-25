package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/user"
)

func (s *service) UpdateAccountPassword(userName, oldPass, newPass string) error {
	return nil
}

func (s *service) DescribeAccount(req *user.DescriptAccountRequest) (*user.User, error) {
	user := user.NewDescribeUser()

	if err := s.uc.FindOne(context.TODO(), bson.M{"_id": req.ID}).Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("user %s not found", req.ID)
		}

		return nil, exception.NewInternalServerError("find user %s error, %s", req.ID, err)
	}

	return user, nil
}

type request struct {
	*user.QueryRAMAccountRequest

	opt *options.FindOptions
}

func (r *request) FindOptions() *options.FindOptions {
	pageSize := int64(r.PageSize)
	skip := int64(r.PageSize) * int64(r.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{"create_at", -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *request) FindFilter() bson.M {
	filter := bson.M{}

	return filter
}
