package namespace

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"

	"github.com/infraboard/keyauth/common/session"
	"github.com/infraboard/keyauth/pkg/department"
)

// use a single instance of Validate, it caches struct info
var (
	validater = validator.New()
)

// NewNamespace todo
func NewNamespace(ctx context.Context, req *CreateNamespaceRequest, depart department.DepartmentServiceServer) (*Namespace, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	tk := session.GetTokenFromContext(ctx)
	ins := &Namespace{
		Id:       xid.New().String(),
		Domain:   tk.Domain,
		Creater:  tk.Account,
		CreateAt: ftime.Now().Timestamp(),
		UpdateAt: ftime.Now().Timestamp(),
		Data:     req,
	}

	descD := department.NewDescribeDepartmentRequest()
	descD.Id = req.DepartmentId
	d, err := depart.DescribeDepartment(ctx, descD)
	if err != nil {
		return nil, err
	}
	// 部门负责人就是空间负责人
	ins.Data.Owner = d.Data.Manager

	return ins, nil
}

// NewDefaultNamespace todo
func NewDefaultNamespace() *Namespace {
	return &Namespace{
		Data: NewCreateNamespaceRequest(),
	}
}

// NewCreateNamespaceRequest todo
func NewCreateNamespaceRequest() *CreateNamespaceRequest {
	return &CreateNamespaceRequest{
		Enabled: true,
	}
}

// Validate todo
func (req *CreateNamespaceRequest) Validate() error {
	return validater.Struct(req)
}

// NewNamespaceSet 实例化
func NewNamespaceSet() *Set {
	return &Set{
		Items: []*Namespace{},
	}
}

// Add 添加应用
func (s *Set) Add(item *Namespace) {
	s.Items = append(s.Items, item)
}
