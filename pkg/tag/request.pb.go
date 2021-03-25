//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/tag/pb/request.proto

package tag

import (
	proto "github.com/golang/protobuf/proto"
	page "github.com/infraboard/mcube/pb/page"
	_ "github.com/infraboard/protoc-gen-go-ext/extension/tag"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// CreateTagRequest todo
type CreateTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tag生效范围
	ScopeType ScopeType `protobuf:"varint,1,opt,name=scope_type,json=scopeType,proto3,enum=keyauth.tag.ScopeType" json:"scope_type"`
	// tag属于哪个namespace
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace"`
	// 建名称
	KeyName string `protobuf:"bytes,3,opt,name=key_name,json=keyName,proto3" json:"key_name"`
	// 建标识
	KeyLabel string `protobuf:"bytes,4,opt,name=key_label,json=keyLabel,proto3" json:"key_label"`
	// 建描述
	KeyDesc string `protobuf:"bytes,5,opt,name=key_desc,json=keyDesc,proto3" json:"key_desc"`
	// 值来源
	ValueFrom ValueFrom `protobuf:"varint,6,opt,name=value_from,json=valueFrom,proto3,enum=keyauth.tag.ValueFrom" json:"value_from"`
	// http 获取Tag 值的参数
	HttpFromOptions *HTTPFromOptions `protobuf:"bytes,7,opt,name=http_from_options,json=httpFromOptions,proto3" json:"http_from_options"`
	// String 类型的值
	Values []*ValueOptions `protobuf:"bytes,8,rep,name=values,proto3" json:"string_value"`
}

func (x *CreateTagRequest) Reset() {
	*x = CreateTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_tag_pb_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTagRequest) ProtoMessage() {}

func (x *CreateTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_tag_pb_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTagRequest.ProtoReflect.Descriptor instead.
func (*CreateTagRequest) Descriptor() ([]byte, []int) {
	return file_pkg_tag_pb_request_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTagRequest) GetScopeType() ScopeType {
	if x != nil {
		return x.ScopeType
	}
	return ScopeType_NAMESPACE
}

func (x *CreateTagRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateTagRequest) GetKeyName() string {
	if x != nil {
		return x.KeyName
	}
	return ""
}

func (x *CreateTagRequest) GetKeyLabel() string {
	if x != nil {
		return x.KeyLabel
	}
	return ""
}

func (x *CreateTagRequest) GetKeyDesc() string {
	if x != nil {
		return x.KeyDesc
	}
	return ""
}

func (x *CreateTagRequest) GetValueFrom() ValueFrom {
	if x != nil {
		return x.ValueFrom
	}
	return ValueFrom_MANUAL
}

func (x *CreateTagRequest) GetHttpFromOptions() *HTTPFromOptions {
	if x != nil {
		return x.HttpFromOptions
	}
	return nil
}

func (x *CreateTagRequest) GetValues() []*ValueOptions {
	if x != nil {
		return x.Values
	}
	return nil
}

// QueryTagRequest todo
type QueryTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *page.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// tag生效范围
	ScopeType ScopeType `protobuf:"varint,2,opt,name=scope_type,json=scopeType,proto3,enum=keyauth.tag.ScopeType" json:"scope_type"`
	// 关键字
	Keywords string `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords"`
	// 为了性能可能不需要查询出value
	SkipValues bool `protobuf:"varint,4,opt,name=skip_values,json=skipValues,proto3" json:"skip_values"`
}

func (x *QueryTagRequest) Reset() {
	*x = QueryTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_tag_pb_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTagRequest) ProtoMessage() {}

func (x *QueryTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_tag_pb_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTagRequest.ProtoReflect.Descriptor instead.
func (*QueryTagRequest) Descriptor() ([]byte, []int) {
	return file_pkg_tag_pb_request_proto_rawDescGZIP(), []int{1}
}

func (x *QueryTagRequest) GetPage() *page.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryTagRequest) GetScopeType() ScopeType {
	if x != nil {
		return x.ScopeType
	}
	return ScopeType_NAMESPACE
}

func (x *QueryTagRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *QueryTagRequest) GetSkipValues() bool {
	if x != nil {
		return x.SkipValues
	}
	return false
}

// DescribeTagRequest todo
type DescribeTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tag Value ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DescribeTagRequest) Reset() {
	*x = DescribeTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_tag_pb_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeTagRequest) ProtoMessage() {}

