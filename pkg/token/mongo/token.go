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
			req.AccessToken, err)
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

	tk, err := s.queryToken(req.AccessToken)
	if err != nil {
		return nil, exception.NewUnauthorized(err.Error())
	}

	return tk, nil
}

func (s *service) RevolkToken(accessToken string) error {
	_, err := s.col.DeleteOne(context.TODO(), bson.M{"access_token": accessToken})
	if err != nil {
		return exception.NewInternalServerError("delete token(%s) error, %s", accessToken, err)
	}
	return nil
}

func (s *service) queryToken(accessToken string) (*token.Token, error) {
	tk := new(token.Token)

	if err := s.col.FindOne(context.TODO(), bson.M{"access_token": accessToken}).Decode(tk); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("token %s not found", accessToken)
		}

		return nil, exception.NewInternalServerError("find token %s error, %s", accessToken, err)
	}

	return tk, nil
}
