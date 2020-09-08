package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/namespace"
	"github.com/infraboard/mcube/exception"
)

func newPaggingQuery(req *namespace.QueryNamespaceRequest) *queryNamespaceRequest {
	return &queryNamespaceRequest{req}
}

type queryNamespaceRequest struct {
	*namespace.QueryNamespaceRequest
}

func (r *queryNamespaceRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.PageSize)
	skip := int64(r.PageSize) * int64(r.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryNamespaceRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Department != "" {
		filter["department"] = r.Department
	}

	return filter
}

func newDescribeQuery(req *namespace.DescriptNamespaceRequest) (*describeNamespaceRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &describeNamespaceRequest{req}, nil
}

type describeNamespaceRequest struct {
	*namespace.DescriptNamespaceRequest
}

func (r *describeNamespaceRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.ID != "" {
		filter["_id"] = r.ID
	}

	return filter
}

func newDeleteRequest(req *namespace.DeleteNamespaceRequest) (*deleteNamespaceRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &deleteNamespaceRequest{req}, nil
}

type deleteNamespaceRequest struct {
	*namespace.DeleteNamespaceRequest
}

func (r *deleteNamespaceRequest) FindFilter() bson.M {
	tk := r.GetToken()

	filter := bson.M{
		"domain": tk.Domain,
		"_id":    r.ID,
	}

	return filter
}
