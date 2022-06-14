package resolver

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/infraboard/keyauth/apps/instance"
	"github.com/infraboard/keyauth/client/rpc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc/resolver"
)

const (
	Scheme = "mcenter"
)

// Following is an example name resolver. It includes a
// ResolverBuilder(https://godoc.org/google.golang.org/grpc/resolver#Builder)
// and a Resolver(https://godoc.org/google.golang.org/grpc/resolver#Resolver).
//
// A ResolverBuilder is registered for a scheme (in this example, "example" is
// the scheme). When a ClientConn is created for this scheme, the
// ResolverBuilder will be picked to build a Resolver. Note that a new Resolver
// is built for each ClientConn. The Resolver will watch the updates for the
// target, and send updates to the ClientConn.

// McenterResolverBuilder is a ResolverBuilder.
type McenterResolverBuilder struct{}

func (*McenterResolverBuilder) Build(
	target resolver.Target,
	cc resolver.ClientConn,
	opts resolver.BuildOptions) (
	resolver.Resolver, error) {

	r := &mcenterResolver{
		mcenter:            rpc.C().Instance(),
		target:             target,
		cc:                 cc,
		queryTimeoutSecond: 3 * time.Second,
		log:                zap.L().Named("mcenter resolver"),
	}

	// 强制触发一次更新
	r.ResolveNow(resolver.ResolveNowOptions{})

	// 添加给Manger管理, Manager负责更新
	M.add(r)
	return r, nil
}

func (*McenterResolverBuilder) Scheme() string {
	return Scheme
}

// exampleResolver is a
// Resolver(https://godoc.org/google.golang.org/grpc/resolver#Resolver).
type mcenterResolver struct {
	mcenter instance.ServiceClient

	target             resolver.Target
	cc                 resolver.ClientConn
	queryTimeoutSecond time.Duration
	log                logger.Logger
}

func (m *mcenterResolver) ResolveNow(o resolver.ResolveNowOptions) {
	// 从mcenter中查询该target对应的服务实例
	addrs, err := m.search()
	if err != nil {
		m.log.Errorf("search target %s error, %s", m.target.URL.String(), err)
	}

	// 更新给client
	m.cc.UpdateState(resolver.State{Addresses: addrs})
}

// 构建mcenter实例查询参数
func (m *mcenterResolver) buildSerchReq() *instance.SearchRequest {
	searchReq := instance.NewSearchRequest()
	searchReq.ServiceName = m.target.URL.Host

	qs := m.target.URL.Query()
	searchReq.Page.PageSize = 500
	searchReq.Region = qs.Get("region")
	searchReq.Environment = qs.Get("environment")
	return searchReq
}

// 查询名称对应的实例
func (m *mcenterResolver) search() ([]resolver.Address, error) {
	req := m.buildSerchReq()

	if req.ServiceName == "" {
		return nil, fmt.Errorf("application name required")
	}

	// 设置查询的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), m.queryTimeoutSecond)
	defer cancel()

	set, err := m.mcenter.Search(ctx, req)
	if err != nil {
		m.log.Errorf("search target %s error, %s", m.target, err)
		return nil, err
	}

	// 优先获取对应匹配的组, 如果没匹配的，则使用最老的那个组
	items := set.GetGroupInstance(req.Group)
	if len(items) == 0 {
		items = set.GetOldestGroup()
	}

	addrString := []string{}
	addrs := make([]resolver.Address, len(items))
	for i, s := range items {
		addrs[i] = resolver.Address{Addr: s.RegistryInfo.Address}
		addrString = append(addrString, s.RegistryInfo.Address)
	}
	m.log.Infof("search service [%s] address: %s", req.ServiceName, strings.Join(addrString, ","))

	return addrs, nil
}

func (m *mcenterResolver) Close() {

}

func init() {
	// Register the mcenter ResolverBuilder. This is usually done in a package's
	// init() function.
	resolver.Register(&McenterResolverBuilder{})
}
