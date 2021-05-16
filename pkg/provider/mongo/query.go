package mongo

import (
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/provider"
)

func newQueryLDAPRequest(req *provider.QueryLDAPConfigRequest) *queryLDAPRequest {
	return &queryLDAPRequest{req}
}

type queryLDAPRequest struct {
	*provider.QueryLDAPConfigRequest
}

func (r *queryLDAPRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.PageSize)
	skip := int64(r.PageSize) * int64(r.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryLDAPRequest) FindFilter() bson.M {
	filter := bson.M{}

	return filter
}

func newDescribeLDAPRequest(req *provider.DescribeLDAPConfig) (*describeLDAPRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &describeLDAPRequest{req}, nil
}

type describeLDAPRequest struct {
	*provider.DescribeLDAPConfig
}

func (r *describeLDAPRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Domain != "" {
		filter["_id"] = r.Domain
	}
	if r.BaseDN != "" {
		filter["base_dn"] = bson.M{"$regex": r.BaseDN + "$", "$options": "im"}
	}

	return filter
}
