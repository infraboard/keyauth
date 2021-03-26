package tag

import (
	"fmt"

	"github.com/infraboard/mcube/types/ftime"

	"github.com/infraboard/keyauth/common/tools"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
)

// New 新创建一个Role
func New(tk *token.Token, req *CreateTagRequest) (*TagKey, error) {
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

	r := &TagKey{
		Id:        req.GenUUID(),
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
			Id:       tools.GenHashID(r.Id + req.Values[i].Value),
			CreateAt: ftime.Now().Timestamp(),
			UpdateAt: ftime.Now().Timestamp(),
			Creater:  tk.Account,
			KeyId:    r.Id,
			Value:    req.Values[i],
		})
	}
	return r, nil
}

// NewTagKeySet 实例化make
func NewTagKeySet() *TagKeySet {
	return &TagKeySet{
		Items: []*TagKey{},
	}
}

// Add todo
func (s *TagKeySet) Add(item *TagKey) {
	s.Items = append(s.Items, item)
}

func NewDefaultTagKey() *TagKey {
	return &TagKey{}
}

// NewTagValueSet 实例化make
func NewTagValueSet() *TagValueSet {
	return &TagValueSet{
		Items: []*TagValue{},
	}
}

// Add todo
func (s *TagValueSet) Add(item *TagValue) {
	s.Items = append(s.Items, item)
}

func NewDefaultTagValue() *TagValue {
	return &TagValue{}
}

// NewCreateTagRequest 实例化请求
func NewCreateTagRequest() *CreateTagRequest {
	return &CreateTagRequest{
		Values:         []*ValueOption{},
		HttpFromOption: &HTTPFromOption{},
	}
}