func (x *DescribeTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_tag_pb_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeTagRequest.ProtoReflect.Descriptor instead.
func (*DescribeTagRequest) Descriptor() ([]byte, []int) {
	return file_pkg_tag_pb_request_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeTagRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// ValueOptions 值描述
type ValueOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=Value,proto3" json:"value" bson:"value"`
	Label string `protobuf:"bytes,2,opt,name=Label,proto3" json:"label" bson:"label"`
	Desc  string `protobuf:"bytes,3,opt,name=Desc,proto3" json:"desc" bson:"desc"`
}

func (x *ValueOptions) Reset() {
	*x = ValueOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_tag_pb_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueOptions) ProtoMessage() {}

func (x *ValueOptions) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_tag_pb_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueOptions.ProtoReflect.Descriptor instead.
func (*ValueOptions) Descriptor() ([]byte, []int) {
	return file_pkg_tag_pb_request_proto_rawDescGZIP(), []int{3}
}

func (x *ValueOptions) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ValueOptions) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *ValueOptions) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

// HTTPFromOptions todo
type HTTPFromOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url       string            `protobuf:"bytes,1,opt,name=url,proto3" json:"url" bson:"url"`
	Headers   map[string]string `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"headers"`
	SearchKey string            `protobuf:"bytes,3,opt,name=search_key,json=searchKey,proto3" json:"search_key" bson:"search_key"`
	ValueKey  string            `protobuf:"bytes,4,opt,name=value_key,json=valueKey,proto3" json:"value_key" bson:"value_key"`
	LabelKey  string            `protobuf:"bytes,5,opt,name=label_key,json=labelKey,proto3" json:"label_key" bson:"label_key"`
	DescKey   string            `protobuf:"bytes,6,opt,name=desc_key,json=descKey,proto3" json:"desc_key" bson:"desc_key"`
}

func (x *HTTPFromOptions) Reset() {
	*x = HTTPFromOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_tag_pb_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPFromOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPFromOptions) ProtoMessage() {}

func (x *HTTPFromOptions) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_tag_pb_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPFromOptions.ProtoReflect.Descriptor instead.
func (*HTTPFromOptions) Descriptor() ([]byte, []int) {
	return file_pkg_tag_pb_request_proto_rawDescGZIP(), []int{4}
}

func (x *HTTPFromOptions) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *HTTPFromOptions) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *HTTPFromOptions) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

func (x *HTTPFromOptions) GetValueKey() string {
	if x != nil {
		return x.ValueKey
	}
	return ""
}

func (x *HTTPFromOptions) GetLabelKey() string {
	if x != nil {
		return x.LabelKey
	}
	return ""
}

func (x *HTTPFromOptions) GetDescKey() string {
	if x != nil {
		return x.DescKey
	}
	return ""
}

var File_pkg_tag_pb_request_proto protoreflect.FileDescriptor

var file_pkg_tag_pb_request_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6b, 0x65, 0x79, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x74, 0x61, 0x67, 0x1a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x78, 0x74,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x74,
	0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f,
	0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x61,
	0x67, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb9, 0x04, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x0a, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x17, 0xc2, 0xde, 0x1f, 0x13, 0x0a, 0x11, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x52, 0x09, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xc2, 0xde, 0x1f, 0x12, 0x0a, 0x10, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x6b, 0x65,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xc2, 0xde,
	0x1f, 0x11, 0x0a, 0x0f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6b, 0x65, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x09,
	0x6b, 0x65, 0x79, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x16, 0xc2, 0xde, 0x1f, 0x12, 0x0a, 0x10, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6b, 0x65, 0x79,
	0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x30, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x15, 0xc2, 0xde, 0x1f, 0x11, 0x0a, 0x0f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x22, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x44,
	0x65, 0x73, 0x63, 0x12, 0x4e, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x42,
	0x17, 0xc2, 0xde, 0x1f, 0x13, 0x0a, 0x11, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x22, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x46,
	0x72, 0x6f, 0x6d, 0x12, 0x68, 0x0a, 0x11, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x66, 0x72, 0x6f, 0x6d,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x48, 0x54, 0x54,
	0x50, 0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x1e, 0xc2, 0xde,
	0x1f, 0x1a, 0x0a, 0x18, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x52, 0x0f, 0x68, 0x74,
	0x74, 0x70, 0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4c, 0x0a,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x19, 0xc2, 0xde, 0x1f, 0x15, 0x0a, 0x13,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x89, 0x02, 0x0a, 0x0f,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x38, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x11, 0xc2, 0xde, 0x1f, 0x0d, 0x0a, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x61,
	0x67, 0x65, 0x22, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x4e, 0x0a, 0x0a, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x53, 0x63, 0x6f, 0x70,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x17, 0xc2, 0xde, 0x1f, 0x13, 0x0a, 0x11, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x52, 0x09,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xc2, 0xde, 0x1f,
	0x11, 0x0a, 0x0f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x22, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x39, 0x0a, 0x0b,
	0x73, 0x6b, 0x69, 0x70, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x42, 0x18, 0xc2, 0xde, 0x1f, 0x14, 0x0a, 0x12, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73,
	0x6b, 0x69, 0x70, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x52, 0x0a, 0x73, 0x6b, 0x69,
	0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x35, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xc2, 0xde, 0x1f, 0x0b, 0x0a,
	0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x22, 0xaf,
	0x01, 0x0a, 0x0c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x35, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1f,
	0xc2, 0xde, 0x1f, 0x1b, 0x0a, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x52,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1f, 0xc2, 0xde, 0x1f, 0x1b, 0x0a, 0x19, 0x62, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x52, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x31, 0x0a,
	0x04, 0x44, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xc2, 0xde, 0x1f,
	0x19, 0x0a, 0x17, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x22, 0x20, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x22, 0x52, 0x04, 0x44, 0x65, 0x73, 0x63,
	0x22, 0xfe, 0x03, 0x0a, 0x0f, 0x48, 0x54, 0x54, 0x50, 0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2d, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x1b, 0xc2, 0xde, 0x1f, 0x17, 0x0a, 0x15, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75,
	0x72, 0x6c, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x72, 0x6c, 0x22, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x68, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74,
	0x61, 0x67, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42,
	0x23, 0xc2, 0xde, 0x1f, 0x1f, 0x0a, 0x1d, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x22, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x48, 0x0a,
	0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x29, 0xc2, 0xde, 0x1f, 0x25, 0x0a, 0x23, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x22, 0x52, 0x09, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x44, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23,
	0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6b, 0x65,
	0x79, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6b,
	0x65, 0x79, 0x22, 0x52, 0x08, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x44, 0x0a,
	0x09, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x22, 0x52, 0x08, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x4b, 0x65, 0x79, 0x12, 0x40, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x25, 0xc2, 0xde, 0x1f, 0x21, 0x0a, 0x1f, 0x62, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x22, 0x20, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x22, 0x52, 0x07, 0x64, 0x65,
	0x73, 0x63, 0x4b, 0x65, 0x79, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_tag_pb_request_proto_rawDescOnce sync.Once
	file_pkg_tag_pb_request_proto_rawDescData = file_pkg_tag_pb_request_proto_rawDesc
)

