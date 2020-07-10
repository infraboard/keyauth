package mongo

import (
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/application"
)

func newPaggingQuery(req *application.QueryApplicationRequest) *queryRequest {
	return &queryRequest{req}
}

type queryRequest struct {
	*application.QueryApplicationRequest
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

	return filter
}

func newDescribeQuery(req *application.DescriptApplicationRequest) (*describeRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &describeRequest{req}, nil
}

type describeRequest struct {
	*application.DescriptApplicationRequest
}

func (r *describeRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.ID != "" {
		filter["_id"] = r.ID
	}
	if r.ClientID != "" {
		filter["client_id"] = r.ClientID
	}

	return filter
}
