package namespace

import (
	"errors"
	"fmt"
	"net/http"

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

// NewQueryNamespaceRequestFromHTTP 列表查询请求
func NewQueryNamespaceRequestFromHTTP(r *http.Request) *QueryNamespaceRequest {
	qs := r.URL.Query()
	return &QueryNamespaceRequest{
		Session:     token.NewSession(),
		PageRequest: request.NewPageRequestFromHTTP(r),
		Department:  qs.Get("department"),
	}
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
	Department string
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

// NewDeleteNamespaceRequestWithID todo
func NewDeleteNamespaceRequestWithID(id string) *DeleteNamespaceRequest {
	return &DeleteNamespaceRequest{
		Session: token.NewSession(),
		ID:      id,
	}
}

// DeleteNamespaceRequest todo
type DeleteNamespaceRequest struct {
	*token.Session
	ID string
}

// Validate todo
func (req *DeleteNamespaceRequest) Validate() error {
	tk := req.GetToken()
	if tk == nil {
		return fmt.Errorf("token required")
	}

	if req.ID == "" {
		return fmt.Errorf("id required")
	}

	return nil
}
