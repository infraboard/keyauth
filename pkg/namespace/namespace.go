package namespace

import (
	"github.com/go-playground/validator/v10"

	"github.com/infraboard/keyauth/pkg/token"
)

// use a single instance of Validate, it caches struct info
var (
	validater = validator.New()
)

// NewDefaultNamespace todo
func NewDefaultNamespace() *Namespace {
	return &Namespace{
		Enabled: true,
	}
}

// NewCreateNamespaceRequest todo
func NewCreateNamespaceRequest() *CreateNamespaceRequest {
	return &CreateNamespaceRequest{}
}

func (req *CreateNamespaceRequest) UpdateOwner(tk *token.Token) {
	req.CreateBy = tk.Account
	req.Domain = tk.Domain
}

// Validate todo
func (req *CreateNamespaceRequest) Validate() error {
	return validater.Struct(req)
}

// NewNamespaceSet 实例化
func NewNamespaceSet() *Set {
	return &Set{
		Items: []*Namespace{},
	}
}

// Add 添加应用
func (s *Set) Add(item *Namespace) {
	s.Items = append(s.Items, item)
}
