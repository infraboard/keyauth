package namespace

import (
	"github.com/infraboard/mcube/types/ftime"
)

// Namespace tenant resource container
type Namespace struct {
	ID                      string     `bson:"_id" json:"id,omitempty"`                // 项目唯一ID
	DomainID                string     `bson:"domain_id" json:"domain_id,omitempty"`   // 所属域ID
	CreaterID               string     `bson:"creater_id" json:"creater_id,omitempty"` // 创建人
	CreateAt                ftime.Time `bson:"create_at" json:"create_at,omitempty"`   // 创建时间
	UpdateAt                ftime.Time `bson:"update_at" json:"update_at,omitempty"`   // 项目修改时间
	Type                    Type       `bson:"type" json:"type,omitempty"`             // 项目类型
	*CreateNamespaceRequest `bson:",inline"`
}

// CreateNamespaceRequest 创建项目请求
type CreateNamespaceRequest struct {
	Name        string `bson:"name" json:"name,omitempty"`               // 项目名称
	Picture     string `bson:"picture" json:"picture,omitempty"`         // 项目描述图片
	Enabled     bool   `bson:"enabled" json:"enabled,omitempty"`         // 禁用项目, 该项目所有人暂时都无法访问
	OwnerID     string `bson:"owner_id" json:"owner_id,omitempty"`       // 项目所有者, PMO
	Description string `bson:"description" json:"description,omitempty"` // 项目描述
}
