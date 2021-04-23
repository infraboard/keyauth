package mconf

import (
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
)

// New 创建服务
func NewGroup(req *CreateGroupRequest) (*Group, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	ins := &Group{
		CreateAt:    ftime.Now().Timestamp(),
		Type:        req.Type,
		Name:        req.Name,
		Description: req.Description,
	}

	return ins, nil
}

// NewMicroSet 实例化
func NewGroupSet() *GroupSet {
	return &GroupSet{
		Items: []*Group{},
	}
}

// Add 添加
func (s *GroupSet) Add(e *Group) {
	s.Items = append(s.Items, e)
}

func NewItem(creater string, req *ItemRequest) *Item {
	return &Item{
		Id:          xid.New().String(),
		Key:         req.Key,
		Group:       req.Group,
		Creater:     creater,
		CreateAt:    ftime.Now().Timestamp(),
		Value:       req.Value,
		Description: req.Description,
	}
}

func NewGroupItemSet(creater string, req *AddItemToGroupRequest) *ItemSet {
	set := NewItemSet()
	for i := range req.Items {
		item := NewItem(creater, req.Items[i])
		item.Group = req.GroupName
		set.Add(item)
	}
	return set
}

// NewMicroSet 实例化
func NewItemSet() *ItemSet {
	return &ItemSet{
		Items: []*Item{},
	}
}

// Add 添加
func (s *ItemSet) Add(e *Item) {
	s.Items = append(s.Items, e)
}

func (s *ItemSet) Docs() []interface{} {
	docs := []interface{}{}
	for i := range s.Items {
		docs = append(docs, s.Items[i])
	}
	return docs
}
