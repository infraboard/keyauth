package endpoint

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/types/ftime"

	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// Service token管理服务
type Service interface {
	DescribeEndpoint(req *DescribeEndpointRequest) (*Endpoint, error)
	QueryEndpoints(req *QueryEndpointRequest) (*Set, error)
	Registry(req *RegistryRequest) error
}

// NewRegistryRequest 注册请求
func NewRegistryRequest(version string, entries []*router.Entry) *RegistryRequest {
	return &RegistryRequest{
		Version: version,
		Entries: entries,
		Session: token.NewSession(),
	}
}

// NewDefaultRegistryRequest todo
func NewDefaultRegistryRequest() *RegistryRequest {
	return &RegistryRequest{
		Session: token.NewSession(),
		Entries: []*router.Entry{},
	}
}

// RegistryRequest 服务注册请求
type RegistryRequest struct {
	*token.Session
	Version string          `json:"version" validate:"required,lte=32"`
	Entries []*router.Entry `json:"entries"`
}

// Validate 校验注册请求合法性
func (req *RegistryRequest) Validate() error {
	if len(req.Entries) == 0 {
		return fmt.Errorf("must require *router.Entry")
	}

	tk := req.GetToken()
	if tk == nil {
		return fmt.Errorf("token required when service endpoints registry")
	}

	if !tk.UserType.Is(types.ServiceAccount) {
		return fmt.Errorf("only service account can registry endpoints")
	}

	return validate.Struct(req)
}

// Endpoints 功能列表
func (req *RegistryRequest) Endpoints(svr string) []*Endpoint {
	eps := make([]*Endpoint, 0, len(req.Entries))
	for i := range req.Entries {
		ep := &Endpoint{
			ID:       GenHashID(svr, req.Entries[i].Path, req.Entries[i].Method),
			CreateAt: ftime.Now(),
			UpdateAt: ftime.Now(),
			Service:  svr,
			Version:  req.Version,
			Entry:    *req.Entries[i],
		}
		eps = append(eps, ep)
	}
	return eps
}

// NewQueryEndpointRequest 列表查询请求
func NewQueryEndpointRequest(pageReq *request.PageRequest) *QueryEndpointRequest {
	return &QueryEndpointRequest{
		PageRequest: pageReq,
	}
}

// QueryEndpointRequest 查询应用列表
type QueryEndpointRequest struct {
	*request.PageRequest
}

// NewDescribeEndpointRequestWithID todo
func NewDescribeEndpointRequestWithID(id string) *DescribeEndpointRequest {
	return &DescribeEndpointRequest{ID: id}
}

// DescribeEndpointRequest todo
type DescribeEndpointRequest struct {
	ID string
}

// Validate 校验
func (req *DescribeEndpointRequest) Validate() error {
	if req.ID == "" {
		return fmt.Errorf("endpoint id is required")
	}

	return nil
}
