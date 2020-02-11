package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/service"
	"github.com/infraboard/mcube/exception"
)

func newPaggingQuery(req *service.QueryServiceRequest) *queryRequest {
	return &queryRequest{req}
}

type queryRequest struct {
	*service.QueryServiceRequest
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

func newDescribeQuery(req *service.DescriptServiceRequest) (*describeRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &describeRequest{req}, nil
}

type describeRequest struct {
	*service.DescriptServiceRequest
}

func (r *describeRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Name != "" {
		filter["_id"] = r.Name
	}
	if r.ServiceID != "" {
		filter["service_id"] = r.ServiceID
	}

	return filter
}
