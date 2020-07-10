package mongo

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/policy"
)

func newDescribeEndpointRequest(req *policy.DescribePolicyRequest) *describePolicyRequest {
	return &describePolicyRequest{req}
}

type describePolicyRequest struct {
	*policy.DescribePolicyRequest
}

func (req *describePolicyRequest) String() string {
	return fmt.Sprintf("policy: %s", req.ID)
}

func (req *describePolicyRequest) FindFilter() bson.M {
	filter := bson.M{}
	if req.ID != "" {
		filter["id"] = req.ID
	}

	return filter
}

func newQueryRequest(req *policy.QueryPolicyRequest) *queryPolicyRequest {
	return &queryPolicyRequest{req}
}

type queryPolicyRequest struct {
	*policy.QueryPolicyRequest
}

func (r *queryPolicyRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.PageSize)
	skip := int64(r.PageSize) * int64(r.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{"create_at", -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryPolicyRequest) FindFilter() bson.M {
	filter := bson.M{}

	return filter
}
