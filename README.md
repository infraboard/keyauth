# keyauth

[![Go Report Card](https://goreportcard.com/badge/github.com/infraboard/keyauth)](https://goreportcard.com/report/github.com/infraboard/keyauth)
![](https://img.shields.io/github/license/infraboard/keyauth)

keyauth是一个微服务权限治理中心, 提供用户认证与授权管理.

demo访问地址: [Demo](http://keyauth.nbtuan.vip/) 用户: admin, 密码: 12345678mM@1

## 功能

+ 认证
    + 用户认证: Web用户访问(Password模式/LDAP模式)/API访问(AccessToken模式)
    + 内部服务认证: AccessToken模式
    + 第三方应用认证: Oauth2.0 auth code认证(比如gitlab, github)

+ 鉴权: 服务将功能注册到keyauth, keyauth基于这些服务功能提供RBAC鉴权机制, 系统角色分类如下
    + 超级管理员: 访问系统所有功能的权限
    + 超级访客:   访问所有系统的只读权限
    + 普通访客：  无权访问所有服务, 但是可以看到服务目录, 用户和申请开通访问那么服务
    + 服务管理员: 访问该服务所有功能的权限
    + 服务Reader: 服务功能的只读权限(GET 方法访问的所有功能)
    + 自定义角色:  直接选择功能定制一个角色的访问权限

+ 多租户:
    + 用户管理: 以域(公司)为单位提供人员管理功能, 

+ 项目: 资源容器
    + 服务提供的资源

+ 审计：
    + 用户登录审计:  用户什么时间在哪儿登录了系统


## 快速开发

```sh
# 安装依赖
make dep
# 配置
mv /etc/keyauth_sample.toml /etc/keyauth.toml
vim /etc/keyauth.toml
# 初始化服务
make init
# 启动服务
make run
```