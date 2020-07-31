package domain

import (
	"fmt"

	"github.com/infraboard/mcube/http/request"

	"github.com/infraboard/keyauth/pkg/token"
)

// Service is an domain service
type Service interface {
	UpdateDomain(*Domain) error
	CreateDomain(ownerID string, req *CreateDomainRequst) (*Domain, error)
	DescriptionDomain(req *DescriptDomainRequest) (*Domain, error)
	QueryDomain(req *QueryDomainRequest) (*Set, error)
	DeleteDomain(id string) error
}

// NewQueryDomainRequest 查询domian列表
func NewQueryDomainRequest(page *request.PageRequest) *QueryDomainRequest {
	return &QueryDomainRequest{
		PageRequest:           page,
		Session:               token.NewSession(),
		DescriptDomainRequest: NewDescriptDomainRequest(),
	}
}

// QueryDomainRequest 请求
type QueryDomainRequest struct {
	*token.Session
	*request.PageRequest
	*DescriptDomainRequest
}

// Validate 校验请求合法
func (req *QueryDomainRequest) Validate() error {
	return nil
}

// NewDescriptDomainRequest 查询详情请求
func NewDescriptDomainRequest() *DescriptDomainRequest {
	return &DescriptDomainRequest{}
}

// DescriptDomainRequest 查询domain详情请求
type DescriptDomainRequest struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Validate todo
func (req *DescriptDomainRequest) Validate() error {
	if req.ID == "" && req.Name == "" {
		return fmt.Errorf("id, name or base_dn required")
	}

	return nil
}

// NewCreateDomainRequst todo
func NewCreateDomainRequst() *CreateDomainRequst {
	return &CreateDomainRequst{
		Enabled: true,
	}
}

// CreateDomainRequst 创建请求
type CreateDomainRequst struct {
	Name           string `bson:"_id" json:"name" validate:"required,lte=60"`                         // 公司或者组织名称
	DisplayName    string `bson:"display_name" json:"display_name,omitempty" validate:"lte=80"`       // 全称
	LogoPath       string `bson:"logo_path" json:"logo_path,omitempty" validate:"lte=200"`            // 公司LOGO图片的URL
	Description    string `bson:"description" json:"description,omitempty" validate:"lte=400"`        // 描述
	Enabled        bool   `bson:"enabled" json:"enabled,omitempty"`                                   // 域状态, 是否需要冻结该域, 冻结时, 该域下面所有用户禁止登录
	Size           string `bson:"size" json:"size,omitempty" validate:"lte=20"`                       // 规模: 50人以下, 50~100, ...
	Location       string `bson:"location" json:"location,omitempty" validate:"lte=20"`               // 位置: 指城市, 比如 中国,四川,成都
	Address        string `bson:"address" json:"address,omitempty" validate:"lte=200"`                // 地址: 比如环球中心 10F 1034
	Industry       string `bson:"industry" json:"industry,omitempty" validate:"lte=100"`              // 所属行业: 比如, 互联网
	Fax            string `bson:"fax" json:"fax,omitempty" validate:"lte=40"`                         // 传真:
	Phone          string `bson:"phone" json:"phone,omitempty" validate:"lte=20"`                     // 固话:
	ContactsName   string `bson:"contacts_name" json:"contacts_name,omitempty" validate:"lte=30"`     // 联系人姓名
	ContactsTitle  string `bson:"contacts_title" json:"contacts_title,omitempty" validate:"lte=40"`   // 联系人职位
	ContactsMobile string `bson:"contacts_mobile" json:"contacts_mobile,omitempty" validate:"lte=20"` // 联系人电话
	ContactsEmail  string `bson:"contacts_email" json:"contacts_email,omitempty" validate:"lte=40"`   // 联系人邮箱
}

// Validate 校验请求是否合法
func (req *CreateDomainRequst) Validate() error {
	return validate.Struct(req)
}

func (req *CreateDomainRequst) String() string {
	return fmt.Sprint(*req)
}
