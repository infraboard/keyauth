package policy

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
)

const (
	MaxUserPolicy = 2048
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewQueryPolicyRequestFromHTTP 列表查询请求
func NewQueryPolicyRequestFromHTTP(r *http.Request) *QueryPolicyRequest {
	page := request.NewPageRequestFromHTTP(r)
	req := NewQueryPolicyRequest(page)

	qs := r.URL.Query()
	req.Account = qs.Get("account")
	req.RoleId = qs.Get("role_id")
	req.NamespaceId = qs.Get("namespace_id")
	req.WithRole = qs.Get("with_role") == "true"
	req.WithNamespace = qs.Get("with_namespace") == "true"
	return req
}

// NewQueryPolicyRequest 列表查询请求
func NewQueryPolicyRequest(pageReq *request.PageRequest) *QueryPolicyRequest {
	return &QueryPolicyRequest{
		Page:          &pageReq.PageRequest,
		WithRole:      false,
		WithNamespace: false,
	}
}

// Validate 校验请求是否合法
func (req *QueryPolicyRequest) Validate() error {
	return validate.Struct(req)
}

func (req *QueryPolicyRequest) CheckOwner(account string) bool {
	return req.Account == account
}

// NewDescriptPolicyRequest new实例
func NewDescriptPolicyRequest() *DescribePolicyRequest {
	return &DescribePolicyRequest{}
}

// Validate todo
func (req *DescribePolicyRequest) Validate() error {
	if req.Id == "" {
		return fmt.Errorf("policy id required")
	}

	return nil
}

// NewDeletePolicyRequestWithID todo
func NewDeletePolicyRequestWithID(id string) *DeletePolicyRequest {
	req := NewDeletePolicyRequest()
	req.Id = id
	return req
}

// NewDeletePolicyRequestWithNamespaceID todo
func NewDeletePolicyRequestWithNamespaceID(namespaceID string) *DeletePolicyRequest {
	req := NewDeletePolicyRequest()
	req.NamespaceId = namespaceID
	return req
}

// NewDeletePolicyRequestWithAccount todo
func NewDeletePolicyRequestWithAccount(account string) *DeletePolicyRequest {
	req := NewDeletePolicyRequest()
	req.Account = account
	return req
}

// NewDeletePolicyRequestWithRoleID todo
func NewDeletePolicyRequestWithRoleID(roleID string) *DeletePolicyRequest {
	req := NewDeletePolicyRequest()
	req.RoleId = roleID
	return req
}

// NewDeletePolicyRequest todo
func NewDeletePolicyRequest() *DeletePolicyRequest {
	return &DeletePolicyRequest{}
}

// Validate todo
func (req *DeletePolicyRequest) Validate() error {
	if req.Id == "" {
		return fmt.Errorf("policy id required")
	}

	return nil
}
