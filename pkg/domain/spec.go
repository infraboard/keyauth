package domain

import (
	"github.com/infraboard/mcube/http/request"
)

// Service is an domain service
type Service interface {
	DescriptionDomain(req *DescriptDomainRequest) (*Domain, error)
	QueryDomain(req *QueryDomainRequest) (domains []*Domain, totalPage int64, err error)
	CreateDomain(req *CreateDomainRequst) (*Domain, error)
	UpdateDomain(*Domain) error
	DeleteDomain(id string) error
}

// NewQueryDomainRequest 查询domian列表
func NewQueryDomainRequest(page *request.PageRequest) *QueryDomainRequest {
	return &QueryDomainRequest{
		PageRequest:           page,
		DescriptDomainRequest: &DescriptDomainRequest{},
	}
}

// QueryDomainRequest 请求
type QueryDomainRequest struct {
	*request.PageRequest
	*DescriptDomainRequest
}

// NewDescriptDomainRequest 查询详情请求
func NewDescriptDomainRequest() *DescriptDomainRequest {
	return &DescriptDomainRequest{}
}

// DescriptDomainRequest 查询domain详情请求
type DescriptDomainRequest struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
