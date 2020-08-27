package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/audit"
)

func newQueryLoginLogRequest(req *audit.QueryLoginRecordRequest) (*queryLoginLogRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &queryLoginLogRequest{
		QueryLoginRecordRequest: req,
	}, nil
}

type queryLoginLogRequest struct {
	*audit.QueryLoginRecordRequest
}

func (r *queryLoginLogRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.PageSize)
	skip := int64(r.PageSize) * int64(r.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "login_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryLoginLogRequest) FindFilter() bson.M {
	tk := r.GetToken()
	filter := bson.M{
		"domain": tk.Domain,
	}

	if r.Account != "" {
		filter["account"] = r.Account
	}

	if r.ApplicationID != "" {
		filter["application_id"] = r.ApplicationID
	}

	if r.GrantType != "" {
		filter["grant_type"] = r.GrantType
	}

	return filter
}
