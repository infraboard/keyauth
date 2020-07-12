package namespace

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"

	"github.com/infraboard/keyauth/pkg/token"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewNamespace todo
func NewNamespace(req *CreateNamespaceRequest) (*Namespace, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	tk := req.GetToken()

	ins := &Namespace{
		ID:                     xid.New().String(),
		DomainID:               tk.DomainID,
		CreaterID:              tk.UserID,
		CreateAt:               ftime.Now(),
		UpdateAt:               ftime.Now(),
		CreateNamespaceRequest: NewCreateNamespaceRequest(),
	}

	return ins, nil
}

// NewDefaultNamespace todo
func NewDefaultNamespace() *Namespace {
	return &Namespace{
		CreateNamespaceRequest: NewCreateNamespaceRequest(),
	}
}

// Namespace tenant resource container
type Namespace struct {
	ID                      string     `bson:"_id" json:"id,omitempty"`                // 项目唯一ID
	DomainID                string     `bson:"domain_id" json:"domain_id,omitempty"`   // 所属域ID
	CreaterID               string     `bson:"creater_id" json:"creater_id,omitempty"` // 创建人
	CreateAt                ftime.Time `bson:"create_at" json:"create_at,omitempty"`   // 创建时间
	UpdateAt                ftime.Time `bson:"update_at" json:"update_at,omitempty"`   // 项目修改时间
	*CreateNamespaceRequest `bson:",inline"`
}

// NewCreateNamespaceRequest todo
func NewCreateNamespaceRequest() *CreateNamespaceRequest {
	return &CreateNamespaceRequest{
		Session: token.NewSession(),
		Enabled: true,
	}
}

// CreateNamespaceRequest 创建项目请求
type CreateNamespaceRequest struct {
	*token.Session `bson:"-" json:"-"`
	Name           string `bson:"name" json:"name,omitempty" validate:"required,lte=80"` // 项目名称
	Picture        string `bson:"picture" json:"picture,omitempty"`                      // 项目描述图片
	Enabled        bool   `bson:"enabled" json:"enabled,omitempty"`                      // 禁用项目, 该项目所有人暂时都无法访问
	OwnerID        string `bson:"owner_id" json:"owner_id,omitempty"`                    // 项目所有者, PMO
	Description    string `bson:"description" json:"description,omitempty"`              // 项目描述
}

// Validate todo
func (req *CreateNamespaceRequest) Validate() error {
	if req.GetToken() == nil {
		return fmt.Errorf("token required")
	}

	return validate.Struct(req)
}

// NewNamespaceSet 实例化
func NewNamespaceSet(req *request.PageRequest) *Set {
	return &Set{
		PageRequest: req,
		Items:       []*Namespace{},
	}
}

// Set 列表
type Set struct {
	*request.PageRequest

	Total int64        `json:"total"`
	Items []*Namespace `json:"items"`
}

// Add 添加应用
func (s *Set) Add(item *Namespace) {
	s.Items = append(s.Items, item)
}
