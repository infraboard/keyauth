package impl

import (
	"github.com/infraboard/keyauth/app/wxwork"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newDescribeWechatWorkRequest(req *wxwork.DescribeWechatWorkConf) (*describeWechatWorkRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &describeWechatWorkRequest{req}, nil
}

type describeWechatWorkRequest struct {
	*wxwork.DescribeWechatWorkConf
}

func (r *describeWechatWorkRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Domain != "" {
		filter["_id"] = r.Domain
	}

	return filter
}

// newQueryConfigRequest 查询
func newQueryConfigRequest(req *wxwork.QueryConfigRequest) *queryConfigRequest {
	return &queryConfigRequest{req}
}

type queryConfigRequest struct {
	*wxwork.QueryConfigRequest
}

func (r *queryConfigRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.PageSize)
	skip := int64(r.PageSize) * int64(r.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryConfigRequest) FindFilter() bson.M {
	filter := bson.M{}

	return filter
}
