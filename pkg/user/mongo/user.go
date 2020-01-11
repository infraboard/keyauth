package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/user"
)

func (s *service) queryAccount(req *queryRequest) (
	users []*user.User, totalPage int64, err error) {
	resp, err := s.col.Find(context.TODO(), req.FindFilter(), req.FindOptions())

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
	count, err := s.col.CountDocuments(context.TODO(), req.FindFilter())
	if err != nil {
		return nil, 0, exception.NewInternalServerError("get device count error, error is %s", err)
	}
	totalPage = count
	return users, totalPage, nil
}

func (s *service) UpdateAccountPassword(userName, oldPass, newPass string) error {
	return nil
}

func (s *service) DescribeAccount(req *user.DescriptAccountRequest) (*user.User, error) {
	r, err := newDescribeRequest(req)
	if err != nil {
		return nil, err
	}
	user := user.NewDescribeUser()

	if err := s.col.FindOne(context.TODO(), r.FindFilter()).Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("user %s not found", req)
		}

		return nil, exception.NewInternalServerError("find user %s error, %s", req, err)
	}

	return user, nil
}

func newPaggingQuery(req *user.QueryAccountRequest) *queryRequest {
	return &queryRequest{
		QueryAccountRequest: req,
	}
}

type queryRequest struct {
	userType user.Type
	*user.QueryAccountRequest
}

func (r *queryRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.PageSize)
	skip := int64(r.PageSize) * int64(r.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{"create_at", -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryRequest) FindFilter() bson.M {
	filter := bson.M{}
	filter["type"] = r.userType

	return filter
}

func newDescribeRequest(req *user.DescriptAccountRequest) (*describeRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &describeRequest{req}, nil
}

type describeRequest struct {
	*user.DescriptAccountRequest
}

func (r *describeRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.ID != "" {
		filter["_id"] = r.ID
	}
	if r.Account != "" {
		filter["account"] = r.Account
	}

	return filter
}
