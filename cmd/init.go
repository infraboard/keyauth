package cmd

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/label"
	"github.com/spf13/cobra"

	"github.com/infraboard/keyauth/app/application"
	"github.com/infraboard/keyauth/app/department"
	"github.com/infraboard/keyauth/app/domain"
	"github.com/infraboard/keyauth/app/micro"
	"github.com/infraboard/keyauth/app/role"
	"github.com/infraboard/keyauth/app/system"
	"github.com/infraboard/keyauth/app/token"
	"github.com/infraboard/keyauth/app/user"
	"github.com/infraboard/keyauth/app/user/types"
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

		// 加载缓存
		if err := loadCache(); err != nil {
			return err
		}

		// 初始化服务层
		app.InitAllApp()

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

	if err := i.checkIsInit(context.Background()); err != nil {
		return nil, err
	}

	err := survey.AskOne(
		&survey.Input{
			Message: "请输入公司(组织)名称:",
			Default: "基础设施服务中心",
		},
		&i.domainDisplayName,
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
		domainName:  domain.AdminDomainName,
		user:        app.GetGrpcApp(user.AppName).(user.ServiceServer),
		domain:      app.GetGrpcApp(domain.AppName).(domain.ServiceServer),
		application: app.GetGrpcApp(application.AppName).(application.ServiceServer),
		token:       app.GetGrpcApp(token.AppName).(token.ServiceServer),
		department:  app.GetGrpcApp(department.AppName).(department.ServiceServer),
		role:        app.GetGrpcApp(role.AppName).(role.ServiceServer),
		micro:       app.GetGrpcApp(micro.AppName).(micro.ServiceServer),
		system:      app.GetInternalApp(system.AppName).(system.Service),
	}
}

// Initialer 初始化控制器
type Initialer struct {
	domainDisplayName string
	domainName        string
	username          string
	password          string
	adminUser         *user.User
	adminRole         *role.Role

	user        user.ServiceServer
	domain      domain.ServiceServer
	application application.ServiceServer
	token       token.ServiceServer
	department  department.ServiceServer
	role        role.ServiceServer
	micro       micro.ServiceServer
	system      system.Service
}

