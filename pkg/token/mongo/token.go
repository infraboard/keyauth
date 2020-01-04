package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"

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
			req.AccessToken, err)
	}

	return tk, nil
}

func (s *service) ValidateToken(accessToken, endpoint string) (*token.Token, error) {
	return nil, nil
}

func (s *service) RevolkToken(accessToken string) error {
	_, err := s.col.DeleteOne(context.TODO(), bson.M{"access_token": accessToken})
	if err != nil {
		return exception.NewInternalServerError("delete token(%s) error, %s", accessToken, err)
	}
	return nil
}

func (s *service) newTokenIssuer(req *token.IssueTokenRequest) (*TokenIssuer, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	issuer := &TokenIssuer{
		IssueTokenRequest: req,
		app:               s.app,
		user:              s.user,
	}
	return issuer, nil
}
