package mongo

import (
	"context"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/mconf"
	"github.com/infraboard/mcube/exception"
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
	return nil, nil
}

func (s *service) DeleteGroup(context.Context, *mconf.DeleteGroupRequest) (
	*mconf.Group, error) {
	return nil, nil
}

func (s *service) QueryItem(context.Context, *mconf.QueryItemRequest) (
	*mconf.ItemSet, error) {
	return nil, nil
}

func (s *service) AddItemToGroup(context.Context, *mconf.AddItemToGroupRequest) (
	*mconf.ItemSet, error) {
	return nil, nil
}

func (s *service) RemoveItemFromGroup(context.Context, *mconf.RemoveItemFromGroupRequest) (
	*mconf.ItemSet, error) {
	return nil, nil
}
