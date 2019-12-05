package domain

import (
	"github.com/infraboard/mcube/http/request"
)

// Service is an domain service
type Service interface {
	Reader
	Writer
}

// NewPageRequest 分页请求
func NewPageRequest(page *request.PageRequest) *QueryDomainRequest {
	return &QueryDomainRequest{
		PageRequest: page,
	}
}

// QueryDomainRequest 请求
type QueryDomainRequest struct {
	*request.PageRequest

	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Reader for read data from store
type Reader interface {
	GetDomainByID(domainID string) (*Domain, error)
	ListDomain(req *QueryDomainRequest) (domains []*Domain, totalPage int64, err error)
}

// Writer for write data to store
type Writer interface {
	CreateDomain(req *CreateDomainRequst) (*Domain, error)
	UpdateDomain(*Domain) error
	DeleteDomainByID(id string) error
}
