package project

// Project tenant resource container
type Project struct {
	ID string `json:"id,omitempty"` // 项目唯一ID
	*CreateProjectRequest
	DomainID string `json:"domain_id,omitempty"` // 所属域ID
	CreateAt int64  `json:"create_at,omitempty"` // 创建时间
	UpdateAt int64  `json:"update_at,omitempty"` // 项目修改时间
}

// CreateProjectRequest 创建项目请求
type CreateProjectRequest struct {
	Name        string `json:"name,omitempty"`        // 项目名称
	Picture     string `json:"picture,omitempty"`     // 项目描述图片
	Type        string `json:"type,omitempty"`        // 项目类型
	Enabled     bool   `json:"enabled,omitempty"`     // 禁用项目, 该项目所有人暂时都无法访问
	Owner       string `json:"owner_id,omitempty"`    // 项目所有者, PMO
	Region      string `json:"region,omitempty"`      // 项目所属地区
	Description string `json:"description,omitempty"` // 项目描述
}
