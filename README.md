# keyauth

[![Go Report Card](https://goreportcard.com/badge/github.com/infraboard/keyauth)](https://goreportcard.com/report/github.com/infraboard/keyauth)
![](https://img.shields.io/github/license/infraboard/keyauth)

keyauth是一个分布式或者微服务场景下的用户中心和鉴权中心

## 功能

+ 认证
    + 用户认证：提供用户名密码认证和API Key认证
    + 服务认证: 提供Service Key接入凭证
    + 第三方应用认证: Oauth2.0 auth code认证 （比如gitlab, github)

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

