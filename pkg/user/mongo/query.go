package mongo

import (
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/user/types"
)

func newPaggingQuery(req *user.QueryAccountRequest) *queryUserRequest {
	return &queryUserRequest{
		QueryAccountRequest: req,
	}
}

type queryUserRequest struct {
	userType types.Type
	*user.QueryAccountRequest
}

func (r *queryUserRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.PageSize)
	skip := int64(r.PageSize) * int64(r.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{"create_at", -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryUserRequest) FindFilter() bson.M {
	filter := bson.M{}
	filter["type"] = r.userType

	return filter
}

func newDescribeRequest(req *user.DescriptAccountRequest) (*describeUserRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &describeUserRequest{req}, nil
}

type describeUserRequest struct {
	*user.DescriptAccountRequest
}

func (r *describeUserRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.ID != "" {
		filter["_id"] = r.ID
	}
	if r.Account != "" {
		filter["account"] = r.Account
	}

	return filter
}
