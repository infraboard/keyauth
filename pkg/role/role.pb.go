//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: pkg/role/pb/role.proto

package role

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

// Role is rbac's role
type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 角色ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 创建时间`
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 更新时间
	UpdateAt int64 `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 角色所属域
	Domain string `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 创建人
	Creater string `protobuf:"bytes,5,opt,name=creater,proto3" json:"creater" bson:"creater"`
	// 角色类型
	Type RoleType `protobuf:"varint,6,opt,name=type,proto3,enum=keyauth.role.RoleType" json:"type" bson:"type"`
	// 应用名称
	Name string `protobuf:"bytes,7,opt,name=name,proto3" json:"name" bson:"name"`
	// 应用简单的描述
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description" bson:"description"`
	// 角色关联的page
	PageMarked string `protobuf:"bytes,10,opt,name=page_marked,json=pageMarked,proto3" json:"page_marked" bson:"page_marked"`
	// 读权限
	Permissions []*Permission `protobuf:"bytes,9,rep,name=permissions,proto3" json:"permissions,omitempty" bson:"-"`
	// 范围, 角色范围限制, 由策略引擎动态补充
	Scope string `protobuf:"bytes,11,opt,name=scope,proto3" json:"scope" bson:"-"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_role_pb_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_role_pb_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_pkg_role_pb_role_proto_rawDescGZIP(), []int{0}
}

func (x *Role) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Role) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Role) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Role) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Role) GetCreater() string {
	if x != nil {
		return x.Creater
	}
	return ""
}

func (x *Role) GetType() RoleType {
	if x != nil {
		return x.Type
	}
	return RoleType_NULL
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Role) GetPageMarked() string {
	if x != nil {
		return x.PageMarked
	}
	return ""
}

func (x *Role) GetPermissions() []*Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *Role) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

// Permission 权限
type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 关联角色ID
	RoleId string `protobuf:"bytes,2,opt,name=role_id,json=roleId,proto3" json:"role_id" bson:"role_id"`
	// 创建时间
	CreateAt int64 `protobuf:"varint,3,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 创建人
	Creater string `protobuf:"bytes,4,opt,name=creater,proto3" json:"creater" bson:"creater"`
	// 效力
	Effect EffectType `protobuf:"varint,5,opt,name=effect,proto3,enum=keyauth.role.EffectType" json:"effect" bson:"effect"`
	// 服务ID
	ServiceId string `protobuf:"bytes,6,opt,name=service_id,json=serviceId,proto3" json:"service_id" bson:"service_id"`
	// 资源列表
	ResourceName string `protobuf:"bytes,7,opt,name=resource_name,json=resourceName,proto3" json:"resource_name" bson:"resource_name"`
	// 维度
	LabelKey string `protobuf:"bytes,8,opt,name=label_key,json=labelKey,proto3" json:"label_key" bson:"label_key"`
	// 适配所有值
	MatchAll bool `protobuf:"varint,9,opt,name=match_all,json=matchAll,proto3" json:"match_all" bson:"match_all"`
	// 标识值
	LabelValues []string `protobuf:"bytes,10,rep,name=label_values,json=labelValues,proto3" json:"label_values" bson:"label_values"`
	// 范围, 角色范围限制, 由策略引擎动态补充
	Scope string `protobuf:"bytes,11,opt,name=scope,proto3" json:"scope" bson:"-"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_role_pb_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_role_pb_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_pkg_role_pb_role_proto_rawDescGZIP(), []int{1}
}

func (x *Permission) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Permission) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *Permission) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Permission) GetCreater() string {
	if x != nil {
		return x.Creater
	}
	return ""
}

func (x *Permission) GetEffect() EffectType {
	if x != nil {
		return x.Effect
	}
	return EffectType_ALLOW
}

func (x *Permission) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *Permission) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *Permission) GetLabelKey() string {
	if x != nil {
		return x.LabelKey
	}
	return ""
}

func (x *Permission) GetMatchAll() bool {
	if x != nil {
		return x.MatchAll
	}
	return false
}

func (x *Permission) GetLabelValues() []string {
	if x != nil {
		return x.LabelValues
	}
	return nil
}

func (x *Permission) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

type Set struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64   `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	Items []*Role `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *Set) Reset() {
	*x = Set{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_role_pb_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_role_pb_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Set.ProtoReflect.Descriptor instead.
func (*Set) Descriptor() ([]byte, []int) {
	return file_pkg_role_pb_role_proto_rawDescGZIP(), []int{2}
}

func (x *Set) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Set) GetItems() []*Role {
	if x != nil {
		return x.Items
	}
	return nil
}

// PermissionSet 用户列表
type PermissionSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64         `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	Items []*Permission `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *PermissionSet) Reset() {
	*x = PermissionSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_role_pb_role_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionSet) ProtoMessage() {}

func (x *PermissionSet) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_role_pb_role_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionSet.ProtoReflect.Descriptor instead.
func (*PermissionSet) Descriptor() ([]byte, []int) {
	return file_pkg_role_pb_role_proto_rawDescGZIP(), []int{3}
}

func (x *PermissionSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PermissionSet) GetItems() []*Permission {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_pkg_role_pb_role_proto protoreflect.FileDescriptor

var file_pkg_role_pb_role_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x1a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x78, 0x74, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x74, 0x61,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x6f, 0x6c,
	0x65, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xef, 0x05, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0xc2, 0xde, 0x1f, 0x16, 0x0a, 0x14, 0x62, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x64, 0x22,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x44, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x20,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x22,
	0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x44, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x27, 0xc2,
	0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x21, 0xc2, 0xde, 0x1f, 0x1d, 0x0a, 0x1b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x22, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x3d, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0xc2, 0xde,
	0x1f, 0x1f, 0x0a, 0x1d, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x72, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72,
	0x22, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x1d, 0xc2, 0xde, 0x1f, 0x19, 0x0a, 0x17, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x1d, 0xc2, 0xde, 0x1f, 0x19, 0x0a, 0x17, 0x62, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xc2,
	0xde, 0x1f, 0x27, 0x0a, 0x25, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xc2, 0xde,
	0x1f, 0x27, 0x0a, 0x25, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x64, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x22, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x64, 0x12, 0x67, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x65, 0x79,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x42, 0x2b, 0xc2, 0xde, 0x1f, 0x27, 0x0a, 0x25, 0x62, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x2d, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x31,
	0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xc2,
	0xde, 0x1f, 0x17, 0x0a, 0x15, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2d, 0x22, 0x20, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x22, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x22, 0x81, 0x06, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x2a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0xc2, 0xde,
	0x1f, 0x16, 0x0a, 0x14, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x07,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0xc2,
	0xde, 0x1f, 0x1f, 0x0a, 0x1d, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x6f, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x22, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x27, 0xc2,
	0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x12, 0x3d, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x23, 0xc2, 0xde, 0x1f, 0x1f, 0x0a, 0x1d, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x72, 0x22, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x12,
	0x53, 0x0a, 0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x18, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x45,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x21, 0xc2, 0xde, 0x1f, 0x1d, 0x0a,
	0x1b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x22, 0x20, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x22, 0x52, 0x06, 0x65, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x12, 0x48, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xc2, 0xde, 0x1f, 0x25, 0x0a, 0x23,
	0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x22, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x54,
	0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xc2, 0xde, 0x1f, 0x2b, 0x0a, 0x29, 0x62, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x22, 0x20,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x22,
	0x52, 0x08, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x4b, 0x65, 0x79, 0x12, 0x44, 0x0a, 0x09, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x42, 0x27, 0xc2,
	0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x5f, 0x61, 0x6c, 0x6c, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x5f, 0x61, 0x6c, 0x6c, 0x22, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x41, 0x6c, 0x6c,
	0x12, 0x50, 0x0a, 0x0c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x42, 0x2d, 0xc2, 0xde, 0x1f, 0x29, 0x0a, 0x27, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x22, 0x52, 0x0b, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x1b, 0xc2, 0xde, 0x1f, 0x17, 0x0a, 0x15, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2d,
	0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x22, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x35, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x1f, 0xc2, 0xde,
	0x1f, 0x1b, 0x0a, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22,
	0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x49, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x1f, 0xc2, 0xde, 0x1f, 0x1b, 0x0a, 0x19, 0x62,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0x97, 0x01, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x74, 0x12, 0x35, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x1f, 0xc2, 0xde, 0x1f, 0x1b, 0x0a, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x4f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x42, 0x1f, 0xc2, 0xde, 0x1f, 0x1b, 0x0a, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_role_pb_role_proto_rawDescOnce sync.Once
	file_pkg_role_pb_role_proto_rawDescData = file_pkg_role_pb_role_proto_rawDesc
)

func file_pkg_role_pb_role_proto_rawDescGZIP() []byte {
	file_pkg_role_pb_role_proto_rawDescOnce.Do(func() {
		file_pkg_role_pb_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_role_pb_role_proto_rawDescData)
	})
	return file_pkg_role_pb_role_proto_rawDescData
}

var file_pkg_role_pb_role_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_role_pb_role_proto_goTypes = []interface{}{
	(*Role)(nil),          // 0: keyauth.role.Role
	(*Permission)(nil),    // 1: keyauth.role.Permission
	(*Set)(nil),           // 2: keyauth.role.Set
	(*PermissionSet)(nil), // 3: keyauth.role.PermissionSet
	(RoleType)(0),         // 4: keyauth.role.RoleType
	(EffectType)(0),       // 5: keyauth.role.EffectType
}
var file_pkg_role_pb_role_proto_depIdxs = []int32{
	4, // 0: keyauth.role.Role.type:type_name -> keyauth.role.RoleType
	1, // 1: keyauth.role.Role.permissions:type_name -> keyauth.role.Permission
	5, // 2: keyauth.role.Permission.effect:type_name -> keyauth.role.EffectType
	0, // 3: keyauth.role.Set.items:type_name -> keyauth.role.Role
	1, // 4: keyauth.role.PermissionSet.items:type_name -> keyauth.role.Permission
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_role_pb_role_proto_init() }
func file_pkg_role_pb_role_proto_init() {
	if File_pkg_role_pb_role_proto != nil {
		return
	}
	file_pkg_role_pb_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_role_pb_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_pkg_role_pb_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
		file_pkg_role_pb_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Set); i {
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
		file_pkg_role_pb_role_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionSet); i {
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
			RawDescriptor: file_pkg_role_pb_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_role_pb_role_proto_goTypes,
		DependencyIndexes: file_pkg_role_pb_role_proto_depIdxs,
		MessageInfos:      file_pkg_role_pb_role_proto_msgTypes,
	}.Build()
	File_pkg_role_pb_role_proto = out.File
	file_pkg_role_pb_role_proto_rawDesc = nil
	file_pkg_role_pb_role_proto_goTypes = nil
	file_pkg_role_pb_role_proto_depIdxs = nil
}
