package namespace

import (
	"errors"

	"github.com/infraboard/mcube/http/request"

	"github.com/infraboard/keyauth/pkg/token"
)

// Service todo
type Service interface {
	CreateNamespace(req *CreateNamespaceRequest) (*Namespace, error)
	QueryNamespace(req *QueryNamespaceRequest) (*Set, error)
	DescribeNamespace(req *DescriptNamespaceRequest) (*Namespace, error)
	DeleteNamespace(req *DeleteNamespaceRequest) error
}

// NewQueryNamespaceRequest 列表查询请求
func NewQueryNamespaceRequest(pageReq *request.PageRequest) *QueryNamespaceRequest {
	return &QueryNamespaceRequest{
		Session:     token.NewSession(),
		PageRequest: pageReq,
	}
}

// QueryNamespaceRequest 查询应用列表
type QueryNamespaceRequest struct {
	*token.Session
	*request.PageRequest
}

// NewNewDescriptNamespaceRequestWithID todo
func NewNewDescriptNamespaceRequestWithID(id string) *DescriptNamespaceRequest {
	req := NewDescriptNamespaceRequest()
	req.ID = id
	return req
}

// NewDescriptNamespaceRequest new实例
func NewDescriptNamespaceRequest() *DescriptNamespaceRequest {
	return &DescriptNamespaceRequest{
		Session: token.NewSession(),
	}
}

// DescriptNamespaceRequest 查询应用详情
type DescriptNamespaceRequest struct {
	*token.Session
	ID string `json:"id,omitempty"`
}

// Validate 校验详情查询请求
func (req *DescriptNamespaceRequest) Validate() error {
	if req.ID == "" {
		return errors.New("id  is required")
	}

	return nil
}

// DeleteNamespaceRequest todo
type DeleteNamespaceRequest struct {
	ID string
}
