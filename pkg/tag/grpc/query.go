package grpc

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/tag"
	"github.com/infraboard/keyauth/pkg/token"
)

func newQueryTagKeyRequest(req *tag.QueryTagKeyRequest) (*queryTagKeyRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &queryTagKeyRequest{
		QueryTagKeyRequest: req}, nil
}

type queryTagKeyRequest struct {
	*tag.QueryTagKeyRequest
}

func (r *queryTagKeyRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryTagKeyRequest) FindFilter() bson.M {
	filter := bson.M{}

	// if r.Type != role.RoleType_NULL {
	// 	filter["type"] = r.Type
	// } else {
	// 	// 获取内建和全局的角色以及域内自己创建的角色
	// 	filter["$or"] = bson.A{
	// 		bson.M{"type": role.RoleType_BUILDIN},
	// 		bson.M{"type": role.RoleType_GLOBAL},
	// 		bson.M{"type": role.RoleType_CUSTOM, "domain": r.tk.Domain},
	// 	}
	// }

	return filter
}

func newQueryTagValueRequest(req *tag.QueryTagValueRequest) (*queryTagValueRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &queryTagValueRequest{
		QueryTagValueRequest: req}, nil
}

type queryTagValueRequest struct {
	*tag.QueryTagValueRequest
}

func (r *queryTagValueRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryTagValueRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.TagId != "" {
		filter["key_id"] = r.TagId
	}

	return filter
}

func newDescribeTagRequest(tk *token.Token, req *tag.DescribeTagRequest) (*describeTagRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &describeTagRequest{
		DescribeTagRequest: req}, nil
}

type describeTagRequest struct {
	*tag.DescribeTagRequest
}

func (r *describeTagRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.TagId != "" {
		filter["_id"] = r.TagId
	}

	return filter
}
