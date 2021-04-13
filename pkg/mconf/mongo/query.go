package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/mconf"
)

func newGroupPaggingQuery(req *mconf.QueryGroupRequest) *queryGroupRequest {
	return &queryGroupRequest{req}
}

type queryGroupRequest struct {
	*mconf.QueryGroupRequest
}

func (r *queryGroupRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{bson.E{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryGroupRequest) FindFilter() bson.M {
	filter := bson.M{}

	return filter
}
