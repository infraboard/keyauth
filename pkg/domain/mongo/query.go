package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/domain"
)

func newQueryDomainRequest(req *domain.QueryDomainRequest) *request {
	return &request{req}
}

type request struct {
	*domain.QueryDomainRequest
}

func (r *request) FindOptions() *options.FindOptions {
	pageSize := int64(r.PageSize)
	skip := int64(r.PageSize) * int64(r.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{"create_at", -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *request) FindFilter() bson.M {
	filter := bson.M{}

	tk := r.GetToken()
	filter["owner_id"] = tk.UserID

	return filter
}

func newDescDomainRequest(req *domain.DescriptDomainRequest) (*descDomain, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &descDomain{req}, nil
}

type descDomain struct {
	*domain.DescriptDomainRequest
}

func (req *descDomain) FindFilter() bson.M {
	filter := bson.M{}

	if req.ID != "" {
		filter["_id"] = req.ID
	}
	if req.Name != "" {
		filter["name"] = req.Name
	}

	return filter
}
