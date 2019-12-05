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
func NewPageRequest(page *request.PageRequest) *Request {
	return &Request{
		PageRequest: page,
	}
}

// Request 请求
type Request struct {
	*request.PageRequest

	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Reader for read data from store
type Reader interface {
	GetDomainByID(domainID string) (*Domain, error)
	ListDomain(req *Request) (domains []*Domain, totalPage int64, err error)
}

// Writer for write data to store
type Writer interface {
	CreateDomain(req *CreateDomainRequst) (*Domain, error)
	UpdateDomain(domain *Domain) error
	DeleteDomainByID(id string) error
}
