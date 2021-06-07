//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: pkg/mconf/pb/conf.proto

package mconf

import (
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

// Micro is service provider
type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 组名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"_id"`
	// 组类型
	Type Type `protobuf:"varint,2,opt,name=type,proto3,enum=keyauth.mconf.Type" json:"type" bson:"type"`
	// 创建人
	Creater string `protobuf:"bytes,3,opt,name=creater,proto3" json:"creater" bson:"creater"`
	// 创建的时间
	CreateAt int64 `protobuf:"varint,4,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 描述信息
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description" bson:"description"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_mconf_pb_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_mconf_pb_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_pkg_mconf_pb_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_GLOBAL
}

func (x *Group) GetCreater() string {
	if x != nil {
		return x.Creater
	}
	return ""
}

func (x *Group) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Group) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Item 健值项
type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 项ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 建的名称
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key" bson:"key"`
	// 关联的组
	Group string `protobuf:"bytes,3,opt,name=group,proto3" json:"group" bson:"group"`
	// 创建人
	Creater string `protobuf:"bytes,4,opt,name=creater,proto3" json:"creater" bson:"creater"`
	// 创建的时间
	CreateAt int64 `protobuf:"varint,5,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 创建的时间
	Updater int64 `protobuf:"varint,6,opt,name=updater,proto3" json:"updater,omitempty" bson:"updater"`
	// 更新时间
	UpdateAt int64 `protobuf:"varint,7,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty" bson:"update_at"`
	// 值对应的值
	Value string `protobuf:"bytes,8,opt,name=value,proto3" json:"value" bson:"value"`
	// 描述信息
	Description string `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty" bson:"description"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_mconf_pb_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_mconf_pb_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_pkg_mconf_pb_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Item) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Item) GetCreater() string {
	if x != nil {
		return x.Creater
	}
	return ""
}

func (x *Item) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Item) GetUpdater() int64 {
	if x != nil {
		return x.Updater
	}
	return 0
}

func (x *Item) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Item) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Item) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ItemSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64   `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	Items []*Item `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *ItemSet) Reset() {
	*x = ItemSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_mconf_pb_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemSet) ProtoMessage() {}

func (x *ItemSet) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_mconf_pb_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemSet.ProtoReflect.Descriptor instead.
func (*ItemSet) Descriptor() ([]byte, []int) {
	return file_pkg_mconf_pb_conf_proto_rawDescGZIP(), []int{2}
}

func (x *ItemSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ItemSet) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type GroupSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64    `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	Items []*Group `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *GroupSet) Reset() {
	*x = GroupSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_mconf_pb_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupSet) ProtoMessage() {}

func (x *GroupSet) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_mconf_pb_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupSet.ProtoReflect.Descriptor instead.
func (*GroupSet) Descriptor() ([]byte, []int) {
	return file_pkg_mconf_pb_conf_proto_rawDescGZIP(), []int{3}
}

