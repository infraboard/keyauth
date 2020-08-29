package mongo

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/role"
)

func newDescribeRoleRequest(req *role.DescribeRoleRequest) (*describeRoleRequest, error) {
	if err := req.Valiate(); err != nil {
		return nil, err
	}
	return &describeRoleRequest{req}, nil
}

type describeRoleRequest struct {
	*role.DescribeRoleRequest
}

func (req *describeRoleRequest) String() string {
	return fmt.Sprintf("role: %s", req.Name)
}

func (req *describeRoleRequest) FindFilter() bson.M {
	filter := bson.M{}

	if req.ID != "" {
		filter["_id"] = req.ID
	}

	if req.Name != "" {
		filter["name"] = req.Name
	}

	return filter
}

// FindOptions todo
func (req *describeRoleRequest) FindOptions() *options.FindOneOptions {
	opt := &options.FindOneOptions{}

	if !req.WithPermissions {
		opt.Projection = bson.M{"permissions": 0}
	}

	return opt
}

func newQueryRoleRequest(req *role.QueryRoleRequest) (*queryRoleRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &queryRoleRequest{req}, nil
}

type queryRoleRequest struct {
	*role.QueryRoleRequest
}

func (r *queryRoleRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.PageSize)
	skip := int64(r.PageSize) * int64(r.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	if !r.WithPermissions {
		opt.Projection = bson.M{"permissions": 0}
	}

	return opt
}

func (r *queryRoleRequest) FindFilter() bson.M {
	filter := bson.M{}
	if r.Type.String() != "unknown" {
		filter["type"] = r.Type.String()
	}

	return filter
}
