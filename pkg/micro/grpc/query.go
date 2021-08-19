package grpc

import (
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newPaggingQuery(req *micro.QueryMicroRequest) *queryRequest {
	return &queryRequest{req}
}

type queryRequest struct {
	*micro.QueryMicroRequest
}

func (r *queryRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{bson.E{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Type != micro.Type_NULL {
		filter["type"] = r.Type
	}

	return filter
}

func newDescribeQuery(req *micro.DescribeMicroRequest) (*describeRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &describeRequest{req}, nil
}

type describeRequest struct {
	*micro.DescribeMicroRequest
}

func (r *describeRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Name != "" {
		filter["name"] = r.Name
	}
	if r.Id != "" {
		filter["_id"] = r.Id
	}
	if r.ClientId != "" {
		filter["client_id"] = r.ClientId
	}

	return filter
}
