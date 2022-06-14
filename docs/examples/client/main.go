package main

import (
	"context"
	"fmt"

	"github.com/infraboard/keyauth/apps/domain"
	"github.com/infraboard/keyauth/apps/user"
	"github.com/infraboard/keyauth/client/rpc"
)

func main() {
	conf := rpc.NewDefaultConfig()
	// 提前注册一个服务, 获取服务的client_id和client_secret
	conf.ClientID = "pz3HiVQA3indzSHzFKtLHaJW"
	conf.ClientSecret = "vDvlAtqN3rS9CZcHugXp6QBuk28zRjud"
	c, err := rpc.NewClient(conf)
	if err != nil {
		panic(err)
	}

	// 查询用户信息, 查询admin domain里面的用户
	req := user.NewQueryAccountRequest()
	req.Domain = domain.AdminDomainName
	eps, err := c.User().QueryAccount(context.Background(), req)
	fmt.Println(eps, err)
}
