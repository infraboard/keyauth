package mongo

import (
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/namespace"
	"github.com/infraboard/keyauth/pkg/token"
)

func newPaggingQuery(req *namespace.QueryNamespaceRequest) *queryNamespaceRequest {
	return &queryNamespaceRequest{req}
}

type queryNamespaceRequest struct {
	*namespace.QueryNamespaceRequest
}

func (r *queryNamespaceRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryNamespaceRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.DepartmentId != "" {
		if r.WithSubDepartment {
			filter["department_id"] = bson.M{"$regex": r.DepartmentId, "$options": "im"}
		} else {
			filter["department_id"] = r.DepartmentId
		}
	}

	if r.Name != "" {
		filter["name"] = r.Name
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

	if r.Id != "" {
		filter["_id"] = r.Id
	}

	return filter
}

func newDeleteRequest(tk *token.Token, req *namespace.DeleteNamespaceRequest) (*deleteNamespaceRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &deleteNamespaceRequest{
		tk:                     tk,
		DeleteNamespaceRequest: req,
	}, nil
}

type deleteNamespaceRequest struct {
	tk *token.Token
	*namespace.DeleteNamespaceRequest
}

func (r *deleteNamespaceRequest) FindFilter() bson.M {
	filter := bson.M{
		"domain": r.tk.Domain,
		"_id":    r.Id,
	}

	return filter
}
