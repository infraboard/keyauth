package resolver

import (
	"sync"

	"google.golang.org/grpc/resolver"
)

// 全局Manager对象
var M = manager{}

type manager struct {
	resolvers []*mcenterResolver
	lock      sync.Mutex
}

func (m *manager) add(r *mcenterResolver) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.resolvers == nil {
		m.resolvers = []*mcenterResolver{}
	}

	m.resolvers = append(m.resolvers, r)
}

func (m *manager) Update() {
	for i := range m.resolvers {
		m.resolvers[i].ResolveNow(resolver.ResolveNowOptions{})
	}
}
