// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: apps/domain/pb/request.proto

package domain

import (
	types "github.com/infraboard/keyauth/common/types"
	request "github.com/infraboard/mcube/http/request"
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

type CreateDomainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 公司或者组织名称
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"name"`
	// Profile 需要修改内容
	// @gotags: bson:"profile" json:"profile"
	Profile *DomainProfile `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile" bson:"profile"`
	// 归属人
	// @gotags: bson:"owner" json:"owner"
	Owner string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner" bson:"owner"`
}

func (x *CreateDomainRequest) Reset() {
	*x = CreateDomainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_domain_pb_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDomainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDomainRequest) ProtoMessage() {}

func (x *CreateDomainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_domain_pb_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDomainRequest.ProtoReflect.Descriptor instead.
func (*CreateDomainRequest) Descriptor() ([]byte, []int) {
	return file_apps_domain_pb_request_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDomainRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDomainRequest) GetProfile() *DomainProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *CreateDomainRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

type DomainProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 全称
	// @gotags: bson:"display_name" json:"display_name"
	DisplayName string `protobuf:"bytes,5,opt,name=display_name,json=displayName,proto3" json:"display_name" bson:"display_name"`
	// 公司LOGO图片的URL
	// @gotags: bson:"logo_path" json:"logo_path"
	LogoPath string `protobuf:"bytes,6,opt,name=logo_path,json=logoPath,proto3" json:"logo_path" bson:"logo_path"`
	// 描述
	// @gotags: bson:"description" json:"description"
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description" bson:"description"`
	// 电话
	// @gotags: bson:"phone" json:"phone"
	Phone string `protobuf:"bytes,9,opt,name=phone,proto3" json:"phone" bson:"phone"`
	// 规模: 50人以下, 50~100, ...
	// @gotags: bson:"size" json:"size"
	Size string `protobuf:"bytes,10,opt,name=size,proto3" json:"size" bson:"size"`
	// 位置: 指城市, 比如 中国,四川,成都
	// @gotags: bson:"location" json:"location"
	Location string `protobuf:"bytes,11,opt,name=location,proto3" json:"location" bson:"location"`
	// 地址: 比如环球中心 10F 1034
	// @gotags: bson:"address" json:"address"
	Address string `protobuf:"bytes,12,opt,name=address,proto3" json:"address" bson:"address"`
	// 所属行业: 比如, 互联网
	// @gotags: bson:"industry" json:"industry"
	Industry string `protobuf:"bytes,13,opt,name=industry,proto3" json:"industry" bson:"industry"`
	// 传真
	// @gotags: bson:"fax" json:"fax"
	Fax string `protobuf:"bytes,14,opt,name=fax,proto3" json:"fax" bson:"fax"`
	// 联系人
	// @gotags: bson:"contack" json:"contack"
	Contack *Contact `protobuf:"bytes,15,opt,name=contack,proto3" json:"contack" bson:"contack"`
}

func (x *DomainProfile) Reset() {
	*x = DomainProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_domain_pb_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainProfile) ProtoMessage() {}

func (x *DomainProfile) ProtoReflect() protoreflect.Message {
	mi := &file_apps_domain_pb_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainProfile.ProtoReflect.Descriptor instead.
func (*DomainProfile) Descriptor() ([]byte, []int) {
	return file_apps_domain_pb_request_proto_rawDescGZIP(), []int{1}
}

func (x *DomainProfile) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *DomainProfile) GetLogoPath() string {
	if x != nil {
		return x.LogoPath
	}
	return ""
}

func (x *DomainProfile) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DomainProfile) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *DomainProfile) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *DomainProfile) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *DomainProfile) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DomainProfile) GetIndustry() string {
	if x != nil {
		return x.Industry
	}
	return ""
}

func (x *DomainProfile) GetFax() string {
	if x != nil {
		return x.Fax
	}
	return ""
}

func (x *DomainProfile) GetContack() *Contact {
	if x != nil {
		return x.Contack
	}
	return nil
}

type UpdateDomainInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 更新模式
	// @gotags: bson:"update_mode" json:"update_mode"
	UpdateMode types.UpdateMode `protobuf:"varint,1,opt,name=update_mode,json=updateMode,proto3,enum=infraboard.keyauth.types.UpdateMode" json:"update_mode" bson:"update_mode"`
	// 公司或者组织名称
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" bson:"name"`
	// CreateDomainRequest 需要修改内容
	// @gotags: bson:"profile" json:"profile"
	Profile *DomainProfile `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile" bson:"profile"`
	// 修改者
	// @gotags: bson:"account" json:"account"
	Account string `protobuf:"bytes,4,opt,name=account,proto3" json:"account" bson:"account"`
}

func (x *UpdateDomainInfoRequest) Reset() {
	*x = UpdateDomainInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_domain_pb_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDomainInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDomainInfoRequest) ProtoMessage() {}

func (x *UpdateDomainInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_domain_pb_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDomainInfoRequest.ProtoReflect.Descriptor instead.
func (*UpdateDomainInfoRequest) Descriptor() ([]byte, []int) {
	return file_apps_domain_pb_request_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateDomainInfoRequest) GetUpdateMode() types.UpdateMode {
	if x != nil {
		return x.UpdateMode
	}
	return types.UpdateMode(0)
}

func (x *UpdateDomainInfoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateDomainInfoRequest) GetProfile() *DomainProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *UpdateDomainInfoRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

// DescribeDomainRequest 查询domain详情请求
type DescribeDomainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"name"`
}

func (x *DescribeDomainRequest) Reset() {
	*x = DescribeDomainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_domain_pb_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeDomainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeDomainRequest) ProtoMessage() {}

func (x *DescribeDomainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_domain_pb_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeDomainRequest.ProtoReflect.Descriptor instead.
func (*DescribeDomainRequest) Descriptor() ([]byte, []int) {
	return file_apps_domain_pb_request_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeDomainRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// DeleteDomainRequest 删除域
type DeleteDomainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"name"`
}

func (x *DeleteDomainRequest) Reset() {
	*x = DeleteDomainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_domain_pb_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDomainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDomainRequest) ProtoMessage() {}

