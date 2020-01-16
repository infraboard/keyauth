package domain

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// New 新建一个domain
func New(ownerID string, req *CreateDomainRequst) (*Domain, error) {
	if ownerID == "" {
		return nil, errors.New("domain required owner")
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &Domain{
		ID:                 xid.New().String(),
		CreateAt:           ftime.Now(),
		UpdateAt:           ftime.Now(),
		OwnerID:            ownerID,
		CreateDomainRequst: req,
	}, nil
}

// NewDomainSet 实例
func NewDomainSet(req *request.PageRequest) *DomainSet {
	return &DomainSet{
		PageRequest: req,
	}
}

// DomainSet domain 列表
type DomainSet struct {
	*request.PageRequest

	Total int64     `json:"total"`
	Items []*Domain `json:"items"`
}

// Add 添加Item
func (s *DomainSet) Add(d *Domain) {
	s.Items = append(s.Items, d)
}

// Domain a tenant container, example an company or organization.
type Domain struct {
	ID                  string     `bson:"_id" json:"id"`                        // 域ID
	CreateAt            ftime.Time `bson:"create_at" json:"create_at,omitempty"` // 创建时间
	UpdateAt            ftime.Time `bson:"update_at" json:"update_at,omitempty"` // 更新时间
	OwnerID             string     `bson:"owner_id" json:"owner_id,omitempty"`   // 域拥有者
	*CreateDomainRequst `bson:",inline"`
}

func (d *Domain) String() string {
	return fmt.Sprint(*d)
}
