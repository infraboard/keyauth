# keyauth

[![Go Report Card](https://goreportcard.com/badge/github.com/infraboard/keyauth)](https://goreportcard.com/report/github.com/infraboard/keyauth)
![](https://img.shields.io/github/license/infraboard/keyauth)

keyauth是一个微服务权限治理中心, 提供用户认证与授权管理.

demo访问地址: [Demo](http://keyauth.nbtuan.vip/) 用户: admin, 密码: 12345678mM@1

## 功能

+ 登录安全
    + 用户认证: Password/LDAP/AccessToken/Oauth2.0其他模式
    + 单点登录: 同一账号同一时间只能在一个客户端登录
    + 密码安全: 密码强度校验, 密码过期提醒, 密码重复限制
    + IP黑白名: 用户登录IP黑白名单限制
    + 异常登录: 失败重试检测, 异地登录检测, 30天未登录检测
    + 智能多因子认证: 异常登录时 支持邮件/短信方式对用户进行二次身份验证
    + 登录日志: 记录用户登录时间, 地点, IP, 登录端 等

+ 权限管理: 
    + 部门管理: 部门人员和部门工作空间管理
    + 权限模型: 基于工作空间的【RBAC】授权 
    + 角色管理: 基于标签的权限条目匹配, 灵活编辑角色
    + 服务目录: 服务将功能注册到keyauth, keyauth基于这些服务功能 提供RBAC鉴权机制

## 快速开发

1. 依赖环境搭建:

+ [MongoDB数据库安装](./docs/mongodb/install.md) (必须)
+ [OpenLDAP安装](./docs/ldap/install.md) (开启LDAP认证时需要安装)
+ [消息总线](./docs/bus/install.md) (开始操作审计时需要安装)
+ [安装protobuf](./docs/protobuf/install.md)(keyauth开发者)

2. 依赖的mcube protobuf文件处理
```
cd GOPATH/pkg/mod/github/infraboard
ln -s mcube@v1.1.2 mcube            // 每次更新版本都需要手动建立最新版本的link
```

2. 快速运行

```sh
# 安装依赖
make install

# 配置
mv /etc/keyauth_sample.toml /etc/keyauth.toml
vim /etc/keyauth.toml

# 初始化服务
make init

# 启动服务
make run
```