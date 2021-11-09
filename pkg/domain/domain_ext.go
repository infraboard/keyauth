package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/types/ftime"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// New 新建一个domain
func New(req *CreateDomainRequest) (*Domain, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	d := &Domain{
		CreateAt:        ftime.Now().Timestamp(),
		UpdateAt:        ftime.Now().Timestamp(),
		Owner:           req.Owner,
		Name:            req.Name,
		Profile:         req.Profile,
		Enabled:         true,
		SecuritySetting: NewDefaultSecuritySetting(),
	}

	return d, nil
}

// NewDefault todo
func NewDefault() *Domain {
	return &Domain{
		Profile:         &DomainProfile{},
		SecuritySetting: NewDefaultSecuritySetting(),
	}
}

// NewDomainSet 实例
func NewDomainSet() *Set {
	return &Set{
		Items: []*Domain{},
	}
}

// Length 总个数
func (s *Set) Length() int {
	return len(s.Items)
}

// Add 添加Item
func (s *Set) Add(d *Domain) {
	s.Items = append(s.Items, d)
}
