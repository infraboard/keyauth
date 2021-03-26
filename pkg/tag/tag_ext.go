package tag

import (
	"fmt"

	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"

	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
)

// New 新创建一个Role
func New(tk *token.Token, req *CreateTagRequest) (*Tag, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	switch req.ScopeType {
	case ScopeType_GLOBAL:
		if !tk.UserType.IsIn(types.UserType_SUPPER) {
			return nil, fmt.Errorf("only supper account can create global tag")
		}
	case ScopeType_DOMAIN:
		if !tk.UserType.IsIn(types.UserType_SUPPER, types.UserType_DOMAIN_ADMIN) {
			return nil, fmt.Errorf("only domain account can create domain tag")
		}
	}

	r := &Tag{
		Id:        xid.New().String(),
		CreateAt:  ftime.Now().Timestamp(),
		UpdateAt:  ftime.Now().Timestamp(),
		Domain:    tk.Domain,
		Creater:   tk.Account,
		ScopeType: req.ScopeType,
		Namespace: req.Namespace,
		KeyName:   req.KeyName,
		KeyLabel:  req.KeyLabel,
		KeyDesc:   req.KeyDesc,
	}

	for i := range req.Values {
		r.Values = append(r.Values, &TagValue{
			Id:       xid.New().String(),
			CreateAt: ftime.Now().Timestamp(),
			UpdateAt: ftime.Now().Timestamp(),
			Creater:  tk.Account,
			KeyId:    r.Id,
			Value:    req.Values[i],
		})
	}
	return r, nil
}

// NewTagSet 实例化make
func NewTagSet() *TagSet {
	return &TagSet{
		Items: []*Tag{},
	}
}

// Add todo
func (s *TagSet) Add(item *Tag) {
	s.Items = append(s.Items, item)
}

func NewDefaultTag() *Tag {
	return &Tag{}
}
