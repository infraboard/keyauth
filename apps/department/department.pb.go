// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: apps/department/pb/department.proto

package department

import (
	role "github.com/infraboard/keyauth/apps/role"
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

// Department user's department
type Department struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 部门ID
	// @gotags: bson:"_id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 路径
	// @gotags: bson:"parent_path" json:"parent_path"
	ParentPath string `protobuf:"bytes,2,opt,name=parent_path,json=parentPath,proto3" json:"parent_path" bson:"parent_path"`
	// 部门编号
	// @gotags: bson:"number" json:"number"
	Number uint64 `protobuf:"varint,3,opt,name=number,proto3" json:"number" bson:"number"`
	// 部门创建时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,4,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 更新时间
	// @gotags: bson:"update_at" json:"update_at"
	UpdateAt int64 `protobuf:"varint,5,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 创建人
	// @gotags: bson:"create_by" json:"create_by"
	CreateBy string `protobuf:"bytes,6,opt,name=create_by,json=createBy,proto3" json:"create_by" bson:"create_by"`
	// 部门所属域
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,7,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 第几级部门, 由层数决定
	// @gotags: bson:"grade" json:"grade"
	Grade int32 `protobuf:"varint,8,opt,name=grade,proto3" json:"grade" bson:"grade"`
	// 子部门数量
	// @gotags: bson:"-" json:"sub_count"
	SubCount int64 `protobuf:"varint,9,opt,name=sub_count,json=subCount,proto3" json:"sub_count" bson:"-"`
	// 部门所有用户数量
	// @gotags: bson:"-" json:"user_count"
	UserCount int64 `protobuf:"varint,10,opt,name=user_count,json=userCount,proto3" json:"user_count" bson:"-"`
	// 部门名称
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,11,opt,name=name,proto3" json:"name" bson:"name"`
	// 显示名称
	// @gotags: bson:"display_name" json:"display_name"
	DisplayName string `protobuf:"bytes,12,opt,name=display_name,json=displayName,proto3" json:"display_name" bson:"display_name"`
	// 上级部门ID
	// @gotags: bson:"parent_id" json:"parent_id"
	ParentId string `protobuf:"bytes,13,opt,name=parent_id,json=parentId,proto3" json:"parent_id" bson:"parent_id"`
	// 部门管理者account
	// @gotags: bson:"manager" json:"manager"
	Manager string `protobuf:"bytes,14,opt,name=manager,proto3" json:"manager" bson:"manager"`
	// 部门成员默认角色
	// @gotags: bson:"default_role_id" json:"default_role_id"
	DefaultRoleId string `protobuf:"bytes,15,opt,name=default_role_id,json=defaultRoleId,proto3" json:"default_role_id" bson:"default_role_id"`
	// 默认角色
	// @gotags: bson:"-" json:"default_role,omitempty"
	DefaultRole *role.Role `protobuf:"bytes,16,opt,name=default_role,json=defaultRole,proto3" json:"default_role,omitempty" bson:"-"`
}

func (x *Department) Reset() {
	*x = Department{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_department_pb_department_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Department) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Department) ProtoMessage() {}

func (x *Department) ProtoReflect() protoreflect.Message {
	mi := &file_apps_department_pb_department_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Department.ProtoReflect.Descriptor instead.
func (*Department) Descriptor() ([]byte, []int) {
	return file_apps_department_pb_department_proto_rawDescGZIP(), []int{0}
}

func (x *Department) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Department) GetParentPath() string {
	if x != nil {
		return x.ParentPath
	}
	return ""
}

func (x *Department) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Department) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Department) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Department) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *Department) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Department) GetGrade() int32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *Department) GetSubCount() int64 {
	if x != nil {
		return x.SubCount
	}
	return 0
}

func (x *Department) GetUserCount() int64 {
	if x != nil {
		return x.UserCount
	}
	return 0
}

func (x *Department) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Department) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Department) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *Department) GetManager() string {
	if x != nil {
		return x.Manager
	}
	return ""
}

func (x *Department) GetDefaultRoleId() string {
	if x != nil {
		return x.DefaultRoleId
	}
	return ""
}

func (x *Department) GetDefaultRole() *role.Role {
	if x != nil {
		return x.DefaultRole
	}
	return nil
}

type Set struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"items" json:"items"
	Items []*Department `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *Set) Reset() {
	*x = Set{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_department_pb_department_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_apps_department_pb_department_proto_msgTypes[1]
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
	return file_apps_department_pb_department_proto_rawDescGZIP(), []int{1}
}

func (x *Set) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Set) GetItems() []*Department {
	if x != nil {
		return x.Items
	}
	return nil
}

// ApplicationForm todo
type ApplicationForm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 申请单ID
	// @gotags: bson:"_id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 域
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 申请人
	// @gotags: bson:"creater" json:"creater"
	Creater string `protobuf:"bytes,3,opt,name=creater,proto3" json:"creater" bson:"creater"`
	// 创建时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,4,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 更新时间
	// @gotags: bson:"update_at" json:"update_at"
	UpdateAt int64 `protobuf:"varint,5,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 申请人
	// @gotags: bson:"account" json:"account"
	Account string `protobuf:"bytes,6,opt,name=account,proto3" json:"account" bson:"account"`
	// 申请加入的部门
	// @gotags: bson:"department_id" json:"department_id"
	DepartmentId string `protobuf:"bytes,7,opt,name=department_id,json=departmentId,proto3" json:"department_id" bson:"department_id"`
	// 留言
	// @gotags: bson:"message" json:"message"
	Message string `protobuf:"bytes,8,opt,name=message,proto3" json:"message" bson:"message"`
	// 状态
	// @gotags: bson:"status" json:"status"
	Status ApplicationFormStatus `protobuf:"varint,9,opt,name=status,proto3,enum=infraboard.keyauth.department.ApplicationFormStatus" json:"status" bson:"status"`
}

