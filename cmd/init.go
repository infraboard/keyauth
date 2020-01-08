package cmd

import (
	"fmt"

	"github.com/infraboard/keyauth/pkg"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
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

		initer := newIniter()
		initer.scanParamsFromCLI()
		// conf := conf.C()
		// fmt.Println(conf)
		return nil
	},
}

func newIniter() *Initialer {
	return &Initialer{}
}

// Initialer 初始化控制器
type Initialer struct {
	domainName string
	username   string
	password   string
}

func (i *Initialer) scanParamsFromCLI() error {
	fmt.Print("请输入公司(组织)名称: ")
	fmt.Scan(&i.domainName)
	fmt.Print("请输入admin用户名: ")
	fmt.Scan(&i.username)
	fmt.Print("请输入admin密码: ")
	bytePassword, err := terminal.ReadPassword(0)
	if err != nil {
		return fmt.Errorf("read password from cli error, %s", err)
	}
	i.password = string(bytePassword)
	fmt.Println(i)
	return nil
}

func init() {
	RootCmd.AddCommand(InitCmd)
}
