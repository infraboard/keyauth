package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/tag"
)

func (s *service) CreateTag(ctx context.Context, req *tag.CreateTagRequest) (*tag.Tag, error) {
	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	r, err := tag.New(tk, req)
	if err != nil {
		return nil, err
	}

	if _, err := s.key.InsertOne(context.TODO(), r); err != nil {
		return nil, exception.NewInternalServerError("inserted tag key(%s) document error, %s",
			r.KeyName, err)
	}

	return r, nil
}

func (s *service) CreateTagValues()