func (x *GroupSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GroupSet) GetItems() []*Group {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_pkg_mconf_pb_conf_proto protoreflect.FileDescriptor

var file_pkg_mconf_pb_conf_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x70, 0x62, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x1a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x78,
	0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x67, 0x2f,
	0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x6d,
	0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd5, 0x02, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x30, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0xc2, 0xde, 0x1f, 0x18,
	0x0a, 0x16, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x46,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6b,
	0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x1d, 0xc2, 0xde, 0x1f, 0x19, 0x0a, 0x17, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0xc2, 0xde, 0x1f, 0x1f, 0x0a, 0x1d, 0x62,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x22, 0x20, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x22, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a, 0x21,
	0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x22,
	0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74,
	0x22, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x4d, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2b, 0xc2, 0xde, 0x1f, 0x27, 0x0a, 0x25, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc6, 0x04, 0x0a, 0x04, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x2a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x1a, 0xc2, 0xde, 0x1f, 0x16, 0x0a, 0x14, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64,
	0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x2d, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xc2, 0xde,
	0x1f, 0x17, 0x0a, 0x15, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6b, 0x65, 0x79, 0x22, 0x20, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6b, 0x65, 0x79, 0x22, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x35,
	0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1f, 0xc2,
	0xde, 0x1f, 0x1b, 0x0a, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x52, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0xc2, 0xde, 0x1f, 0x1f, 0x0a, 0x1d, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x22, 0x20, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x22, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x20,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x22,
	0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x47, 0x0a, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x2d, 0xc2, 0xde, 0x1f,
	0x29, 0x0a, 0x27, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72,
	0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x2c,
	0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x72, 0x12, 0x4e, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x42, 0x31, 0xc2, 0xde, 0x1f, 0x2d, 0x0a, 0x2b, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x20, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x2c, 0x6f,
	0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x1f, 0xc2, 0xde, 0x1f, 0x1b, 0x0a, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x57, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x35, 0xc2, 0xde, 0x1f, 0x31, 0x0a, 0x2f, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2c, 0x6f, 0x6d, 0x69, 0x74,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x8c, 0x01, 0x0a, 0x07, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x74, 0x12,
	0x35, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x1f,
	0xc2, 0xde, 0x1f, 0x1b, 0x0a, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x4a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x1f, 0xc2, 0xde, 0x1f, 0x1b,
	0x0a, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x20, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x8e, 0x01, 0x0a, 0x08, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x74, 0x12,
	0x35, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x1f,
	0xc2, 0xde, 0x1f, 0x1b, 0x0a, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x4b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x1f, 0xc2, 0xde, 0x1f,
	0x1b, 0x0a, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x20,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_mconf_pb_conf_proto_rawDescOnce sync.Once
	file_pkg_mconf_pb_conf_proto_rawDescData = file_pkg_mconf_pb_conf_proto_rawDesc
)

func file_pkg_mconf_pb_conf_proto_rawDescGZIP() []byte {
	file_pkg_mconf_pb_conf_proto_rawDescOnce.Do(func() {
		file_pkg_mconf_pb_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_mconf_pb_conf_proto_rawDescData)
	})
	return file_pkg_mconf_pb_conf_proto_rawDescData
}

var file_pkg_mconf_pb_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_mconf_pb_conf_proto_goTypes = []interface{}{
	(*Group)(nil),    // 0: keyauth.mconf.Group
	(*Item)(nil),     // 1: keyauth.mconf.Item
	(*ItemSet)(nil),  // 2: keyauth.mconf.ItemSet
	(*GroupSet)(nil), // 3: keyauth.mconf.GroupSet
	(Type)(0),        // 4: keyauth.mconf.Type
}
var file_pkg_mconf_pb_conf_proto_depIdxs = []int32{
	4, // 0: keyauth.mconf.Group.type:type_name -> keyauth.mconf.Type
	1, // 1: keyauth.mconf.ItemSet.items:type_name -> keyauth.mconf.Item
	0, // 2: keyauth.mconf.GroupSet.items:type_name -> keyauth.mconf.Group
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_mconf_pb_conf_proto_init() }
func file_pkg_mconf_pb_conf_proto_init() {
	if File_pkg_mconf_pb_conf_proto != nil {
		return
	}
	file_pkg_mconf_pb_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_mconf_pb_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_pkg_mconf_pb_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_pkg_mconf_pb_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemSet); i {
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
		file_pkg_mconf_pb_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupSet); i {
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
			RawDescriptor: file_pkg_mconf_pb_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_mconf_pb_conf_proto_goTypes,
		DependencyIndexes: file_pkg_mconf_pb_conf_proto_depIdxs,
		MessageInfos:      file_pkg_mconf_pb_conf_proto_msgTypes,
	}.Build()
	File_pkg_mconf_pb_conf_proto = out.File
	file_pkg_mconf_pb_conf_proto_rawDesc = nil
	file_pkg_mconf_pb_conf_proto_goTypes = nil
	file_pkg_mconf_pb_conf_proto_depIdxs = nil
}
