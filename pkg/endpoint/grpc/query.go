package grpc

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/mcube/exception"
)

func newDescribeEndpointRequest(req *endpoint.DescribeEndpointRequest) (*describeEndpointRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &describeEndpointRequest{req}, nil
}

type describeEndpointRequest struct {
	*endpoint.DescribeEndpointRequest
}

func (req *describeEndpointRequest) String() string {
	return fmt.Sprintf("endpoint: %s", req.Id)
}

func (req *describeEndpointRequest) FindFilter() bson.M {
	filter := bson.M{}
	if req.Id != "" {
		filter["_id"] = req.Id
	}

	return filter
}

func newQueryEndpointRequest(req *endpoint.QueryEndpointRequest) *queryEndpointRequest {
	return &queryEndpointRequest{req}
}

type queryEndpointRequest struct {
	*endpoint.QueryEndpointRequest
}

func (r *queryEndpointRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryEndpointRequest) FindFilter() bson.M {
	filter := bson.M{}

	if len(r.ServiceIds) > 0 {
		filter["service_id"] = bson.M{"$in": r.ServiceIds}
	}
	if r.Method != "" {
		filter["entry.method"] = r.Method
	}
	if len(r.Resources) > 0 {
		filter["entry.resource"] = bson.M{"$in": r.Resources}
	}
	if r.Path != "" {
		filter["entry.path"] = bson.M{"$regex": r.Path, "$options": "im"}
	}
	if r.FunctionName != "" {
		filter["entry.function_name"] = r.FunctionName
	}
	switch r.PermissionEnable {
	case endpoint.BoolQuery_TRUE:
		filter["entry.permission_enable"] = true
	case endpoint.BoolQuery_FALSE:
		filter["entry.permission_enable"] = false
	}

	return filter
}
