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
		&i.domainName,
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
	domainName string
	username   string
	password   string
}

// Run 执行初始化
func (i *Initialer) Run() error {
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

	a, err := i.initApp(u.ID)
	if err != nil {
		return err
	}
	fmt.Println("初始化应用 [成功]")
	fmt.Printf("应用客户端ID: %s\n", a.ClientID)
	fmt.Printf("应用客户端凭证: %s\n", a.ClientSecret)

	return nil
}

func (i *Initialer) checkIsInit() error {
	req := user.NewQueryAccountRequest(request.NewPageRequest(20, 1))
	_, total, err := pkg.User.QuerySupperAccount(req)
	if err != nil {
		return err
	}

	if total > 0 {
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
	req.Name = i.domainName
	return pkg.Domain.CreateDomain(ownerID, req)
}

func (i *Initialer) initApp(ownerID string) (*application.Application, error) {
	req := application.NewCreateApplicatonRequest()
	req.Name = "初始化应用"
	return pkg.Application.CreateUserApplication(ownerID, req)
}

func init() {
	RootCmd.AddCommand(InitCmd)
}
