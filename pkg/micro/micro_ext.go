package micro

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// New 创建服务
func New(req *CreateMicroRequest) (*Micro, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	ins := &Micro{
		Id:            xid.New().String(),
		CreateAt:      ftime.Now().Timestamp(),
		UpdateAt:      ftime.Now().Timestamp(),
		Enabled:       true,
		Type:          req.Type,
		Name:          req.Name,
		Label:         req.Label,
		Description:   req.Description,
		ClientId:      token.MakeBearer(24),
		ClientSecret:  token.MakeBearer(32),
		ClientEnabled: true,
	}

	return ins, nil
}

// Desensitize 数据脱敏
func (m *Micro) Desensitize() {
	m.ClientSecret = ""
}

// ValiateClientCredential todo
func (m *Micro) ValiateClientCredential(clientSecret string) error {
	if !m.ClientEnabled {
		return exception.NewBadRequest("client not enabled")
	}

	if m.ClientSecret != clientSecret {
		return exception.NewUnauthorized("client credentail invalidate")
	}
	return nil
}

// NewCreateMicroRequest todo
func NewCreateMicroRequest() *CreateMicroRequest {
	return &CreateMicroRequest{
		Label: map[string]string{},
		Type:  Type_CUSTOM,
	}
}

// Validate 校验请求是否合法
func (req *CreateMicroRequest) Validate() error {
	return validate.Struct(req)
}

// NewMicroSet 实例化
func NewMicroSet() *Set {
	return &Set{
		Items: []*Micro{},
	}
}

// Add 添加
func (s *Set) Add(e *Micro) {
	s.Items = append(s.Items, e)
}
