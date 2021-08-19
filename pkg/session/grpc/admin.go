package grpc

import (
	"context"

	"github.com/infraboard/keyauth/pkg/session"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/mongo"
)

type adminimpl struct {
	*service
	session.UnimplementedAdminServiceServer
}

func (s *adminimpl) QueryUserLastSession(ctx context.Context, req *session.QueryUserLastSessionRequest) (*session.Session, error) {
	r, err := newQueryUserLastSessionRequest(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate query session request error, %s", err)
	}

	ins := session.NewDefaultSession()
	if err := s.col.FindOne(context.TODO(), r.FindFilter(), r.FindOptions()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("session %s not found", req)
		}

		return nil, exception.NewInternalServerError("find session %s error, %s", req.Account, err)
	}
	return ins, nil
}
