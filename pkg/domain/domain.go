package domain

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

const (
	// Personal 个人域
	Personal Type = iota
	// Enterprise 企业域
	Enterprise
	// Paterner 合作伙伴域
	Paterner
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// Type 域类型
type Type int

// Domain a tenant container, example an company or organization.
type Domain struct {
	ID       string `bson:"_id" json:"id"`                        // 域ID
	CreateAt int64  `bson:"create_at" json:"create_at,omitempty"` // 创建时间
	UpdateAt int64  `bson:"update_at" json:"update_at,omitempty"` // 更新时间

	Type           Type   `bson:"type" json:"type,omitempty"`                                         // 域类型: Personal: 个人, Enterprise: 企业, Paterner: 合作伙伴伙伴
	Name           string `bson:"name" json:"name,omitempty" validate:"required,gte=1,lte=30"`        // 公司或者组织名称
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

func (d *Domain) String() string {
	return fmt.Sprint(*d)
}

// Validate 校验请求是否合法
func (d *Domain) Validate() error {
	return validate.Struct(d)
}
