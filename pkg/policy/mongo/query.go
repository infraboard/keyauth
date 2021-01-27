package mongo

import (
	"fmt"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/token"
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
		filter["id"] = req.Id
	}

	return filter
}

func newQueryPolicyRequest(tk *token.Token, req *policy.QueryPolicyRequest) (*queryPolicyRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &queryPolicyRequest{
		tk:                 tk,
		QueryPolicyRequest: req,
	}, nil
}

type queryPolicyRequest struct {
	tk *token.Token
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
	tk := r.tk

	filter := bson.M{}
	filter["domain"] = tk.Domain

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

func newDeletePolicyRequest(tk *token.Token, req *policy.DeletePolicyRequest) (*deletePolicyRequest, error) {
	return &deletePolicyRequest{
		tk:                  tk,
		DeletePolicyRequest: req,
	}, nil
}

type deletePolicyRequest struct {
	tk *token.Token
	*policy.DeletePolicyRequest
}

func (r *deletePolicyRequest) FindFilter() bson.M {
	tk := r.tk

	filter := bson.M{}
	filter["domain"] = tk.Domain

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
