//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: pkg/tag/pb/request.proto

package tag

import (
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
	KeyName string `protobuf:"bytes,3,opt,name=key_name,json=keyName,proto3" json:"key_name" validate:"lte=200"`
	// 建标识
	KeyLabel string `protobuf:"bytes,4,opt,name=key_label,json=keyLabel,proto3" json:"key_label"`
	// 建描述
	KeyDesc string `protobuf:"bytes,5,opt,name=key_desc,json=keyDesc,proto3" json:"key_desc"`
	// 值来源
	ValueFrom ValueFrom `protobuf:"varint,6,opt,name=value_from,json=valueFrom,proto3,enum=keyauth.tag.ValueFrom" json:"value_from"`
	// http 获取Tag 值的参数
	HttpFromOption *HTTPFromOption `protobuf:"bytes,7,opt,name=http_from_option,json=httpFromOption,proto3" json:"http_from_option"`
	// String 类型的值
	Values []*ValueOption `protobuf:"bytes,8,rep,name=values,proto3" json:"values"`
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

func (x *CreateTagRequest) GetHttpFromOption() *HTTPFromOption {
	if x != nil {
		return x.HttpFromOption
	}
	return nil
}

func (x *CreateTagRequest) GetValues() []*ValueOption {
	if x != nil {
		return x.Values
	}
	return nil
}

// QueryTagKeyRequest todo
type QueryTagKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *page.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// tag生效范围
	ScopeType ScopeType `protobuf:"varint,2,opt,name=scope_type,json=scopeType,proto3,enum=keyauth.tag.ScopeType" json:"scope_type"`
	// 关键字
	Keywords string `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords"`
}

func (x *QueryTagKeyRequest) Reset() {
	*x = QueryTagKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_tag_pb_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTagKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTagKeyRequest) ProtoMessage() {}

func (x *QueryTagKeyRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use QueryTagKeyRequest.ProtoReflect.Descriptor instead.
func (*QueryTagKeyRequest) Descriptor() ([]byte, []int) {
	return file_pkg_tag_pb_request_proto_rawDescGZIP(), []int{1}
}

func (x *QueryTagKeyRequest) GetPage() *page.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryTagKeyRequest) GetScopeType() ScopeType {
	if x != nil {
		return x.ScopeType
	}
	return ScopeType_NAMESPACE
}

func (x *QueryTagKeyRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

// QueryTagValueRequest todo
type QueryTagValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *page.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// Tag Value ID
	TagId string `protobuf:"bytes,2,opt,name=tag_id,json=tagId,proto3" json:"tag_id"`
}

func (x *QueryTagValueRequest) Reset() {
	*x = QueryTagValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_tag_pb_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTagValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTagValueRequest) ProtoMessage() {}

func (x *QueryTagValueRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use QueryTagValueRequest.ProtoReflect.Descriptor instead.
func (*QueryTagValueRequest) Descriptor() ([]byte, []int) {
	return file_pkg_tag_pb_request_proto_rawDescGZIP(), []int{2}
}

func (x *QueryTagValueRequest) GetPage() *page.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryTagValueRequest) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

// DeleteTagRequest todo
type DeleteTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tag Value ID
	TagId string `protobuf:"bytes,2,opt,name=tag_id,json=tagId,proto3" json:"tag_id" validate:"lte=200"`
}

func (x *DeleteTagRequest) Reset() {
	*x = DeleteTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_tag_pb_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTagRequest) ProtoMessage() {}

func (x *DeleteTagRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeleteTagRequest.ProtoReflect.Descriptor instead.
func (*DeleteTagRequest) Descriptor() ([]byte, []int) {
	return file_pkg_tag_pb_request_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteTagRequest) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

type DescribeTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tag Value ID
	TagId string `protobuf:"bytes,2,opt,name=tag_id,json=tagId,proto3" json:"tag_id" validate:"lte=200"`
}

func (x *DescribeTagRequest) Reset() {
	*x = DescribeTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_tag_pb_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeTagRequest) ProtoMessage() {}

func (x *DescribeTagRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DescribeTagRequest.ProtoReflect.Descriptor instead.
func (*DescribeTagRequest) Descriptor() ([]byte, []int) {
	return file_pkg_tag_pb_request_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeTagRequest) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

// ValueOptions 值描述
type ValueOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=Value,proto3" json:"value" bson:"value" validate:"lte=200"`
	Label string `protobuf:"bytes,2,opt,name=Label,proto3" json:"label" bson:"label"`
	Desc  string `protobuf:"bytes,3,opt,name=Desc,proto3" json:"desc" bson:"desc"`
}

