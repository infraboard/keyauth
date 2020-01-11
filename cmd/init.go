package cmd

import (
	"errors"
	"fmt"

	"github.com/infraboard/mcube/http/request"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"

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

	fmt.Print("请输入公司(组织)名称: ")
	fmt.Scan(&i.domainName)
	fmt.Print("请输入admin用户名: ")
	fmt.Scan(&i.username)
	fmt.Print("请输入admin密码: ")
	bytePassword, err := terminal.ReadPassword(0)
	if err != nil {
		return nil, fmt.Errorf("read password from cli error, %s", err)
	}
	fmt.Println()
	fmt.Print("请再次输入admin密码: ")
	checkPassword, err := terminal.ReadPassword(0)
	if err != nil {
		return nil, fmt.Errorf("read password from cli error, %s", err)
	}
	if string(bytePassword) != string(checkPassword) {
		return nil, fmt.Errorf("两次密码输入不一致")
	}

	i.password = string(bytePassword)
	fmt.Println()

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
	fmt.Printf("初始化用户%s [成功]\n", u.Account)

	d, err := i.initDomain(u.ID)
	if err != nil {
		return err
	}
	fmt.Printf("初始化域%s [成功]\n", d.Name)

	a, err := i.initApp(u.ID)
	if err != nil {
		return err
	}
	fmt.Printf("初始化应用%s [成功]\n", a.Name)

	return nil
}

func (i *Initialer) checkIsInit() error {
	req := user.NewQueryAccountRequest(request.NewPageRequest(20, 1))
	_, total, err := pkg.User.QuerySupperAccount(req)
	if err != nil {
		return err
	}
	fmt.Println(total)

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
