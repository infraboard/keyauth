package domain

// Service is an domain service
type Service interface {
	Reader
	Writer
}

// NewRequest 请求
func NewRequest() *Request {
	return &Request{
		PageSize:   20,
		PageNumber: 1,
	}
}

// Request 请求
type Request struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	PageSize   uint   `json:"page_size,omitempty" validate:"gte=1,lte=200"`
	PageNumber uint   `json:"page_number,omitempty" validate:"gte=1"`
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