func (x *ValueOption) Reset() {
	*x = ValueOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_tag_pb_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueOption) ProtoMessage() {}

func (x *ValueOption) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_tag_pb_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueOption.ProtoReflect.Descriptor instead.
func (*ValueOption) Descriptor() ([]byte, []int) {
	return file_pkg_tag_pb_request_proto_rawDescGZIP(), []int{5}
}

func (x *ValueOption) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ValueOption) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *ValueOption) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

// HTTPFromOptions todo
type HTTPFromOption struct {
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

func (x *HTTPFromOption) Reset() {
	*x = HTTPFromOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_tag_pb_request_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPFromOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPFromOption) ProtoMessage() {}

func (x *HTTPFromOption) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_tag_pb_request_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPFromOption.ProtoReflect.Descriptor instead.
func (*HTTPFromOption) Descriptor() ([]byte, []int) {
	return file_pkg_tag_pb_request_proto_rawDescGZIP(), []int{6}
}

func (x *HTTPFromOption) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *HTTPFromOption) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *HTTPFromOption) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

func (x *HTTPFromOption) GetValueKey() string {
	if x != nil {
		return x.ValueKey
	}
	return ""
}

func (x *HTTPFromOption) GetLabelKey() string {
	if x != nil {
		return x.LabelKey
	}
	return ""
}

