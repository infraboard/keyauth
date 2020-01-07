package mongo

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg/token"
)

func (s *service) IssueToken(req *token.IssueTokenRequest) (*token.Token, error) {
	issuer, err := s.newTokenIssuer(req)
	if err != nil {
		return nil, exception.NewUnauthorized(err.Error())
	}

	tk, err := issuer.IssueToken()
	if err != nil {
		return nil, err
	}

	if _, err := s.col.InsertOne(context.TODO(), tk); err != nil {
		return nil, exception.NewInternalServerError("inserted token(%s) document error, %s",
			tk.AccessToken, err)
	}

	return tk, nil
}

func (s *service) ValidateToken(req *token.ValidateTokenRequest) (*token.Token, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	ck := newClientChecker(s.app)
	if _, err := ck.CheckClient(req.ClientID, req.ClientSecret); err != nil {
		return nil, exception.NewUnauthorized(err.Error())
	}

	tk, err := s.describeToken(newDescribeTokenRequestWithAccess(req.AccessToken))
	if err != nil {
		return nil, exception.NewUnauthorized(err.Error())
	}

	if tk.CheckAccessIsExpired() {
		return nil, exception.NewAccessTokenExpired("access_token: %s has expired", tk.AccessToken)
	}

	// 校验用户权限
	if req.Endpoint != "" {
		// 找到用户角色
		// 判断该角色是否有该Endpoint调用权限
	}

	return tk, nil
}

func (s *service) RevolkToken(req *token.DescribeTokenRequest) error {
	if err := req.Validate(); err != nil {
		return exception.NewBadRequest(err.Error())
	}

	ck := newClientChecker(s.app)
	if _, err := ck.CheckClient(req.ClientID, req.ClientSecret); err != nil {
		return exception.NewUnauthorized(err.Error())
	}

	descReq := newDescribeTokenRequestWithAccess(req.AccessToken)
	return s.destoryToken(descReq)
}

func (s *service) destoryToken(req *describeTokenRequest) error {
	resp, err := s.col.DeleteOne(context.TODO(), req.FindFilter())
	if err != nil {
		return exception.NewInternalServerError("delete token(%s) error, %s", req, err)
	}

	if resp.DeletedCount == 0 {
		return exception.NewNotFound("token(%s) not found", req)
	}

	return nil
}

func (s *service) describeToken(req *describeTokenRequest) (*token.Token, error) {
	tk := new(token.Token)

	if err := s.col.FindOne(context.TODO(), req.FindFilter()).Decode(tk); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("token %s not found", req)
		}

		return nil, exception.NewInternalServerError("find token %s error, %s", req, err)
	}

	return tk, nil
}

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
		filter["access_token"] = req.AccessToken
	}
	if req.RefreshToken != "" {
		filter["refresh_token"] = req.RefreshToken
	}

	return filter
}
