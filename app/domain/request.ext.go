package domain

import (
	"encoding/json"
	"fmt"

	"github.com/infraboard/mcube/http/request"

	"github.com/infraboard/keyauth/common/types"
)

// NewQueryDomainRequest 查询domian列表
func NewQueryDomainRequest(page *request.PageRequest) *QueryDomainRequest {
	return &QueryDomainRequest{
		Page: &page.PageRequest,
	}
}

// Validate 校验请求合法
func (req *QueryDomainRequest) Validate() error {
	return nil
}

// NewDescribeDomainRequest 查询详情请求
func NewDescribeDomainRequest() *DescribeDomainRequest {
	return &DescribeDomainRequest{}
}

// NewDescribeDomainRequestWithName 查询详情请求
func NewDescribeDomainRequestWithName(name string) *DescribeDomainRequest {
	return &DescribeDomainRequest{
		Name: name,
	}
}

// Validate todo
func (req *DescribeDomainRequest) Validate() error {
	if req.Name == "" {
		return fmt.Errorf("name required")
	}

	return nil
}

// NewCreateDomainRequest todo
func NewCreateDomainRequest() *CreateDomainRequest {
	return &CreateDomainRequest{
		Profile: &DomainProfile{},
	}
}

// Validate 校验请求是否合法
func (req *CreateDomainRequest) Validate() error {
	return validate.Struct(req)
}

// Patch todo
func (req *DomainProfile) Patch(data *DomainProfile) {
	patchData, _ := json.Marshal(data)
	json.Unmarshal(patchData, req)
}

// NewPutDomainRequest todo
func NewPutDomainRequest() *UpdateDomainInfoRequest {
	return &UpdateDomainInfoRequest{
		UpdateMode: types.UpdateMode_PUT,
	}
}

// NewPatchDomainRequest todo
func NewPatchDomainRequest() *UpdateDomainInfoRequest {
	return &UpdateDomainInfoRequest{
		UpdateMode: types.UpdateMode_PATCH,
	}
}

// NewPutDomainSecurityRequest todo
func NewPutDomainSecurityRequest() *UpdateDomainSecurityRequest {
	return &UpdateDomainSecurityRequest{
		UpdateMode:      types.UpdateMode_PUT,
		SecuritySetting: NewDefaultSecuritySetting(),
	}
}

// NewDeleteDomainRequestByName todo
func NewDeleteDomainRequestByName(name string) *DeleteDomainRequest {
	return &DeleteDomainRequest{
		Name: name,
	}
}

// Validate todo
func (req *UpdateDomainSecurityRequest) Validate() error {
	return validate.Struct(req)
}

// Validate 更新校验
func (req *UpdateDomainInfoRequest) Validate() error {
	return validate.Struct(req)
}
