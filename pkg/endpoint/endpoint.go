package endpoint

import (
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/types/ftime"
)

// NewDefaultEndpoint todo
func NewDefaultEndpoint() *Endpoint {
	return &Endpoint{}
}

// NewEndpoint todo
func NewEndpoint(serviceName, version string, entry router.Entry) *Endpoint {
	return &Endpoint{
		ID:       GenHashID(serviceName, entry.Path, entry.Method),
		CreateAt: ftime.Now(),
		UpdateAt: ftime.Now(),
		Service:  serviceName,
		Version:  version,
		Entry:    entry,
	}
}

// GenHashID hash id
func GenHashID(service, path, method string) string {
	hashedStr := fmt.Sprintf("%s-%s-%s", service, path, method)
	h := fnv.New32a()
	h.Write([]byte(hashedStr))
	return fmt.Sprintf("%x", h.Sum32())
}

// Endpoint Service's features
type Endpoint struct {
	ID       string     `bson:"_id" json:"id" validate:"required,lte=64"`                    // 端点名称
	CreateAt ftime.Time `bson:"create_at" json:"create_at,omitempty"`                        // 创建时间
	UpdateAt ftime.Time `bson:"update_at" json:"update_at,omitempty"`                        // 更新时间
	Service  string     `bson:"service" json:"service,omitempty" validate:"required,lte=64"` // 该功能属于那个服务
	Version  string     `bson:"version" json:"version,omitempty" validate:"required,lte=64"` // 服务那个版本的功能

	router.Entry `bson:",inline"`
}

// LabelsToStr 扁平化标签  action:get;action:list;action-list-echo
func (e *Endpoint) LabelsToStr() string {
	labels := make([]string, 0, len(e.Labels))
	for k, v := range e.Labels {
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
		e.Labels[kvItem[0]] = kvItem[1]
	}
	return nil
}

// NewEndpointSet 实例化
func NewEndpointSet(req *request.PageRequest) *Set {
	return &Set{
		PageRequest: req,
		Items:       []*Endpoint{},
	}
}

// Set 列表
type Set struct {
	*request.PageRequest

	Total int64       `json:"total"`
	Items []*Endpoint `json:"items"`
}

// Add 添加
func (s *Set) Add(e *Endpoint) {
	s.Items = append(s.Items, e)
}
