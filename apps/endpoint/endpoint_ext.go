package endpoint

import (
	"fmt"
	"hash/fnv"
	"strings"

	http "github.com/infraboard/mcube/pb/http"
	"github.com/infraboard/mcube/types/ftime"
)

// NewDefaultEndpoint todo
func NewDefaultEndpoint() *Endpoint {
	return &Endpoint{}
}

// NewEndpoint todo
func NewEndpoint(serviceID, version string, entry http.Entry) *Endpoint {
	return &Endpoint{
		Id:        GenHashID(serviceID, entry.Path),
		CreateAt:  ftime.Now().Timestamp(),
		UpdateAt:  ftime.Now().Timestamp(),
		ServiceId: serviceID,
		Version:   version,
		Entry:     &entry,
	}
}

// GenHashID hash id
func GenHashID(service, grpcPath string) string {
	hashedStr := fmt.Sprintf("%s-%s", service, grpcPath)
	h := fnv.New32a()
	h.Write([]byte(hashedStr))
	return fmt.Sprintf("%x", h.Sum32())
}

// LabelsToStr 扁平化标签  action:get;action:list;action-list-echo
func (e *Endpoint) LabelsToStr() string {
	labels := make([]string, 0, len(e.Entry.Labels))
	for k, v := range e.Entry.Labels {
		labels = append(labels, fmt.Sprintf("%s:%s;", k, v))
	}
	return strings.Join(labels, "")
}

// ParseLabels 解析Str格式的label
func (e *Endpoint) ParseLabels(labels string) error {
	kvs := strings.Split(strings.TrimSuffix(labels, ";"), ";")
	for _, kv := range kvs {
		kvItem := strings.Split(kv, ":")
		if len(kvItem) != 2 {
			return fmt.Errorf("labels format error, format: k:v;k:v;")
		}
		e.Entry.Labels[kvItem[0]] = kvItem[1]
	}
	return nil
}

// NewEndpointSet 实例化
func NewEndpointSet() *Set {
	return &Set{
		Items: []*Endpoint{},
	}
}

// Add 添加
func (s *Set) Add(e *Endpoint) {
	s.Items = append(s.Items, e)
}

// UpdatePath todo
func (r *Resource) UpdatePath(path string) {
	for _, p := range r.Paths {
		if p == path {
			return
		}
	}

	r.Paths = append(r.Paths, path)
}

// UpdateMethod todo
func (r *Resource) UpdateMethod(mothod string) {
	for _, p := range r.Methods {
		if p == mothod {
			return
		}
	}

	r.Methods = append(r.Methods, mothod)
}

// UpdateFunction todo
func (r *Resource) UpdateFunction(fuction string) {
	for _, p := range r.Functions {
		if p == fuction {
			return
		}
	}

	r.Functions = append(r.Functions, fuction)
}

// UpdateAction todo
func (r *Resource) UpdateAction(action string) {
	for _, p := range r.Actions {
		if p == action {
			return
		}
	}

	r.Actions = append(r.Actions, action)
}

// NewResourceSet todo
func NewResourceSet() *ResourceSet {
	return &ResourceSet{
		Items: []*Resource{},
	}
}

// AddEndpointSet todo
func (s *ResourceSet) AddEndpointSet(eps *Set) {
	for i := range eps.Items {
		s.addEndpint(eps.Items[i])
	}
}

func (s *ResourceSet) addEndpint(ep *Endpoint) {
	if ep.Entry == nil || ep.Entry.Resource == "" {
		return
	}

	rs := s.getOrCreateResource(ep.ServiceId, ep.Entry.Resource)
	rs.UpdateMethod(ep.Entry.Method)
	rs.UpdatePath(ep.Entry.Path)
	rs.UpdateFunction(ep.Entry.FunctionName)
	if v, ok := ep.Entry.Labels["action"]; ok {
		rs.UpdateAction(v)
	}
}

func (s *ResourceSet) getOrCreateResource(serviceID, name string) *Resource {
	var rs *Resource

	for i := range s.Items {
		rs = s.Items[i]
		if rs.ServiceId == serviceID && rs.Name == name {
			return rs
		}
	}

	// 添加新resource
	rs = &Resource{
		ServiceId: serviceID,
		Name:      name,
	}
	s.Items = append(s.Items, rs)
	return rs
}