func file_pkg_tag_pb_request_proto_rawDescGZIP() []byte {
	file_pkg_tag_pb_request_proto_rawDescOnce.Do(func() {
		file_pkg_tag_pb_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_tag_pb_request_proto_rawDescData)
	})
	return file_pkg_tag_pb_request_proto_rawDescData
}

var file_pkg_tag_pb_request_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_tag_pb_request_proto_goTypes = []interface{}{
	(*CreateTagRequest)(nil),   // 0: keyauth.tag.CreateTagRequest
	(*QueryTagRequest)(nil),    // 1: keyauth.tag.QueryTagRequest
	(*DescribeTagRequest)(nil), // 2: keyauth.tag.DescribeTagRequest
	(*ValueOptions)(nil),       // 3: keyauth.tag.ValueOptions
	(*HTTPFromOptions)(nil),    // 4: keyauth.tag.HTTPFromOptions
	nil,                        // 5: keyauth.tag.HTTPFromOptions.HeadersEntry
	(ScopeType)(0),             // 6: keyauth.tag.ScopeType
	(ValueFrom)(0),             // 7: keyauth.tag.ValueFrom
	(*page.PageRequest)(nil),   // 8: page.PageRequest
}
var file_pkg_tag_pb_request_proto_depIdxs = []int32{
	6, // 0: keyauth.tag.CreateTagRequest.scope_type:type_name -> keyauth.tag.ScopeType
	7, // 1: keyauth.tag.CreateTagRequest.value_from:type_name -> keyauth.tag.ValueFrom
	4, // 2: keyauth.tag.CreateTagRequest.http_from_options:type_name -> keyauth.tag.HTTPFromOptions
	3, // 3: keyauth.tag.CreateTagRequest.values:type_name -> keyauth.tag.ValueOptions
	8, // 4: keyauth.tag.QueryTagRequest.page:type_name -> page.PageRequest
	6, // 5: keyauth.tag.QueryTagRequest.scope_type:type_name -> keyauth.tag.ScopeType
	5, // 6: keyauth.tag.HTTPFromOptions.headers:type_name -> keyauth.tag.HTTPFromOptions.HeadersEntry
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_pkg_tag_pb_request_proto_init() }
func file_pkg_tag_pb_request_proto_init() {
	if File_pkg_tag_pb_request_proto != nil {
		return
	}
	file_pkg_tag_pb_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_tag_pb_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTagRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_tag_pb_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTagRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_tag_pb_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeTagRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_tag_pb_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_tag_pb_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPFromOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_tag_pb_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_tag_pb_request_proto_goTypes,
		DependencyIndexes: file_pkg_tag_pb_request_proto_depIdxs,
		MessageInfos:      file_pkg_tag_pb_request_proto_msgTypes,
	}.Build()
	File_pkg_tag_pb_request_proto = out.File
	file_pkg_tag_pb_request_proto_rawDesc = nil
	file_pkg_tag_pb_request_proto_goTypes = nil
	file_pkg_tag_pb_request_proto_depIdxs = nil
}