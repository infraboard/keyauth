package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/infraboard/keyauth/pkg/application"
	"github.com/infraboard/keyauth/pkg/token"
)

func (s *service) IssueToken(req *token.IssueTokenRequest) (*token.Token, error) {
	app, err := token.New(req)
	if err != nil {
		return nil, err
	}

	if _, err := s.col.InsertOne(context.TODO(), app); err != nil {
		return nil, exception.NewInternalServerError("inserted token(%s) document error, %s",
			req.AccessToken, err)
	}

	appReq := application.NewDescriptApplicationRequest()
	s.app.DescriptionApplication(appReq)

	return app, nil
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
