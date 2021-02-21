package department

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"

	"github.com/infraboard/keyauth/common/session"
	"github.com/infraboard/keyauth/common/types"
	"github.com/infraboard/keyauth/pkg/counter"
	"github.com/infraboard/keyauth/pkg/role"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewDepartment 新建实例
func NewDepartment(ctx context.Context, req *CreateDepartmentRequest, d DepartmentServiceServer, r role.RoleServiceServer, counter counter.Service) (*Department, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	tk := session.GetTokenFromContext(ctx)
	ins := &Department{
		CreateAt:      ftime.Now().Timestamp(),
		UpdateAt:      ftime.Now().Timestamp(),
		Creater:       tk.Account,
		Domain:        tk.Domain,
		Grade:         1,
		Name:          req.Name,
		DisplayName:   req.DisplayName,
		ParentId:      req.ParentId,
		Manager:       req.Manager,
		DefaultRoleId: req.DefaultRoleId,
	}

	if req.ParentId != "" {
		pd, err := d.DescribeDepartment(ctx, NewDescribeDepartmentRequestWithID(req.ParentId))
		if err != nil {
			return nil, err
		}
		ins.ParentPath = pd.Path()
		ins.Grade = int32(len(strings.Split(pd.Path(), ".")))
	}

	if req.Manager == "" {
		ins.Manager = tk.Account
	}

	var err error
	// 检查Role是否存在
	if req.DefaultRoleId != "" {
		ins.DefaultRole, err = r.DescribeRole(ctx, role.NewDescribeRoleRequestWithID(req.DefaultRoleId))
		if err != nil {
			return nil, err
		}
	} else {
		// 默认补充访客角色
		ins.DefaultRole, err = r.DescribeRole(ctx, role.NewDescribeRoleRequestWithName(role.VisitorRoleName))
		if err != nil {
			return nil, err
		}
		ins.DefaultRoleId = ins.DefaultRole.Id
	}

	// 计算ID
	count, err := counter.GetNextSequenceValue(ins.CounterKey())
	if err != nil {
		return nil, err
	}
	ins.Number = count.Value
	ins.Id = fmt.Sprintf("%s.%d", ins.ParentPath, ins.Number)

	return ins, nil
}

// NewDefaultDepartment todo
func NewDefaultDepartment() *Department {
	return &Department{}
}

// Update todo
func (d *Department) Update(req *CreateDepartmentRequest) {
	d.Name = req.Name
	d.DisplayName = req.DisplayName
	d.ParentId = req.ParentId
	d.Manager = req.Manager
	d.DefaultRoleId = req.DefaultRoleId
}

// Patch todo
func (d *Department) Patch(req *CreateDepartmentRequest) {
	patchData, _ := json.Marshal(req)
	json.Unmarshal(patchData, d)
}

// HasSubDepartment todo
func (d *Department) HasSubDepartment() bool {
	return d.SubCount > 0
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
	return &CreateDepartmentRequest{}
}

// Validate 校验参数的合法性
func (req *CreateDepartmentRequest) Validate() error {
	return validate.Struct(req)
}

// Patch todo
func (req *CreateDepartmentRequest) Patch(data *CreateDepartmentRequest) {
	patchData, _ := json.Marshal(data)
	json.Unmarshal(patchData, req)
}

// NewDepartmentSet 实例化
func NewDepartmentSet() *Set {
	return &Set{
		Items: []*Department{},
	}
}

// Add 添加应用
func (s *Set) Add(item *Department) {
	s.Items = append(s.Items, item)
}

// NewPutUpdateDepartmentRequest todo
func NewPutUpdateDepartmentRequest(id string) *UpdateDepartmentRequest {
	return &UpdateDepartmentRequest{
		Id:         id,
		UpdateMode: types.UpdateMode_PUT,
		Data:       NewCreateDepartmentRequest(),
	}
}

// NewPatchUpdateDepartmentRequest todo
func NewPatchUpdateDepartmentRequest(id string) *UpdateDepartmentRequest {
	return &UpdateDepartmentRequest{
		Id:         id,
		UpdateMode: types.UpdateMode_PATCH,
		Data:       NewCreateDepartmentRequest(),
	}
}

// Validate 校验入参
func (req *UpdateDepartmentRequest) Validate() error {
	if req.Id == "" {
		return fmt.Errorf("department id requred")
	}

	return req.Data.Validate()
}
