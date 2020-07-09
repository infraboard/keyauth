package cmd

import (
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/infraboard/mcube/http/request"
	"github.com/spf13/cobra"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/application"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user"
)

// InitCmd 初始化系统
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化服务",
	Long:  `初始化admin用户相关基础信息`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 初始化全局变量
		if err := loadGlobalConfig(confType); err != nil {
			return err
		}

		// 初始化全局日志配置
		if err := loadGlobalLogger(); err != nil {
			return err
		}

		// 初始化服务层
		if err := pkg.InitService(); err != nil {
			return err
		}

		initer, err := NewInitialerFromCLI()
		if err != nil {
			return err
		}
		if err := initer.Run(); err != nil {
			return err
		}
		return nil
	},
}

// NewInitialerFromCLI 初始化
func NewInitialerFromCLI() (*Initialer, error) {
	i := new(Initialer)

	if err := i.checkIsInit(); err != nil {
		return nil, err
	}

	err := survey.AskOne(
		&survey.Input{
			Message: "请输入公司(组织)名称:",
			Default: "基础设施服务中心",
		},
		&i.domainDesc,
		survey.WithValidator(survey.Required),
	)
	if err != nil {
		return nil, err
	}

	err = survey.AskOne(
		&survey.Input{
			Message: "请输入管理员用户名称:",
			Default: "admin",
		},
		&i.username,
		survey.WithValidator(survey.Required),
	)
	if err != nil {
		return nil, err
	}

	var repeatPass string
	err = survey.AskOne(
		&survey.Password{
			Message: "请输入管理员密码:",
		},
		&i.password,
		survey.WithValidator(survey.Required),
	)
	if err != nil {
		return nil, err
	}
	err = survey.AskOne(
		&survey.Password{
			Message: "再次输入管理员密码:",
		},
		&repeatPass,
		survey.WithValidator(survey.Required),
		survey.WithValidator(func(ans interface{}) error {
			if ans.(string) != i.password {
				return fmt.Errorf("两次输入的密码不一致")
			}
			return nil
		}),
	)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// Initialer 初始化控制器
type Initialer struct {
	domainDesc string
	username   string
	password   string
}

// Run 执行初始化
func (i *Initialer) Run() error {
	fmt.Println()
	fmt.Println("开始初始化...")

	u, err := i.initUser()
	if err != nil {
		return err
	}
	fmt.Println("初始化用户 [成功]")

	_, err = i.initDomain(u.ID)
	if err != nil {
		return err
	}
	fmt.Println("初始化域   [成功]")

	apps, err := i.initApp(u.ID)
	if err != nil {
		return err
	}
	for index := range apps {
		fmt.Printf("初始化应用: %s [成功]\n", apps[index].Name)
		fmt.Printf("应用客户端ID: %s\n", apps[index].ClientID)
		fmt.Printf("应用客户端凭证: %s\n", apps[index].ClientSecret)
	}

	return nil
}

func (i *Initialer) checkIsInit() error {
	req := user.NewQueryAccountRequest(request.NewPageRequest(20, 1))
	userSet, err := pkg.User.QuerySupperAccount(req)
	if err != nil {
		return err
	}

	if userSet.Total > 0 {
		return errors.New("supper admin user has exist")
	}
	return nil
}

func (i *Initialer) initUser() (*user.User, error) {
	req := user.NewCreateUserRequest()
	req.Account = i.username
	req.Password = i.password
	return pkg.User.CreateSupperAccount(req)
}

func (i *Initialer) initDomain(ownerID string) (*domain.Domain, error) {
	req := domain.NewCreateDomainRequst()
	req.Name = domain.AdminDomainName
	req.Description = i.domainDesc
	return pkg.Domain.CreateDomain(ownerID, req)
}

func (i *Initialer) initApp(ownerID string) ([]*application.Application, error) {
	tk := &token.Token{UserID: ownerID}

	req := application.NewCreateApplicatonRequest()
	req.Name = application.AdminWebApplicationName
	req.ClientType = application.Public
	req.Description = "Admin Web管理端"
	req.WithToken(tk)
	web, err := pkg.Application.CreateBuildInApplication(req)
	if err != nil {
		return nil, fmt.Errorf("create admin web applicaton error, %s", err)
	}

	req = application.NewCreateApplicatonRequest()
	req.Name = application.AdminServiceApplicationName
	req.ClientType = application.Confidential
	req.Description = "Admin Service 内置管理端, 服务注册后, 使用该端管理他们的凭证, 默认token不过期"
	req.AccessTokenExpireSecond = 0
	req.RefreshTokenExpiredSecond = 0
	req.WithToken(tk)
	svr, err := pkg.Application.CreateBuildInApplication(req)
	if err != nil {
		return nil, fmt.Errorf("create admin web applicaton error, %s", err)
	}

	apps := []*application.Application{web, svr}
	return apps, nil
}

func init() {
	RootCmd.AddCommand(InitCmd)
}