func (x *HTTPFromOption) GetDescKey() string {
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
	0xc1, 0x04, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x0a, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x17, 0xc2, 0xde, 0x1f, 0x13, 0x0a, 0x11, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x52, 0x09, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xc2, 0xde, 0x1f, 0x12, 0x0a, 0x10, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x6b, 0x65,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xc2, 0xde,
	0x1f, 0x24, 0x0a, 0x22, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6b, 0x65, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x6c, 0x74,
	0x65, 0x3d, 0x32, 0x30, 0x30, 0x22, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x33, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x16, 0xc2, 0xde, 0x1f, 0x12, 0x0a, 0x10, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x6b, 0x65, 0x79, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x30, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xc2, 0xde, 0x1f, 0x11, 0x0a, 0x0f, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x22, 0x52, 0x07, 0x6b,
	0x65, 0x79, 0x44, 0x65, 0x73, 0x63, 0x12, 0x4e, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6b, 0x65, 0x79,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x42, 0x17, 0xc2, 0xde, 0x1f, 0x13, 0x0a, 0x11, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x22, 0x52, 0x09, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x64, 0x0a, 0x10, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x48,
	0x54, 0x54, 0x50, 0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x1d, 0xc2,
	0xde, 0x1f, 0x19, 0x0a, 0x17, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x52, 0x0e, 0x68, 0x74,
	0x74, 0x70, 0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b,
	0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x13, 0xc2, 0xde, 0x1f, 0x0f, 0x0a, 0x0d, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x52, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x22, 0xd1, 0x01, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x67,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x11, 0xc2, 0xde, 0x1f,
	0x0d, 0x0a, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x61, 0x67, 0x65, 0x22, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x4e, 0x0a, 0x0a, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x17, 0xc2, 0xde, 0x1f, 0x13, 0x0a, 0x11, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x52, 0x09, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xc2, 0xde, 0x1f, 0x11, 0x0a, 0x0f, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x52, 0x08, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x7c, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x54, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x38, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x11, 0xc2, 0xde, 0x1f, 0x0d, 0x0a, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x61,
	0x67, 0x65, 0x22, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x74, 0x61, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xc2, 0xde, 0x1f, 0x0f, 0x0a,
	0x0d, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x22, 0x52, 0x05,
	0x74, 0x61, 0x67, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x06, 0x74, 0x61, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xc2, 0xde, 0x1f, 0x22, 0x0a,
	0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x6c, 0x74, 0x65, 0x3d, 0x32, 0x30, 0x30,
	0x22, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d,
	0x0a, 0x06, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26,
	0xc2, 0xde, 0x1f, 0x22, 0x0a, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x61, 0x67, 0x5f,
	0x69, 0x64, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x6c, 0x74,
	0x65, 0x3d, 0x32, 0x30, 0x30, 0x22, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x22, 0xc1, 0x01,
	0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x32, 0xc2, 0xde,
	0x1f, 0x2e, 0x0a, 0x2c, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x20, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x6c, 0x74, 0x65, 0x3d, 0x32, 0x30, 0x30, 0x22,
	0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1f, 0xc2, 0xde, 0x1f, 0x1b, 0x0a, 0x19, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x52, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x31,
	0x0a, 0x04, 0x44, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xc2, 0xde,
	0x1f, 0x19, 0x0a, 0x17, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x22, 0x20,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x22, 0x52, 0x04, 0x44, 0x65, 0x73,
	0x63, 0x22, 0xfc, 0x03, 0x0a, 0x0e, 0x48, 0x54, 0x54, 0x50, 0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x1b, 0xc2, 0xde, 0x1f, 0x17, 0x0a, 0x15, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75,
	0x72, 0x6c, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x72, 0x6c, 0x22, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x67, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74,
	0x61, 0x67, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x23,
	0xc2, 0xde, 0x1f, 0x1f, 0x0a, 0x1d, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x22, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x48, 0x0a, 0x0a,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x29, 0xc2, 0xde, 0x1f, 0x25, 0x0a, 0x23, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x22, 0x52, 0x09, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x44, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a,
	0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6b, 0x65, 0x79,
	0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6b, 0x65,
	0x79, 0x22, 0x52, 0x08, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x44, 0x0a, 0x09,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x22, 0x52, 0x08, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x4b,
	0x65, 0x79, 0x12, 0x40, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x25, 0xc2, 0xde, 0x1f, 0x21, 0x0a, 0x1f, 0x62, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x22, 0x52, 0x07, 0x64, 0x65, 0x73,
	0x63, 0x4b, 0x65, 0x79, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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

var file_pkg_tag_pb_request_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_tag_pb_request_proto_goTypes = []interface{}{
	(*CreateTagRequest)(nil),     // 0: keyauth.tag.CreateTagRequest
	(*QueryTagKeyRequest)(nil),   // 1: keyauth.tag.QueryTagKeyRequest
	(*QueryTagValueRequest)(nil), // 2: keyauth.tag.QueryTagValueRequest
	(*DeleteTagRequest)(nil),     // 3: keyauth.tag.DeleteTagRequest
	(*DescribeTagRequest)(nil),   // 4: keyauth.tag.DescribeTagRequest
	(*ValueOption)(nil),          // 5: keyauth.tag.ValueOption
	(*HTTPFromOption)(nil),       // 6: keyauth.tag.HTTPFromOption
	nil,                          // 7: keyauth.tag.HTTPFromOption.HeadersEntry
	(ScopeType)(0),               // 8: keyauth.tag.ScopeType
	(ValueFrom)(0),               // 9: keyauth.tag.ValueFrom
	(*page.PageRequest)(nil),     // 10: page.PageRequest
}
var file_pkg_tag_pb_request_proto_depIdxs = []int32{
	8,  // 0: keyauth.tag.CreateTagRequest.scope_type:type_name -> keyauth.tag.ScopeType
	9,  // 1: keyauth.tag.CreateTagRequest.value_from:type_name -> keyauth.tag.ValueFrom
	6,  // 2: keyauth.tag.CreateTagRequest.http_from_option:type_name -> keyauth.tag.HTTPFromOption
	5,  // 3: keyauth.tag.CreateTagRequest.values:type_name -> keyauth.tag.ValueOption
	10, // 4: keyauth.tag.QueryTagKeyRequest.page:type_name -> page.PageRequest
	8,  // 5: keyauth.tag.QueryTagKeyRequest.scope_type:type_name -> keyauth.tag.ScopeType
	10, // 6: keyauth.tag.QueryTagValueRequest.page:type_name -> page.PageRequest
	7,  // 7: keyauth.tag.HTTPFromOption.headers:type_name -> keyauth.tag.HTTPFromOption.HeadersEntry
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
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
			switch v := v.(*QueryTagKeyRequest); i {
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
			switch v := v.(*QueryTagValueRequest); i {
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
			switch v := v.(*DeleteTagRequest); i {
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
		file_pkg_tag_pb_request_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueOption); i {
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
		file_pkg_tag_pb_request_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPFromOption); i {
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
			NumMessages:   8,
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