// Run 执行初始化
func (i *Initialer) Run() error {
	fmt.Println()
	fmt.Println("开始初始化...")

	ctx := context.Background()

	u, err := i.initUser(ctx)
	if err != nil {
		return err
	}
	i.adminUser = u
	fmt.Printf("初始化用户: %s [成功]\n", i.username)

	_, err = i.initDomain(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("初始化域: %s   [成功]\n", i.domainDisplayName)

	apps, err := i.initApp(ctx)
	if err != nil {
		return err
	}
	for index := range apps {
		fmt.Printf("初始化应用: %s [成功]\n", apps[index].Name)
		fmt.Printf("应用客户端ID: %s\n", apps[index].ClientId)
		fmt.Printf("应用客户端凭证: %s\n", apps[index].ClientSecret)
	}

	roles, err := i.initRole(ctx)
	if err != nil {
		return err
	}
	for index := range roles {
		r := roles[index]
		fmt.Printf("初始化角色: %s [成功]\n", r.Name)
	}

	dep, err := i.initDepartment(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("初始化根部门: %s   [成功]\n", dep.DisplayName)

	sysconf, err := i.initSystemConfig()
	if err != nil {
		return err
	}
	fmt.Printf("初始化系统配置: %s   [成功]\n", sysconf.Version)

	return nil
}

func (i *Initialer) checkIsInit(ctx context.Context) error {
	req := user.NewQueryAccountRequest()
	req.UserType = types.UserType_SUPPER
	userSet, err := i.user.QueryAccount(ctx, req)
	if err != nil {
		return err
	}

	if userSet.Total > 0 {
		return errors.New("supper admin user has exist")
	}
	return nil
}

func (i *Initialer) initUser(ctx context.Context) (*user.User, error) {
	req := user.NewCreateUserRequest()
	req.UserType = types.UserType_SUPPER
	req.Account = strings.TrimSpace(i.username)
	req.Password = strings.TrimSpace(i.password)
	req.Domin = i.domainName
	return i.user.CreateAccount(ctx, req)
}

func (i *Initialer) initDomain(ctx context.Context) (*domain.Domain, error) {
	req := domain.NewCreateDomainRequest()
	req.Name = i.domainName
	req.Owner = i.adminUser.Account
	req.Profile.DisplayName = strings.TrimSpace(i.domainDisplayName)
	return i.domain.CreateDomain(ctx, req)
}

func (i *Initialer) initApp(ctx context.Context) ([]*application.Application, error) {
	req := application.NewCreateApplicatonRequest()
	req.Name = application.AdminWebApplicationName
	req.ClientType = application.ClientType_PUBLIC
	req.Description = "Admin Web管理端"
	req.BuildIn = true
	req.Domain = i.domainName
	req.CreateBy = i.adminUser.Account

	web, err := i.application.CreateApplication(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("create admin web applicaton error, %s", err)
	}

	req = application.NewCreateApplicatonRequest()
	req.Name = application.AdminServiceApplicationName
	req.ClientType = application.ClientType_CONFIDENTIAL
	req.Description = "Admin Service 内置管理端, 服务注册后, 使用该端管理他们的凭证, 默认token不过期"
	req.AccessTokenExpireSecond = 0
	req.RefreshTokenExpireSecond = 0
	req.BuildIn = true
	req.Domain = i.domainName
	req.CreateBy = i.adminUser.Account
	svr, err := i.application.CreateApplication(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("create admin web applicaton error, %s", err)
	}

	apps := []*application.Application{web, svr}
	return apps, nil
}

func (i *Initialer) initRole(ctx context.Context) ([]*role.Role, error) {
	admin := role.NewDefaultPermission()
	admin.ServiceId = "*"
	admin.ResourceName = "*"
	admin.LabelKey = "*"
	admin.LabelValues = []string{"*"}

	req := role.NewCreateRoleRequest()
	req.Name = role.AdminRoleName
	req.Description = "系统管理员, 有系统所有功能的访问权限"
	req.Permissions = []*role.CreatePermssionRequest{admin}
	req.Type = role.RoleType_BUILDIN
	req.CreateBy = i.adminUser.Account
	adminRole, err := i.role.CreateRole(ctx, req)
	if err != nil {
		return nil, err
	}
	i.adminRole = adminRole

	vistor := role.NewDefaultPermission()
	vistor.ServiceId = "*"
	vistor.ResourceName = "*"
	vistor.LabelKey = label.ActionLableKey
	vistor.LabelValues = []string{label.Get.Value(), label.List.Value()}

	req = role.NewCreateRoleRequest()
	req.Name = role.VisitorRoleName
	req.Description = "访客, 登录系统后, 默认的权限"
	req.Permissions = []*role.CreatePermssionRequest{vistor}
	req.Type = role.RoleType_BUILDIN
	req.CreateBy = i.adminUser.Account
	vistorRole, err := i.role.CreateRole(ctx, req)
	if err != nil {
		return nil, err
	}

	return []*role.Role{adminRole, vistorRole}, nil
}

func (i *Initialer) initDepartment(ctx context.Context) (*department.Department, error) {
	req := department.NewCreateDepartmentRequest()
	req.Name = department.DefaultDepartmentName
	req.DisplayName = i.domainDisplayName
	req.Domain = i.domainName
	req.Manager = strings.TrimSpace(i.username)
	req.CreateBy = i.adminUser.Account
	return i.department.CreateDepartment(ctx, req)
}

func (i *Initialer) initService(ctx context.Context, r *role.Role) (*micro.Micro, error) {
	req := micro.NewCreateMicroRequest()
	req.Name = version.ServiceName
	req.Description = version.Description
	req.Type = micro.Type_BUILD_IN
	req.Creater = i.adminUser.Account
	return i.micro.CreateService(ctx, req)
}

func (i *Initialer) initSystemConfig() (*system.Config, error) {
	sysConf := system.NewDefaultConfig()
	if err := i.system.InitConfig(sysConf); err != nil {
		return nil, err
	}
	return sysConf, nil
}

func init() {
	RootCmd.AddCommand(InitCmd)
}
