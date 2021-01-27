package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/common/session"
	"github.com/infraboard/keyauth/pkg/application"
)

type adminimpl struct {
	*service
	application.UnimplementedAdminServiceServer
}

func (s *adminimpl) GetBuildInApplication(ctx context.Context, req *application.GetBuildInApplicationRequest) (
	*application.Application, error) {
	app := new(application.Application)
	if err := s.col.FindOne(context.TODO(), bson.M{"name": req.Name, "build_in": true}).Decode(app); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("applicaiton %s not found", req.Name)
		}

		return nil, exception.NewInternalServerError("find application %s error, %s", req.Name, err)
	}

	return app, nil
}

func (s *adminimpl) CreateBuildInApplication(ctx context.Context, req *application.CreateApplicatonRequest) (
	*application.Application, error) {
	account := session.GetTokenFromContext(ctx).Account
	app, err := application.NewBuildInApplication(account, req)
	if err != nil {
		return nil, err
	}

	return s.save(app)
}
