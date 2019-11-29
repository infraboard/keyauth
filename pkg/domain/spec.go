package domain

// Service is an domain service
type Service interface {
	Reader
	Writer
}

// Request 请求
type Request struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	PageNumber int64  `json:"page_number,omitempty"`
	PageSize   int64  `json:"page_size,omitempty"`
}

// Reader for read data from store
type Reader interface {
	GetDomainByID(domainID string) (*Domain, error)
	ListDomain(req *Request) (domains []*Domain, totalPage int64, err error)
}

// Writer for write data to store
type Writer interface {
	CreateDomain(domain *Domain) error
	UpdateDomain(domain *Domain) error
	DeleteDomainByID(id string) error
}