func (x *ApplicationForm) Reset() {
	*x = ApplicationForm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_department_pb_department_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationForm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationForm) ProtoMessage() {}

func (x *ApplicationForm) ProtoReflect() protoreflect.Message {
	mi := &file_apps_department_pb_department_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationForm.ProtoReflect.Descriptor instead.
func (*ApplicationForm) Descriptor() ([]byte, []int) {
	return file_apps_department_pb_department_proto_rawDescGZIP(), []int{2}
}

func (x *ApplicationForm) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ApplicationForm) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *ApplicationForm) GetCreater() string {
	if x != nil {
		return x.Creater
	}
	return ""
}

func (x *ApplicationForm) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *ApplicationForm) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *ApplicationForm) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *ApplicationForm) GetDepartmentId() string {
	if x != nil {
		return x.DepartmentId
	}
	return ""
}

func (x *ApplicationForm) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ApplicationForm) GetStatus() ApplicationFormStatus {
	if x != nil {
		return x.Status
	}
	return ApplicationFormStatus_NULL
}

type ApplicationFormSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"items" json:"items"
	Items []*ApplicationForm `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *ApplicationFormSet) Reset() {
	*x = ApplicationFormSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_department_pb_department_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationFormSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationFormSet) ProtoMessage() {}

func (x *ApplicationFormSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_department_pb_department_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationFormSet.ProtoReflect.Descriptor instead.
func (*ApplicationFormSet) Descriptor() ([]byte, []int) {
	return file_apps_department_pb_department_proto_rawDescGZIP(), []int{3}
}

func (x *ApplicationFormSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ApplicationFormSet) GetItems() []*ApplicationForm {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_apps_department_pb_department_proto protoreflect.FileDescriptor

var file_apps_department_pb_department_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1d, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70,
	0x62, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x03, 0x0a,
	0x0a, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x75,
	0x62, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x0c, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b,
	0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x5c, 0x0a,
	0x03, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x3f, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xb4, 0x02, 0x0a, 0x0f,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x72, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x4c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x6f, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x70, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x44,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65,
	0x79, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_department_pb_department_proto_rawDescOnce sync.Once
	file_apps_department_pb_department_proto_rawDescData = file_apps_department_pb_department_proto_rawDesc
)

func file_apps_department_pb_department_proto_rawDescGZIP() []byte {
	file_apps_department_pb_department_proto_rawDescOnce.Do(func() {
		file_apps_department_pb_department_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_department_pb_department_proto_rawDescData)
	})
	return file_apps_department_pb_department_proto_rawDescData
}

var file_apps_department_pb_department_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_department_pb_department_proto_goTypes = []interface{}{
	(*Department)(nil),         // 0: infraboard.keyauth.department.Department
	(*Set)(nil),                // 1: infraboard.keyauth.department.Set
	(*ApplicationForm)(nil),    // 2: infraboard.keyauth.department.ApplicationForm
	(*ApplicationFormSet)(nil), // 3: infraboard.keyauth.department.ApplicationFormSet
	(*role.Role)(nil),          // 4: infraboard.keyauth.role.Role
	(ApplicationFormStatus)(0), // 5: infraboard.keyauth.department.ApplicationFormStatus
}
var file_apps_department_pb_department_proto_depIdxs = []int32{
	4, // 0: infraboard.keyauth.department.Department.default_role:type_name -> infraboard.keyauth.role.Role
	0, // 1: infraboard.keyauth.department.Set.items:type_name -> infraboard.keyauth.department.Department
	5, // 2: infraboard.keyauth.department.ApplicationForm.status:type_name -> infraboard.keyauth.department.ApplicationFormStatus
	2, // 3: infraboard.keyauth.department.ApplicationFormSet.items:type_name -> infraboard.keyauth.department.ApplicationForm
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_apps_department_pb_department_proto_init() }
func file_apps_department_pb_department_proto_init() {
	if File_apps_department_pb_department_proto != nil {
		return
	}
	file_apps_department_pb_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_department_pb_department_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Department); i {
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
		file_apps_department_pb_department_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_apps_department_pb_department_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationForm); i {
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
		file_apps_department_pb_department_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationFormSet); i {
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
			RawDescriptor: file_apps_department_pb_department_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_department_pb_department_proto_goTypes,
		DependencyIndexes: file_apps_department_pb_department_proto_depIdxs,
		MessageInfos:      file_apps_department_pb_department_proto_msgTypes,
	}.Build()
	File_apps_department_pb_department_proto = out.File
	file_apps_department_pb_department_proto_rawDesc = nil
	file_apps_department_pb_department_proto_goTypes = nil
	file_apps_department_pb_department_proto_depIdxs = nil
}