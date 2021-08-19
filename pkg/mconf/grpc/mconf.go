package grpc

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/mconf"
)

func (s *service) CreateGroup(ctx context.Context, req *mconf.CreateGroupRequest) (
	*mconf.Group, error) {
	ins, err := mconf.NewGroup(req)
	if err != nil {
		return nil, err
	}

	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	ins.Creater = tk.Account

	if _, err := s.group.InsertOne(context.TODO(), ins); err != nil {
		return nil, exception.NewInternalServerError("inserted group document error, %s", err)
	}
	return ins, nil
}

func (s *service) QueryGroup(ctx context.Context, req *mconf.QueryGroupRequest) (
	*mconf.GroupSet, error) {
	r := newGroupPaggingQuery(req)
	resp, err := s.group.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find group error, error is %s", err)
	}

	set := mconf.NewGroupSet()
	// 循环
	for resp.Next(context.TODO()) {
		ins := new(mconf.Group)
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode group error, error is %s", err)
		}

		set.Add(ins)
	}

	// count
	count, err := s.group.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get group count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

func (s *service) DescribeGroup(ctx context.Context, req *mconf.DescribeGroupRequest) (
	*mconf.Group, error) {
	r, err := newDescribeGroupQuery(req)
	if err != nil {
		return nil, err
	}

	ins := new(mconf.Group)
	if err := s.group.FindOne(context.TODO(), r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("group %s not found", req)
		}

		return nil, exception.NewInternalServerError("find group %s error, %s", req, err)
	}
	return ins, nil
}

func (s *service) DeleteGroup(ctx context.Context, req *mconf.DeleteGroupRequest) (
	*mconf.Group, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate delete group error, %s", err)
	}

	descReq := mconf.NewDescribeGroupRequestWithName(req.Name)
	ins, err := s.DescribeGroup(ctx, descReq)
	if err != nil {
		return nil, err
	}

	// 清除服务实体
	_, err = s.group.DeleteOne(context.TODO(), bson.M{"_id": req.Name})
	if err != nil {
		return nil, exception.NewInternalServerError("delete group(%s) error, %s", req.Name, err)
	}

	return ins, nil
}

func (s *service) QueryItem(ctx context.Context, req *mconf.QueryItemRequest) (
	*mconf.ItemSet, error) {
	r := newItemPaggingQuery(req)
	resp, err := s.item.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find group error, error is %s", err)
	}

	set := mconf.NewItemSet()
	// 循环
	for resp.Next(context.TODO()) {
		ins := new(mconf.Item)
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode group error, error is %s", err)
		}

		set.Add(ins)
	}

	// count
	count, err := s.group.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get item count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

func (s *service) AddItemToGroup(ctx context.Context, req *mconf.AddItemToGroupRequest) (
	*mconf.ItemSet, error) {
	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	set := mconf.NewGroupItemSet(tk.Account, req)
	if _, err := s.item.InsertMany(context.TODO(), set.Docs()); err != nil {
		return nil, exception.NewInternalServerError("inserted item document error, %s", err)
	}
	return set, nil
}

func (s *service) RemoveItemFromGroup(ctx context.Context, req *mconf.RemoveItemFromGroupRequest) (
	*mconf.ItemSet, error) {
	// if err := req.Validate(); err != nil {
	// 	return nil, exception.NewBadRequest("validate delete service error, %s", err)
	// }

	// describeReq := mconf.NewDescribeServiceRequest()
	// describeReq.Id = req.Id
	// svr, err := s.DescribeService(ctx, describeReq)
	// if err != nil {
	// 	return nil, err
	// }

	// // 清除服务实体
	// _, err = s.item.DeleteOne(context.TODO(), bson.M{"_id": req.Id})
	// if err != nil {
	// 	return nil, exception.NewInternalServerError("delete service(%s) error, %s", req.Id, err)
	// }

	return nil, nil
}
