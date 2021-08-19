package department

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"

	"github.com/infraboard/keyauth/common/types"
	"github.com/infraboard/keyauth/pkg/token"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

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

func (req *CreateDepartmentRequest) UpdateOwner(tk *token.Token) {
	req.Domain = tk.Domain
	req.CreateBy = tk.Account
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