func (x *DeleteDomainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_domain_pb_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDomainRequest.ProtoReflect.Descriptor instead.
func (*DeleteDomainRequest) Descriptor() ([]byte, []int) {
	return file_apps_domain_pb_request_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteDomainRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// QueryDomainRequest 请求
type QueryDomainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// @gotags: json:"name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	// @gotags: json:"owner"
	Owner string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner"`
}

func (x *QueryDomainRequest) Reset() {
	*x = QueryDomainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_domain_pb_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDomainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDomainRequest) ProtoMessage() {}

func (x *QueryDomainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_domain_pb_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryDomainRequest.ProtoReflect.Descriptor instead.
func (*QueryDomainRequest) Descriptor() ([]byte, []int) {
	return file_apps_domain_pb_request_proto_rawDescGZIP(), []int{5}
}

func (x *QueryDomainRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryDomainRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *QueryDomainRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

// UpdateDomainSecurityRequest todo
type UpdateDomainSecurityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 更新模式
	// @gotags: bson:"update_mode" json:"update_mode"
	UpdateMode types.UpdateMode `protobuf:"varint,1,opt,name=update_mode,json=updateMode,proto3,enum=infraboard.keyauth.types.UpdateMode" json:"update_mode" bson:"update_mode"`
	// 名称
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" bson:"name"`
	// SecuritySetting 域安全设置
	// @gotags: bson:"security_setting" json:"security_setting"
	SecuritySetting *SecuritySetting `protobuf:"bytes,3,opt,name=security_setting,json=securitySetting,proto3" json:"security_setting" bson:"security_setting"`
}

func (x *UpdateDomainSecurityRequest) Reset() {
	*x = UpdateDomainSecurityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_domain_pb_request_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDomainSecurityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDomainSecurityRequest) ProtoMessage() {}

func (x *UpdateDomainSecurityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_domain_pb_request_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDomainSecurityRequest.ProtoReflect.Descriptor instead.
func (*UpdateDomainSecurityRequest) Descriptor() ([]byte, []int) {
	return file_apps_domain_pb_request_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateDomainSecurityRequest) GetUpdateMode() types.UpdateMode {
	if x != nil {
		return x.UpdateMode
	}
	return types.UpdateMode(0)
}

func (x *UpdateDomainSecurityRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateDomainSecurityRequest) GetSecuritySetting() *SecuritySetting {
	if x != nil {
		return x.SecuritySetting
	}
	return nil
}

// SecuritySetting 安全策略
type SecuritySetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 密码安全
	// @gotags: bson:"password_security" json:"password_security"
	PasswordSecurity *PasswordSecurity `protobuf:"bytes,1,opt,name=password_security,json=passwordSecurity,proto3" json:"password_security" bson:"password_security"`
	// 登录安全
	// @gotags: bson:"login_security" json:"login_security"
	LoginSecurity *LoginSecurity `protobuf:"bytes,2,opt,name=login_security,json=loginSecurity,proto3" json:"login_security" bson:"login_security"`
}

func (x *SecuritySetting) Reset() {
	*x = SecuritySetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_domain_pb_request_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecuritySetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecuritySetting) ProtoMessage() {}

func (x *SecuritySetting) ProtoReflect() protoreflect.Message {
	mi := &file_apps_domain_pb_request_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecuritySetting.ProtoReflect.Descriptor instead.
func (*SecuritySetting) Descriptor() ([]byte, []int) {
	return file_apps_domain_pb_request_proto_rawDescGZIP(), []int{7}
}

func (x *SecuritySetting) GetPasswordSecurity() *PasswordSecurity {
	if x != nil {
		return x.PasswordSecurity
	}
	return nil
}

func (x *SecuritySetting) GetLoginSecurity() *LoginSecurity {
	if x != nil {
		return x.LoginSecurity
	}
	return nil
}

// 联系人
type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 姓名
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"name"`
	// 职位
	// @gotags: bson:"title" json:"title"
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title" bson:"title"`
	// 电话
	// @gotags: bson:"phone" json:"phone"
	Phone string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone" bson:"phone"`
	// 邮箱
	// @gotags: bson:"email" json:"email"
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email" bson:"email"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_domain_pb_request_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_apps_domain_pb_request_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_apps_domain_pb_request_proto_rawDescGZIP(), []int{8}
}

func (x *Contact) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Contact) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Contact) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Contact) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_apps_domain_pb_request_proto protoreflect.FileDescriptor

var file_apps_domain_pb_request_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x62,
	0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x1c, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70,
	0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x83, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x42, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0xbd, 0x02, 0x0a, 0x0d, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x50, 0x61, 0x74, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x61, 0x78, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x61, 0x78, 0x12, 0x3c, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x6b, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x6b, 0x22, 0xd2, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x42,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x15,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x76, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0xcf, 0x01, 0x0a,
	0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b,
	0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x55, 0x0a, 0x10, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b,
	0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x22, 0xbc,
	0x01, 0x0a, 0x0f, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x58, 0x0a, 0x11, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x10, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x4f, 0x0a, 0x0e,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x0d,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x22, 0x5f, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x2b,
	0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_apps_domain_pb_request_proto_rawDescOnce sync.Once
	file_apps_domain_pb_request_proto_rawDescData = file_apps_domain_pb_request_proto_rawDesc
)

func file_apps_domain_pb_request_proto_rawDescGZIP() []byte {
	file_apps_domain_pb_request_proto_rawDescOnce.Do(func() {
		file_apps_domain_pb_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_domain_pb_request_proto_rawDescData)
	})
	return file_apps_domain_pb_request_proto_rawDescData
}

