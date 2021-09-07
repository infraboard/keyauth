package grpc

import (
	"fmt"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/policy"
)

func newDescribePolicyRequest(req *policy.DescribePolicyRequest) (*describePolicyRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}
	return &describePolicyRequest{req}, nil
}

type describePolicyRequest struct {
	*policy.DescribePolicyRequest
}

func (req *describePolicyRequest) String() string {
	return fmt.Sprintf("policy: %s", req.Id)
}

func (req *describePolicyRequest) FindFilter() bson.M {
	filter := bson.M{}
	if req.Id != "" {
		filter["_id"] = req.Id
	}
	return filter
}

func newQueryPolicyRequest(req *policy.QueryPolicyRequest) (*queryPolicyRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &queryPolicyRequest{
		QueryPolicyRequest: req,
	}, nil
}

type queryPolicyRequest struct {
	*policy.QueryPolicyRequest
}

func (r *queryPolicyRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryPolicyRequest) FindFilter() bson.M {
	filter := bson.M{}
	if r.Domain != "" {
		filter["domain"] = r.Domain
	}

	if r.NamespaceId != "" {
		filter["namespace_id"] = r.NamespaceId
	}
	if r.RoleId != "" {
		filter["role_id"] = r.RoleId
	}
	if r.Account != "" {
		filter["account"] = r.Account
	}
	if r.Type != policy.PolicyType_NULL {
		filter["type"] = r.Type
	}

	return filter
}

func newDeletePolicyRequest(req *policy.DeletePolicyRequest) (*deletePolicyRequest, error) {
	return &deletePolicyRequest{
		DeletePolicyRequest: req,
	}, nil
}

type deletePolicyRequest struct {
	*policy.DeletePolicyRequest
}

func (r *deletePolicyRequest) FindFilter() bson.M {
	filter := bson.M{}
	filter["domain"] = r.Domain

	if r.Id != "" {
		filter["_id"] = r.Id
	}
	if r.Account != "" {
		filter["account"] = r.Account
	}
	if r.RoleId != "" {
		filter["role_id"] = r.RoleId
	}
	if r.NamespaceId != "" {
		filter["namespace_id"] = r.NamespaceId
	}
	if r.Type != policy.PolicyType_NULL {
		filter["type"] = r.Type
	}

	return filter
}
