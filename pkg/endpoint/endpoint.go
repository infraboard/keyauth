package endpoint

import (
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/types/ftime"
)

// NewDefaultEndpoint todo
func NewDefaultEndpoint() *Endpoint {
	return &Endpoint{}
}

// NewEndpoint todo
func NewEndpoint(serviceID, version string, entry router.Entry) *Endpoint {
	return &Endpoint{
		Id:        GenHashID(serviceID, entry.Path, entry.Method),
		CreateAt:  ftime.Now().Timestamp(),
		UpdateAt:  ftime.Now().Timestamp(),
		ServiceId: serviceID,
		Version:   version,
		Entry:     &entry,
	}
}

// GenHashID hash id
func GenHashID(service, path, method string) string {
	hashedStr := fmt.Sprintf("%s-%s-%s", service, path, method)
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
