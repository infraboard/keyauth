package endpoint

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	httpb "github.com/infraboard/mcube/pb/http"
	"github.com/infraboard/mcube/types/ftime"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewRegistryRequest 注册请求
func NewRegistryRequest(version string, entries []*httpb.Entry) *RegistryRequest {
	return &RegistryRequest{
		Version: version,
		Entries: entries,
	}
}

// NewDefaultRegistryRequest todo
func NewDefaultRegistryRequest() *RegistryRequest {
	return &RegistryRequest{
		Entries: []*httpb.Entry{},
	}
}

// Validate 校验注册请求合法性
func (req *RegistryRequest) Validate() error {
	if len(req.Entries) == 0 {
		return fmt.Errorf("must require *router.Entry")
	}

	return validate.Struct(req)
}

// Endpoints 功能列表
func (req *RegistryRequest) Endpoints(serviceID string) []*Endpoint {
	eps := make([]*Endpoint, 0, len(req.Entries))
	for i := range req.Entries {
		ep := &Endpoint{
			Id:        GenHashID(serviceID, req.Entries[i].Path),
			CreateAt:  ftime.Now().Timestamp(),
			UpdateAt:  ftime.Now().Timestamp(),
			ServiceId: serviceID,
			Version:   req.Version,
			Entry:     req.Entries[i],
		}
		eps = append(eps, ep)
	}
	return eps
}

// NewRegistryResponse todo
func NewRegistryResponse(message string) *RegistryResponse {
	return &RegistryResponse{Message: message}
}

// NewQueryEndpointRequestFromHTTP 列表查询请求
func NewQueryEndpointRequestFromHTTP(r *http.Request) *QueryEndpointRequest {
	page := request.NewPageRequestFromHTTP(r)
	qs := r.URL.Query()

	query := &QueryEndpointRequest{
		Page:         &page.PageRequest,
		Path:         qs.Get("path"),
		Method:       qs.Get("method"),
		FunctionName: qs.Get("function_name"),
	}

	sids := qs.Get("service_ids")
	if sids != "" {
		query.ServiceIds = strings.Split(sids, ",")
	}
	rs := qs.Get("resources")
	if rs != "" {
		query.Resources = strings.Split(rs, ",")
	}

	return query
}

// NewQueryEndpointRequest 列表查询请求
func NewQueryEndpointRequest(pageReq *request.PageRequest) *QueryEndpointRequest {
	return &QueryEndpointRequest{
		Page: &pageReq.PageRequest,
	}
}

// NewDescribeEndpointRequestWithID todo
func NewDescribeEndpointRequestWithID(id string) *DescribeEndpointRequest {
	return &DescribeEndpointRequest{Id: id}
}

// Validate 校验
func (req *DescribeEndpointRequest) Validate() error {
	if req.Id == "" {
		return fmt.Errorf("endpoint id is required")
	}

	return nil
}

// NewDeleteEndpointRequestWithServiceID todo
func NewDeleteEndpointRequestWithServiceID(id string) *DeleteEndpointRequest {
	return &DeleteEndpointRequest{ServiceId: id}
}

// NewQueryResourceRequestFromHTTP 列表查询请求
func NewQueryResourceRequestFromHTTP(r *http.Request) *QueryResourceRequest {
	page := request.NewPageRequestFromHTTP(r)
	qs := r.URL.Query()
	pe, err := ParseBoolQueryFromString(qs.Get("permission_enable"))
	if err != nil {
		pe = BoolQuery_ALL
	}

	query := &QueryResourceRequest{
		Page:             &page.PageRequest,
		PermissionEnable: pe,
	}

	strIds := qs.Get("service_ids")
	if strIds != "" {
		query.ServiceIds = strings.Split(strIds, ",")
	}
	rs := qs.Get("resources")
	if rs != "" {
		query.Resources = strings.Split(rs, ",")
	}

	return query
}

// Validate todo
func (req *QueryResourceRequest) Validate() error {
	if len(req.ServiceIds) == 0 {
		return exception.NewBadRequest("service_ids required, but \"\"")
	}

	return nil
}
