package grpc

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/mconf"
	"github.com/infraboard/mcube/exception"
)

func newGroupPaggingQuery(req *mconf.QueryGroupRequest) *queryGroupRequest {
	return &queryGroupRequest{req}
}

type queryGroupRequest struct {
	*mconf.QueryGroupRequest
}

func (r *queryGroupRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{bson.E{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryGroupRequest) FindFilter() bson.M {
	filter := bson.M{}

	return filter
}

func newDescribeGroupQuery(req *mconf.DescribeGroupRequest) (*describeGroupRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &describeGroupRequest{req}, nil
}

type describeGroupRequest struct {
	*mconf.DescribeGroupRequest
}

func (r *describeGroupRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Name != "" {
		filter["_id"] = r.Name
	}

	return filter
}

func newItemPaggingQuery(req *mconf.QueryItemRequest) *queryItemRequest {
	return &queryItemRequest{req}
}

type queryItemRequest struct {
	*mconf.QueryItemRequest
}

func (r *queryItemRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{bson.E{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryItemRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.GroupName != "" {
		filter["group"] = r.GroupName
	}

	return filter
}
