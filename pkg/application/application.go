package application

import (
	"errors"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"

	"github.com/infraboard/keyauth/pkg/token"
)

// ClientType 客户端类型
type ClientType string

const (
	// Confidential (server-based) https://tools.ietf.org/html/rfc6749#section-2.1
	Confidential ClientType = "confidential"
	// Public （client-based)
	Public ClientType = "public"
)

// NewUserApplicartion 新建实例
func NewUserApplicartion(account string, req *CreateApplicatonRequest) (*Application, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	app := newDeafultApplication(req)
	app.User = account

	return app, nil
}

// NewBuildInApplication 构建内建应用
func NewBuildInApplication(account string, req *CreateApplicatonRequest) (*Application, error) {
	app, err := NewUserApplicartion(account, req)
	if err != nil {
		return nil, err
	}
	app.BuildIn = true

	return app, nil
}

func newDeafultApplication(req *CreateApplicatonRequest) *Application {
	return &Application{
		ID:                      xid.New().String(),
		BuildIn:                 false,
		CreateAt:                ftime.Now(),
		UpdateAt:                ftime.Now(),
		ClientID:                token.MakeBearer(16),
		ClientSecret:            token.MakeBearer(32),
		CreateApplicatonRequest: req,
	}
}

// Application is oauth2's client: https://tools.ietf.org/html/rfc6749#section-2
type Application struct {
	ID                       string     `bson:"_id" json:"id,omitempty"`                      // 唯一ID
	BuildIn                  bool       `bson:"build_in" json:"build_in"`                     // 是否是内建应用
	Domain                   string     `bosn:"domain" json:"domain,omitempty"`               // 所处于域
	User                     string     `bson:"user" json:"user,omitempty"`                   // 应用属于那个用户
	CreateAt                 ftime.Time `bson:"create_at" json:"create_at,omitempty"`         // 应用创建的时间
	UpdateAt                 ftime.Time `bson:"update_at" json:"update_at,omitempty"`         // 应用更新的时间
	ClientID                 string     `bson:"client_id" json:"client_id,omitempty"`         // 应用客户端ID
	ClientSecret             string     `bson:"client_secret" json:"client_secret,omitempty"` // 应用客户端秘钥
	Locked                   bool       `bson:"locked" json:"locked,omitempty"`               // 是否冻结应用, 冻结应用后, 该应用无法通过凭证获取访问凭证(token)
	*CreateApplicatonRequest `bson:",inline"`
}

// CheckClientSecret 判断凭证是否合法
func (a *Application) CheckClientSecret(secret string) error {
	if a.ClientSecret != secret {
		return errors.New("client_secret is not correct")
	}

	return nil
}

// NewApplicationSet 实例化
func NewApplicationSet(req *request.PageRequest) *Set {
	return &Set{
		PageRequest: req,
	}
}

// Set 应用列表
type Set struct {
	*request.PageRequest

	Total int64          `json:"total"`
	Items []*Application `json:"items"`
}

// Add 添加应用
func (s *Set) Add(app *Application) {
	s.Items = append(s.Items, app)
}
