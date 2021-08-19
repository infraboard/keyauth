package grpc

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/token"
)

func newDescribeTokenRequestWithAccess(token string) *describeTokenRequest {
	return &describeTokenRequest{
		AccessToken: token,
	}
}

func newDescribeTokenRequestWithRefresh(token string) *describeTokenRequest {
	return &describeTokenRequest{
		RefreshToken: token,
	}
}

func newDescribeTokenRequest(req *token.DescribeTokenRequest) *describeTokenRequest {
	return &describeTokenRequest{
		AccessToken:  req.AccessToken,
		RefreshToken: req.RefreshToken,
	}
}

type describeTokenRequest struct {
	AccessToken  string
	RefreshToken string
}

func (req *describeTokenRequest) String() string {
	return fmt.Sprintf("access_token: %s, refresh_token: %s",
		req.AccessToken, req.RefreshToken)
}

func (req *describeTokenRequest) FindFilter() bson.M {
	filter := bson.M{}
	if req.AccessToken != "" {
		filter["_id"] = req.AccessToken
	}
	if req.RefreshToken != "" {
		filter["refresh_token"] = req.RefreshToken
	}

	return filter
}

func newQueryRequest(req *token.QueryTokenRequest) *queryRequest {
	return &queryRequest{req}
}

type queryRequest struct {
	*token.QueryTokenRequest
}

func (r *queryRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{"create_at", -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryRequest) FindFilter() bson.M {
	filter := bson.M{}
	if r.ApplicationId != "" {
		filter["application_id"] = r.ApplicationId
	}
	if r.GrantType != token.GrantType_NULL {
		filter["grant_type"] = r.GrantType
	}
	if r.Account != "" {
		filter["account"] = r.Account
	}
	return filter
}

func newDeleteTokenRequest(req *token.DeleteTokenRequest) *deleteTokenRequest {
	return &deleteTokenRequest{
		DeleteTokenRequest: req,
	}
}

type deleteTokenRequest struct {
	*token.DeleteTokenRequest
}

func (req *deleteTokenRequest) String() string {
	return fmt.Sprintf("access_token: %s",
		req.AccessToken)
}

func (req *deleteTokenRequest) FindFilter() bson.M {
	filter := bson.M{}

	filter["domain"] = req.Domain
	filter["account"] = req.Account

	if len(req.AccessToken) > 0 {
		filter["_id"] = bson.M{"$in": req.AccessToken}
	}

	return filter
}
