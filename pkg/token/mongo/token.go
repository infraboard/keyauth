package mongo

import (
	"context"

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

	tk, err := s.queryToken(newQueryTokenRequestWithAccess(req.AccessToken))
	if err != nil {
		return nil, exception.NewUnauthorized(err.Error())
	}

	if tk.CheckAccessIsExpired() {
		return nil, exception.NewTokenExpired("access_token: %s has expired", tk.AccessToken)
	}

	return tk, nil
}

func (s *service) RevolkToken(req *queryTokenRequest) error {
	_, err := s.col.DeleteOne(context.TODO(), req.FindFilter())
	if err != nil {
		return exception.NewInternalServerError("delete token(%s) error, %s", req, err)
	}
	return nil
}

func (s *service) queryToken(req *queryTokenRequest) (*token.Token, error) {
	tk := new(token.Token)

	if err := s.col.FindOne(context.TODO(), req.FindFilter()).Decode(tk); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("token %s not found", req)
		}

		return nil, exception.NewInternalServerError("find token %s error, %s", req, err)
	}

	return tk, nil
}

func newQueryTokenRequestWithAccess(token string) *queryTokenRequest {
	return &queryTokenRequest{
		AccessToken: token,
	}
}

func newQueryTokenRequestWithRefresh(token string) *queryTokenRequest {
	return &queryTokenRequest{
		RefreshToken: token,
	}
}

type queryTokenRequest struct {
	AccessToken  string
	RefreshToken string
}

func (req *queryTokenRequest) FindFilter() bson.M {
	filter := bson.M{}
	if req.AccessToken != "" {
		filter["access_token"] = req.AccessToken
	}
	if req.RefreshToken != "" {
		filter["refresh_token"] = req.RefreshToken
	}

	return filter
}
