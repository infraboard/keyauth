package mongo

import (
	"context"

	"github.com/infraboard/keyauth/pkg/mconf"
)

func (s *service) CreateGroup(ctx context.Context, req *mconf.CreateGroupRequest) (
	*mconf.Group, error) {
	return nil, nil
}

func (s *service) QueryGroup(ctx context.Context, req *mconf.QueryGroupRequest) (
	*mconf.GroupSet, error) {
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