var file_apps_domain_pb_request_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_apps_domain_pb_request_proto_goTypes = []interface{}{
	(*CreateDomainRequest)(nil),         // 0: infraboard.keyauth.domain.CreateDomainRequest
	(*DomainProfile)(nil),               // 1: infraboard.keyauth.domain.DomainProfile
	(*UpdateDomainInfoRequest)(nil),     // 2: infraboard.keyauth.domain.UpdateDomainInfoRequest
	(*DescribeDomainRequest)(nil),       // 3: infraboard.keyauth.domain.DescribeDomainRequest
	(*DeleteDomainRequest)(nil),         // 4: infraboard.keyauth.domain.DeleteDomainRequest
	(*QueryDomainRequest)(nil),          // 5: infraboard.keyauth.domain.QueryDomainRequest
	(*UpdateDomainSecurityRequest)(nil), // 6: infraboard.keyauth.domain.UpdateDomainSecurityRequest
	(*SecuritySetting)(nil),             // 7: infraboard.keyauth.domain.SecuritySetting
	(*Contact)(nil),                     // 8: infraboard.keyauth.domain.Contact
	(types.UpdateMode)(0),               // 9: infraboard.keyauth.types.UpdateMode
	(*request.PageRequest)(nil),         // 10: infraboard.mcube.page.PageRequest
	(*PasswordSecurity)(nil),            // 11: infraboard.keyauth.domain.PasswordSecurity
	(*LoginSecurity)(nil),               // 12: infraboard.keyauth.domain.LoginSecurity
}
var file_apps_domain_pb_request_proto_depIdxs = []int32{
	1,  // 0: infraboard.keyauth.domain.CreateDomainRequest.profile:type_name -> infraboard.keyauth.domain.DomainProfile
	8,  // 1: infraboard.keyauth.domain.DomainProfile.contack:type_name -> infraboard.keyauth.domain.Contact
	9,  // 2: infraboard.keyauth.domain.UpdateDomainInfoRequest.update_mode:type_name -> infraboard.keyauth.types.UpdateMode
	1,  // 3: infraboard.keyauth.domain.UpdateDomainInfoRequest.profile:type_name -> infraboard.keyauth.domain.DomainProfile
	10, // 4: infraboard.keyauth.domain.QueryDomainRequest.page:type_name -> infraboard.mcube.page.PageRequest
	9,  // 5: infraboard.keyauth.domain.UpdateDomainSecurityRequest.update_mode:type_name -> infraboard.keyauth.types.UpdateMode
	7,  // 6: infraboard.keyauth.domain.UpdateDomainSecurityRequest.security_setting:type_name -> infraboard.keyauth.domain.SecuritySetting
	11, // 7: infraboard.keyauth.domain.SecuritySetting.password_security:type_name -> infraboard.keyauth.domain.PasswordSecurity
	12, // 8: infraboard.keyauth.domain.SecuritySetting.login_security:type_name -> infraboard.keyauth.domain.LoginSecurity
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_apps_domain_pb_request_proto_init() }
func file_apps_domain_pb_request_proto_init() {
	if File_apps_domain_pb_request_proto != nil {
		return
	}
	file_apps_domain_pb_setting_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_domain_pb_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDomainRequest); i {
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
		file_apps_domain_pb_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainProfile); i {
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
		file_apps_domain_pb_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDomainInfoRequest); i {
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
		file_apps_domain_pb_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeDomainRequest); i {
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
		file_apps_domain_pb_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDomainRequest); i {
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
		file_apps_domain_pb_request_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDomainRequest); i {
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
		file_apps_domain_pb_request_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDomainSecurityRequest); i {
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
		file_apps_domain_pb_request_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecuritySetting); i {
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
		file_apps_domain_pb_request_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
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
			RawDescriptor: file_apps_domain_pb_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_domain_pb_request_proto_goTypes,
		DependencyIndexes: file_apps_domain_pb_request_proto_depIdxs,
		MessageInfos:      file_apps_domain_pb_request_proto_msgTypes,
	}.Build()
	File_apps_domain_pb_request_proto = out.File
	file_apps_domain_pb_request_proto_rawDesc = nil
	file_apps_domain_pb_request_proto_goTypes = nil
	file_apps_domain_pb_request_proto_depIdxs = nil
}