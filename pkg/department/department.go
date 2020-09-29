package department

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"

	"github.com/infraboard/keyauth/common/types"
	"github.com/infraboard/keyauth/pkg/counter"
	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/keyauth/pkg/token"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewDepartment 新建实例
func NewDepartment(req *CreateDepartmentRequest, d Service, r role.Service, counter counter.Service) (*Department, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	tk := req.GetToken()

	ins := &Department{
		CreateAt:                ftime.Now(),
		UpdateAt:                ftime.Now(),
		Creater:                 tk.Account,
		Domain:                  tk.Domain,
		Grade:                   1,
		CreateDepartmentRequest: req,
	}

	if req.ParentID != "" {
		pd, err := d.DescribeDepartment(NewDescriptDepartmentRequestWithID(req.ParentID))
		if err != nil {
			return nil, err
		}
		ins.ParentPath = pd.Path()
		ins.Grade = len(strings.Split(pd.Path(), "."))
	}

	if req.Manager == "" {
		req.Manager = tk.Account
	}

	var err error
	// 检查Role是否存在
	if req.DefaultRoleID != "" {
		ins.DefaultRole, err = r.DescribeRole(role.NewDescribeRoleRequestWithID(req.DefaultRoleID))
		if err != nil {
			return nil, err
		}
	}

	// 默认补充访客角色
	if req.DefaultRoleID == "" {
		ins.DefaultRole, err = r.DescribeRole(role.NewDescribeRoleRequestWithName(role.VisitorRoleName))
		if err != nil {
			return nil, err
		}
		ins.DefaultRoleID = ins.DefaultRole.ID
	}

	// 计算ID
	count, err := counter.GetNextSequenceValue(ins.CounterKey())
	if err != nil {
		return nil, err
	}
	ins.Number = count.Value
	ins.ID = fmt.Sprintf("%s.%d", ins.ParentPath, ins.Number)

	return ins, nil
}

// NewDefaultDepartment todo
func NewDefaultDepartment() *Department {
	return &Department{
		CreateDepartmentRequest: NewCreateDepartmentRequest(),
	}
}

// Department user's department
type Department struct {
	ID                       string     `bson:"_id" json:"id"`                        // 部门ID
	ParentPath               string     `bson:"parent_path" json:"parent_path"`       // 路径
	Number                   uint64     `bson:"number" json:"number,omitempty"`       // 部门编号
	CreateAt                 ftime.Time `bson:"create_at" json:"create_at,omitempty"` // 部门创建时间
	UpdateAt                 ftime.Time `bson:"update_at" json:"update_at,omitempty"` // 更新时间
	Creater                  string     `bson:"creater" json:"creater,omitempty"`     // 创建人
	Domain                   string     `bson:"domain" json:"domain,omitempty"`       // 部门所属域
	Grade                    int        `bson:"grade" json:"grade,omitempty"`         // 第几级部门, 由层数决定
	SubCount                 *int64     `bson:"-" json:"sub_count,omitempty"`         // 子部门数量
	UserCount                *int64     `bson:"-" json:"user_count,omitempty"`        // 部门所有用户数量
	*CreateDepartmentRequest `bson:",inline"`

	DefaultRole *role.Role `bson:"-" json:"default_role,omitempty"` // 默认角色
}

func (d *Department) String() string {
	return d.Name
}

// HasSubDepartment todo
func (d *Department) HasSubDepartment() bool {
	return *d.SubCount > 0
}

// CounterKey 编号计算的key
func (d *Department) CounterKey() string {
	return fmt.Sprintf("%s.depart.%d", d.Domain, d.Grade)
}

// Path 具体路径
func (d *Department) Path() string {
	return fmt.Sprintf("%s.%d", d.ParentPath, d.Number)
}

// NewCreateDepartmentRequest todo
func NewCreateDepartmentRequest() *CreateDepartmentRequest {
	return &CreateDepartmentRequest{
		Session: token.NewSession(),
	}
}

// CreateDepartmentRequest 创建部门请求
type CreateDepartmentRequest struct {
	*token.Session `bson:"-" json:"-"`
	Name           string `bson:"name" json:"name" validate:"required,lte=60"`               // 部门名称
	DisplayName    string `bson:"display_name" json:"display_name"`                          // 显示名称
	ParentID       string `bson:"parent_id" json:"parent_id" validate:"lte=200"`             // 上级部门ID
	Manager        string `bson:"manager" json:"manager" validate:"required,lte=200"`        // 部门管理者ID
	DefaultRoleID  string `bson:"default_role_id" json:"default_role_id" validate:"lte=200"` // 部门成员默认角色
}

// Validate 校验参数的合法性
func (req *CreateDepartmentRequest) Validate() error {

	if req.Session == nil {
		return fmt.Errorf("session required")
	}

	tk := req.GetToken()
	if tk == nil {
		return fmt.Errorf("session token required")
	}
	if tk.Domain == "" {
		return fmt.Errorf("user must create domain first")
	}

	return validate.Struct(req)
}

// Patch todo
func (req *CreateDepartmentRequest) Patch(data *CreateDepartmentRequest) {
	patchData, _ := json.Marshal(data)
	json.Unmarshal(patchData, req)
}

// NewDepartmentSet 实例化
func NewDepartmentSet(req *request.PageRequest) *Set {
	return &Set{
		PageRequest: req,
		Items:       []*Department{},
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

// NewPutUpdateDepartmentRequest todo
func NewPutUpdateDepartmentRequest(id string) *UpdateDepartmentRequest {
	return &UpdateDepartmentRequest{
		ID:                      id,
		UpdateMode:              types.PutUpdateMode,
		CreateDepartmentRequest: NewCreateDepartmentRequest(),
	}
}

// NewPatchUpdateDepartmentRequest todo
func NewPatchUpdateDepartmentRequest(id string) *UpdateDepartmentRequest {
	return &UpdateDepartmentRequest{
		ID:                      id,
		UpdateMode:              types.PatchUpdateMode,
		CreateDepartmentRequest: NewCreateDepartmentRequest(),
	}
}

// UpdateDepartmentRequest todo
type UpdateDepartmentRequest struct {
	ID         string           `json:"id"`
	UpdateMode types.UpdateMode `json:"update_mode"`
	*CreateDepartmentRequest
}

// Validate 校验入参
func (req *UpdateDepartmentRequest) Validate() error {
	if req.ID == "" {
		return fmt.Errorf("department id requred")
	}

	return req.CreateDepartmentRequest.Validate()
}
