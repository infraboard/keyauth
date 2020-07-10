package mongo

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/role"
)

func newDescribeRoleRequest(req *role.DescribeRoleRequest) *describeRoleRequest {
	return &describeRoleRequest{req}
}

type describeRoleRequest struct {
	*role.DescribeRoleRequest
}

func (req *describeRoleRequest) String() string {
	return fmt.Sprintf("role: %s", req.Name)
}

func (req *describeRoleRequest) FindFilter() bson.M {
	filter := bson.M{}
	if req.Name != "" {
		filter["name"] = req.Name
	}

	return filter
}

func newQueryRequest(req *role.QueryRoleRequest) *queryRoleRequest {
	return &queryRoleRequest{req}
}

type queryRoleRequest struct {
	*role.QueryRoleRequest
}

func (r *queryRoleRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.PageSize)
	skip := int64(r.PageSize) * int64(r.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{"create_at", -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryRoleRequest) FindFilter() bson.M {
	filter := bson.M{}
	if r.Type.String() != "" {
		filter["type"] = r.Type.String()
	}

	return filter
}
