package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/register"
	"github.com/infraboard/mcube/register/etcd"
	"github.com/spf13/cobra"

	"github.com/infraboard/keyauth/api"
	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/version"
)

var (
	// pusher service config option
	confType   string
	confFile   string
	confEtcd   string
	isEncrypto bool
)

// startCmd represents the start command
var serviceCmd = &cobra.Command{
	Use:   "service [start/stop/reload/restart]",
	Short: "权限中心服务",
	Long:  `权限中心服务`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("[start] are required")
		}

		// 初始化全局变量
		if err := loadGloabl(confType); err != nil {
			return err
		}

		conf := conf.C()
		switch args[0] {
		case "start":
			// 启动服务
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)

			// 初始化服务
			svr, err := newService(conf)
			if err != nil {
				return err
			}

			// 注册服务
			r, err := etcd.NewEtcdRegister(conf.Etcd.EndPoints, conf.Etcd.Username, conf.Etcd.Password, zap.L().Named("Register"))
			if err != nil {
				svr.log.Warn(err)
			}
			defer r.UnRegiste()

			if err := svr.registry(r, conf); err != nil {
				return err
			}

			// 等待信号处理
			go svr.waitSign(ch)

			// 启动服务
			if err := svr.start(); err != nil {
				if !strings.Contains(err.Error(), "http: Server closed") {
					return err
				}
			}
		default:
			return errors.New("not support argument, support [start]")
		}

		return nil
	},
}

func newService(cnf *conf.Config) (*service, error) {
	http, err := api.NewHTTPService(cnf)
	if err != nil {
		return nil, err
	}

	svr := &service{
		http: http,
		log:  zap.L().Named("CLI"),
	}

	return svr, nil
}

type service struct {
	http *api.HTTPService
	hb   <-chan register.HeatbeatResonse

	log  logger.Logger
	stop context.CancelFunc
}

func (s *service) start() error {
	return s.http.Start()
}

// config 为全局变量, 只需要load 即可全局可用户
// 日志需要初始化并配置
func loadGloabl(configType string) error {
	// 配置加载
	switch configType {
	case "file":
		err := conf.LoadConfigFromTomlFile(confFile, isEncrypto)
		if err != nil {
			return err
		}
	case "env":
		return errors.New("not implemented")
	case "etcd":
		return errors.New("not implemented")
	default:
		return errors.New("unknown config type")
	}

	// 加载日志组件
	lc := conf.C().Log
	var (
		logInitMsg string
		level      zap.Level
	)

	lv, err := zap.NewLevel(lc.Level)
	if err != nil {
		logInitMsg = fmt.Sprintf("%s, use default level INFO", err)
		level = zap.InfoLevel
	} else {
		level = lv
		logInitMsg = fmt.Sprintf("log level: %s", lv)
	}

	if lc.ToStdout {
		// 输出到标准输出
		if lc.JSON {
			if err := zap.DevelopmentSetup(zap.WithLevel(level), zap.AsJSON()); err != nil {
				return err
			}
		} else {
			if err := zap.DevelopmentSetup(zap.WithLevel(level)); err != nil {
				return err
			}
		}
	} else {
		// 输出到文件
		logconf := zap.DefaultConfig()
		logconf.Files.Name = "api.log"
		logconf.Files.Path = lc.LogDirPath
		logconf.Level = level

		if err := zap.Configure(logconf); err != nil {
			return err
		}
	}

	zap.L().Named("Init").Info(logInitMsg)
	return nil
}

func (s *service) registry(r register.Register, conf *conf.Config) error {
	instance := &register.ServiceInstance{
		InstanceName: conf.APP.Name,
		ServiceName:  version.ServiceName,
		Type:         register.API,
		Address:      conf.APP.Host,
		Version:      version.GIT_TAG,
		GitBranch:    version.GIT_BRANCH,
		GitCommit:    version.GIT_COMMIT,
		BuildEnv:     version.GO_VERSION,
		BuildAt:      version.BUILD_TIME,
		Online:       time.Now().UnixNano() / 1000000, // 毫秒时间戳

		Prefix:   conf.Etcd.InstancePrefixKey,
		TTL:      conf.Etcd.InstanceTTL,
		Interval: time.Duration(conf.Etcd.InstanceTTL/3) * time.Second,
	}

	heatbeatResp, err := r.Registe(instance)
	if err != nil {
		return err
	}

	s.hb = heatbeatResp
	return nil
}

func (s *service) waitSign(sign chan os.Signal) {
	for {
		select {
		case sg := <-sign:
			switch v := sg.(type) {
			default:
				s.log.Infof("receive signal '%v', start graceful shutdown", v.String())
				if err := s.http.Stop(); err != nil {
					s.log.Errorf("graceful shutdown err: %s, force exit", err)
				}
				s.log.Infof("service stop complete")
				return
			}
		case <-s.hb:
			// s.log.Debug(hb.TTL())
		}
	}
}

func init() {
	serviceCmd.Flags().StringVarP(&confType, "config-type", "t", "file", "the service config type [file/env/etcd]")
	serviceCmd.Flags().StringVarP(&confFile, "config-file", "f", "etc/keyauth.toml", "the service config from file")
	serviceCmd.Flags().StringVarP(&confEtcd, "config-etcd", "e", "127.0.0.1:2379", "the service config from etcd")
	serviceCmd.Flags().BoolVarP(&isEncrypto, "is-encrypto", "p", false, "config is encrypto")
	RootCmd.AddCommand(serviceCmd)
}
