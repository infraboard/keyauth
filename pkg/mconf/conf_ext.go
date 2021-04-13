package mconf

import (
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
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
