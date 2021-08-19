package application

import (
	"errors"
	"time"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"github.com/rs/xid"

	"github.com/infraboard/keyauth/pkg/token"
)

// NewUserApplicartion 新建实例
func NewUserApplicartion(account string, req *CreateApplicatonRequest) (*Application, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	app := newDeafultApplication(req)
	app.CreateBy = account

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
	now := time.Now().Unix() * 1000
	return &Application{
		Id:           xid.New().String(),
		BuildIn:      false,
		CreateAt:     now,
		UpdateAt:     now,
		ClientId:     token.MakeBearer(24),
		ClientSecret: token.MakeBearer(32),

		Name:                     req.Name,
		Website:                  req.Website,
		LogoImage:                req.LogoImage,
		Description:              req.Description,
		RedirectUri:              req.RedirectUri,
		AccessTokenExpireSecond:  req.AccessTokenExpireSecond,
		RefreshTokenExpireSecond: req.RefreshTokenExpireSecond,
		ClientType:               req.ClientType,
	}
}

// CheckClientSecret 判断凭证是否合法
func (a *Application) CheckClientSecret(secret string) error {
	if a.ClientSecret != secret {
		return errors.New("client_secret is not correct")
	}

	return nil
}

func (a *Application) IsOwner(account string) bool {
	return a.CreateBy == account
}

// NewApplicationSet 实例化
func NewApplicationSet(req *request.PageRequest) *Set {
	return &Set{
		Items: []*Application{},
	}
}

// Add 添加应用
func (s *Set) Add(app *Application) {
	s.Items = append(s.Items, app)
}

// NewGetBuildInAdminApplicationRequest todo
func NewGetBuildInAdminApplicationRequest() *GetBuildInApplicationRequest {
	return &GetBuildInApplicationRequest{
		Name: AdminServiceApplicationName,
	}
}
