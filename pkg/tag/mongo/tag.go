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

	t, err := tag.New(tk, req)
	if err != nil {
		return nil, err
	}

	// insert key
	if _, err := s.key.InsertOne(context.TODO(), t); err != nil {
		return nil, exception.NewInternalServerError("inserted tag key(%s) document error, %s",
			t.KeyName, err)
	}

	// insert values
	values := []interface{}{}
	for i := range t.Values {
		values = append(values, t.Values[i])
	}
	if s.value.InsertMany(context.TODO(), values); err != nil {
		return nil, exception.NewInternalServerError("inserted tag value(%s) document error, %s",
			t.KeyName, err)
	}

	return t, nil
}

func (s *service) QueryTag(ctx context.Context, req *tag.QueryTagRequest) (*tag.TagSet, error) {
	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	query, err := newQueryTagRequest(tk, req)
	if err != nil {
		return nil, err
	}

	resp, err := s.key.Find(context.TODO(), query.FindFilter(), query.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find role error, error is %s", err)
	}

	set := tag.NewTagSet()
	// 循环
	for resp.Next(context.TODO()) {
		ins := tag.NewDefaultTag()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode role error, error is %s", err)
		}
		set.Add(ins)
	}

	// count
	count, err := s.key.CountDocuments(context.TODO(), query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get tag key count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (s *service) DescribeTag(context.Context, *tag.DescribeTagRequest) (*tag.Tag, error) {
	return nil, nil
}
