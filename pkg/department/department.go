package department

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// Department user's department
type Department struct {
	ID       string `bson:"_id" json:"id"`                        // 部门ID
	Number   string `bson:"number" json:"number,omitempty"`       // 部门编号
	Path     string `bson:"path" json:"path,omitempty"`           // 部门访问路径
	CreateAt int64  `bson:"create_at" json:"create_at,omitempty"` // 部门创建时间
	DomainID string `bson:"domain_id" json:"domain_id,omitempty"` // 部门所属域
	Grade    uint   `bson:"grade" json:"grade,omitempty"`         // 第几级部门, 由层数决定
	*CreateDepartmentRequest
}

// CreateDepartmentRequest 创建部门请求
type CreateDepartmentRequest struct {
	Name      string `json:"name,omitempty" validate:"required,lte=60"`        // 部门名称
	ParentID  string `json:"parent_id,omitempty" validate:"required,lte=200"`  // 上级部门ID
	ManagerID string `json:"manager_id,omitempty" validate:"required,lte=200"` // 部门管理者ID
}

// Validate 校验参数的合法性
func (req *CreateDepartmentRequest) Validate() error {
	return validate.Struct(req)
}

// NewDepartmentSet 实例化
func NewDepartmentSet(req *request.PageRequest) *Set {
	return &Set{
		PageRequest: req,
	}
}

// Set 集合
type Set struct {
	*request.PageRequest

	Total int64         `json:"total"`
	Items []*Department `json:"items"`
}

// Add 添加应用
func (s *Set) Add(item *Department) {
	s.Items = append(s.Items, item)
}
