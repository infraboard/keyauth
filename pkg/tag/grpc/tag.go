package grpc

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/tag"
	"github.com/infraboard/keyauth/pkg/user/types"
)

func (s *service) CreateTag(ctx context.Context, req *tag.CreateTagRequest) (*tag.TagKey, error) {
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

func (s *service) QueryTagKey(ctx context.Context, req *tag.QueryTagKeyRequest) (*tag.TagKeySet, error) {
	query, err := newQueryTagKeyRequest(req)
	if err != nil {
		return nil, err
	}

	resp, err := s.key.Find(context.TODO(), query.FindFilter(), query.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find role error, error is %s", err)
	}

	set := tag.NewTagKeySet()
	// 循环
	for resp.Next(context.TODO()) {
		ins := tag.NewDefaultTagKey()
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

func (s *service) QueryTagValue(ctx context.Context, req *tag.QueryTagValueRequest) (*tag.TagValueSet, error) {
	query, err := newQueryTagValueRequest(req)
	if err != nil {
		return nil, err
	}

	resp, err := s.value.Find(context.TODO(), query.FindFilter(), query.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find role error, error is %s", err)
	}

	set := tag.NewTagValueSet()
	// 循环
	for resp.Next(context.TODO()) {
		ins := tag.NewDefaultTagValue()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode role error, error is %s", err)
		}
		set.Add(ins)
	}

	// count
	count, err := s.value.CountDocuments(context.TODO(), query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get tag key count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

func (s *service) DescribeTag(ctx context.Context, req *tag.DescribeTagRequest) (*tag.TagKey, error) {
	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	query, err := newDescribeTagRequest(tk, req)
	if err != nil {
		return nil, err
	}

	ins := tag.NewDefaultTagKey()
	if err := s.key.FindOne(context.TODO(), query.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("tag %s not found", req)
		}

		return nil, exception.NewInternalServerError("find tag %s error, %s", req, err)
	}

	return ins, nil
}

func (s *service) DeleteTag(ctx context.Context, req *tag.DeleteTagRequest) (*tag.TagKey, error) {
	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	ins, err := s.DescribeTag(ctx, tag.NewDescribeTagRequestWithID(req.TagId))
	if err != nil {
		return nil, err
	}

	switch ins.ScopeType {
	case tag.ScopeType_GLOBAL:
		if !tk.UserType.IsIn(types.UserType_SUPPER) {
			return nil, exception.NewBadRequest("only supper can delete global tag")
		}
	case tag.ScopeType_DOMAIN:
		if !tk.UserType.IsIn(types.UserType_SUPPER, types.UserType_DOMAIN_ADMIN) {
			return nil, exception.NewBadRequest("only domain admin can delete domain tag")
		}
	}

	_, err = s.key.DeleteOne(context.TODO(), bson.M{"_id": req.TagId})
	if err != nil {
		return nil, exception.NewInternalServerError("delete tag(%s) error, %s", req.TagId, err)
	}

	// 清除tag 关联的值
	_, err = s.value.DeleteMany(context.TODO(), bson.M{"key_id": req.TagId})
	if err != nil {
		s.log.Errorf("delete tag values error, %s", err)
	}

	return ins, nil
}
