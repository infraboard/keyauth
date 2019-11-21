package domain

import (
	"fmt"
)

const (
	// Personal 个人域
	Personal Type = iota
	// Enterprise 企业域
	Enterprise
	// Paterner 合作伙伴域
	Paterner
)

// Type 域类型
type Type int

// Domain a tenant container, example an company or organization.
type Domain struct {
	ID       string `json:"id"`                  // 域ID
	CreateAt int64  `json:"create_at,omitempty"` // 创建时间
	UpdateAt int64  `json:"update_at,omitempty"` // 更新时间

	Type           Type   `json:"type,omitempty"`            // 域类型: Personal: 个人, Enterprise: 企业, Paterner: 合作伙伴伙伴
	Name           string `json:"name,omitempty"`            // 公司或者组织名称
	DisplayName    string `json:"display_name,omitempty"`    // 全称
	LogoPath       string `json:"logo_path,omitempty"`       // 公司LOGO图片的URL
	Description    string `json:"description,omitempty"`     // 描述
	Enabled        bool   `json:"enabled,omitempty"`         // 域状态, 是否需要冻结该域, 冻结时, 该域下面所有用户禁止登录
	Size           string `json:"size,omitempty"`            // 规模: 50人以下, 50~100, ...
	Location       string `json:"location,omitempty"`        // 位置: 指城市, 比如 中国,四川,成都
	Address        string `json:"address,omitempty"`         // 地址: 比如环球中心 10F 1034
	Industry       string `json:"industry,omitempty"`        // 所属行业: 比如, 互联网
	Fax            string `json:"fax,omitempty"`             // 传真:
	Phone          string `json:"phone,omitempty"`           // 固话:
	ContactsName   string `json:"contacts_name,omitempty"`   // 联系人姓名
	ContactsTitle  string `json:"contacts_title,omitempty"`  // 联系人职位
	ContactsMobile string `json:"contacts_mobile,omitempty"` // 联系人电话
	ContactsEmail  string `json:"contacts_email,omitempty"`  // 联系人邮箱
}

func (d *Domain) String() string {
	return fmt.Sprint(*d)
}
