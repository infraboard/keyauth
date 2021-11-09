package tag

import (
	"github.com/infraboard/mcube/types/ftime"

	"github.com/infraboard/keyauth/common/tools"
)

// New 新创建一个Role
func New(req *CreateTagRequest) (*TagKey, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	r := &TagKey{
		Id:        req.GenUUID(),
		CreateAt:  ftime.Now().Timestamp(),
		UpdateAt:  ftime.Now().Timestamp(),
		Domain:    req.Domain,
		Creater:   req.Creater,
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
			Creater:  req.Creater,
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
