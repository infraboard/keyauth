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


## SDK使用

注意SDK权限很大, 仅供内部服务调用, 如果对外，请走HTTP API
```go
func main() {
	conf := client.NewDefaultConfig()
    // 提前注册一个服务, 获取服务的client_id和client_secret
	conf.SetClientCredentials("VYizVq1fsK7olinqVHrBvFOl", "qS9FGBoFGRaVfbgeqFVDRcgH7nNJi9fp")
	c, err := client.NewClient(conf)
	if err != nil {
		panic(err)
	}
	page := request.NewPageRequest(20, 1)
	meta := metadata.Pairs("access_token", "NEjvVOhmhAQXFuYSrZdJaBsH")
	ctx := metadata.NewOutgoingContext(context.Background(), meta)
	eps, err := c.Endpoint().QueryEndpoints(ctx, endpoint.NewQueryEndpointRequest(page))
	fmt.Println(eps, err)
}
```

## 快速开发

1. 依赖环境搭建:

+ [MongoDB数据库安装](./docs/mongodb/install.md) (必须)
+ [OpenLDAP安装](./docs/ldap/install.md) (开启LDAP认证时需要安装)
+ [消息总线](./docs/bus/install.md) (开始操作审计时需要安装)
+ [安装protobuf](./docs/protobuf/install.md)(keyauth开发者)

2. grpc 环境准备
```sh
# 1.安装protoc编译器,  项目使用版本: v3.19.1
# 下载预编译包安装: https://github.com/protocolbuffers/protobuf/releases

# 2.protoc-gen-go go语言查询, 项目使用版本: v1.27.1   
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# 3.安装protoc-gen-go-grpc插件, 项目使用版本: 1.1.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 4.安装自定义proto tag插件
go install github.com/favadi/protoc-go-inject-tag@latest

# 安装项目依赖的protobuf
cp -r docs/include/github.com /usr/local/include
```

3. 快速运行

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