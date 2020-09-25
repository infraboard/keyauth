package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/infraboard/mcube/http/label"
	"github.com/spf13/cobra"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/application"
	"github.com/infraboard/keyauth/pkg/department"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/user/types"
	"github.com/infraboard/keyauth/version"
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
	i := NewInitialer()

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

// NewInitialer todo
func NewInitialer() *Initialer {
	return &Initialer{
		mockTK: &token.Token{
			UserType: types.SupperAccount,
			Domain:   domain.AdminDomainName,
		},
	}
}

// Initialer 初始化控制器
type Initialer struct {
	domainDesc string
	username   string
	password   string
	tk         *token.Token
	mockTK     *token.Token
}

// Run 执行初始化
func (i *Initialer) Run() error {
	fmt.Println()
	fmt.Println("开始初始化...")

	u, err := i.initUser()
	if err != nil {
		return err
	}
	fmt.Printf("初始化用户: %s [成功]\n", i.username)

	_, err = i.initDomain(u.Account)
	if err != nil {
		return err
	}
	fmt.Printf("初始化域: %s   [成功]\n", i.domainDesc)

	apps, err := i.initApp(u.Account)
	if err != nil {
		return err
	}
	for index := range apps {
		fmt.Printf("初始化应用: %s [成功]\n", apps[index].Name)
		fmt.Printf("应用客户端ID: %s\n", apps[index].ClientID)
		fmt.Printf("应用客户端凭证: %s\n", apps[index].ClientSecret)
	}

	if err := i.getAdminToken(apps[0], u); err != nil {
		return err
	}

	roles, err := i.initRole()
	if err != nil {
		return err
	}
	for index := range roles {
		fmt.Printf("初始化角色: %s [成功]\n", roles[index].Name)
	}

	svr, err := i.initService()
	if err != nil {
		return err
	}
	fmt.Printf("初始化服务: %s   [成功]\n", svr.Name)

	dep, err := i.initDepartment()
	if err != nil {
		return err
	}
	fmt.Printf("初始化部门: %s   [成功]\n", dep.DisplayName)

	return nil
}

func (i *Initialer) checkIsInit() error {
	req := user.NewQueryAccountRequest()
	req.WithToken(i.mockTK)
	userSet, err := pkg.User.QueryAccount(types.SupperAccount, req)
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
	req.WithToken(i.mockTK)
	req.Account = strings.TrimSpace(i.username)
	req.Password = strings.TrimSpace(i.password)
	return pkg.User.CreateAccount(types.SupperAccount, req)
}

func (i *Initialer) initDomain(ownerID string) (*domain.Domain, error) {
	req := domain.NewCreateDomainRequst()
	req.Name = domain.AdminDomainName
	req.Description = strings.TrimSpace(i.domainDesc)
	return pkg.Domain.CreateDomain(ownerID, req)
}

func (i *Initialer) initApp(ownerID string) ([]*application.Application, error) {
	tk := &token.Token{Account: ownerID}

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

func (i *Initialer) getAdminToken(app *application.Application, u *user.User) error {
	if app == nil || u == nil {
		return fmt.Errorf("get admin token need app and admin user")
	}

	req := token.NewIssueTokenByPassword(app.ClientID, app.ClientSecret, u.Account, u.Password)
	tk, err := pkg.Token.IssueToken(req)
	if err != nil {
		return err
	}
	i.tk = tk
	return nil
}

func (i *Initialer) initRole() ([]*role.Role, error) {
	admin := role.NewDefaultPermission()
	admin.ResourceName = "*"
	admin.LabelKey = "*"
	admin.LabelValues = []string{"*"}

	req := role.NewCreateRoleRequest()
	req.WithToken(i.tk)
	req.Name = role.AdminRoleName
	req.Description = "系统管理员, 有系统所有功能的访问权限"
	req.Permissions = []*role.Permission{admin}
	req.Type = role.BuildInType
	adminRole, err := pkg.Role.CreateRole(req)
	if err != nil {
		return nil, err
	}

	vistor := role.NewDefaultPermission()
	vistor.ResourceName = "*"
	vistor.LabelKey = label.ActionLableKey
	vistor.LabelValues = []string{label.Get.Value(), label.List.Value()}

	req = role.NewCreateRoleRequest()
	req.WithToken(i.tk)
	req.Name = role.VisitorRoleName
	req.Description = "访客, 登录系统后, 默认的权限"
	req.Permissions = []*role.Permission{vistor}
	req.Type = role.BuildInType
	vistorRole, err := pkg.Role.CreateRole(req)
	if err != nil {
		return nil, err
	}

	return []*role.Role{adminRole, vistorRole}, nil
}

func (i *Initialer) initDepartment() (*department.Department, error) {
	if pkg.Department == nil {
		return nil, fmt.Errorf("dependence service department is nil")
	}

	req := department.NewCreateDepartmentRequest()
	req.Name = department.DefaultDepartmentName
	req.DisplayName = "默认部门"
	req.WithToken(i.tk)
	return pkg.Department.CreateDepartment(req)
}

func (i *Initialer) initService() (*micro.Micro, error) {
	req := micro.NewCreateMicroRequest()
	req.WithToken(i.tk)
	req.Name = version.ServiceName
	req.Description = version.Description
	req.Label = map[string]string{"type": "build_in"}
	return pkg.Micro.CreateService(req)
}

func init() {
	RootCmd.AddCommand(InitCmd)
}
