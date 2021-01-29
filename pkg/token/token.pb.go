//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/token/pb/token.proto

package token

import (
	proto "github.com/golang/protobuf/proto"
	types "github.com/infraboard/keyauth/pkg/user/types"
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

type GrantType int32

const (
	GrantType_NULL      GrantType = 0
	GrantType_UNKNOWN   GrantType = 1
	GrantType_PASSWORD  GrantType = 2
	GrantType_LDAP      GrantType = 3
	GrantType_REFRESH   GrantType = 4
	GrantType_ACCESS    GrantType = 5
	GrantType_CLIENT    GrantType = 6
	GrantType_AUTH_CODE GrantType = 7
	GrantType_IMPLICIT  GrantType = 8
)

// Enum value maps for GrantType.
var (
	GrantType_name = map[int32]string{
		0: "NULL",
		1: "UNKNOWN",
		2: "PASSWORD",
		3: "LDAP",
		4: "REFRESH",
		5: "ACCESS",
		6: "CLIENT",
		7: "AUTH_CODE",
		8: "IMPLICIT",
	}
	GrantType_value = map[string]int32{
		"NULL":      0,
		"UNKNOWN":   1,
		"PASSWORD":  2,
		"LDAP":      3,
		"REFRESH":   4,
		"ACCESS":    5,
		"CLIENT":    6,
		"AUTH_CODE": 7,
		"IMPLICIT":  8,
	}
)

func (x GrantType) Enum() *GrantType {
	p := new(GrantType)
	*p = x
	return p
}

func (x GrantType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GrantType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_token_pb_token_proto_enumTypes[0].Descriptor()
}

func (GrantType) Type() protoreflect.EnumType {
	return &file_pkg_token_pb_token_proto_enumTypes[0]
}

func (x GrantType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GrantType.Descriptor instead.
func (GrantType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_token_pb_token_proto_rawDescGZIP(), []int{0}
}

type TokenType int32

const (
	TokenType_BEARER TokenType = 0
	TokenType_MAC    TokenType = 1
	TokenType_JWT    TokenType = 2
)

// Enum value maps for TokenType.
var (
	TokenType_name = map[int32]string{
		0: "BEARER",
		1: "MAC",
		2: "JWT",
	}
	TokenType_value = map[string]int32{
		"BEARER": 0,
		"MAC":    1,
		"JWT":    2,
	}
)

func (x TokenType) Enum() *TokenType {
	p := new(TokenType)
	*p = x
	return p
}

func (x TokenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_token_pb_token_proto_enumTypes[1].Descriptor()
}

func (TokenType) Type() protoreflect.EnumType {
	return &file_pkg_token_pb_token_proto_enumTypes[1]
}

func (x TokenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenType.Descriptor instead.
func (TokenType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_token_pb_token_proto_rawDescGZIP(), []int{1}
}

type BlockType int32

const (
	BlockType_SESSION_TERMINATED     BlockType = 0
	BlockType_OTHER_CLIENT_LOGGED_IN BlockType = 1
	BlockType_OTHER_PLACE_LOGGED_IN  BlockType = 2
	BlockType_OTHER_IP_LOGGED_IN     BlockType = 3
)

// Enum value maps for BlockType.
var (
	BlockType_name = map[int32]string{
		0: "SESSION_TERMINATED",
		1: "OTHER_CLIENT_LOGGED_IN",
		2: "OTHER_PLACE_LOGGED_IN",
		3: "OTHER_IP_LOGGED_IN",
	}
	BlockType_value = map[string]int32{
		"SESSION_TERMINATED":     0,
		"OTHER_CLIENT_LOGGED_IN": 1,
		"OTHER_PLACE_LOGGED_IN":  2,
		"OTHER_IP_LOGGED_IN":     3,
	}
)

func (x BlockType) Enum() *BlockType {
	p := new(BlockType)
	*p = x
	return p
}

func (x BlockType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_token_pb_token_proto_enumTypes[2].Descriptor()
}

func (BlockType) Type() protoreflect.EnumType {
	return &file_pkg_token_pb_token_proto_enumTypes[2]
}

func (x BlockType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlockType.Descriptor instead.
func (BlockType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_token_pb_token_proto_rawDescGZIP(), []int{2}
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId        string         `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id" bson:"session_id"`
	AccessToken      string         `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token" bson:"_id"`
	RefreshToken     string         `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty" bson:"refresh_token"`
	CreateAt         int64          `protobuf:"varint,5,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty" bson:"create_at"`
	AccessExpiredAt  int64          `protobuf:"varint,6,opt,name=access_expired_at,json=accessExpiredAt,proto3" json:"access_expired_at,omitempty" bson:"access_expired_at"`
	RefreshExpiredAt int64          `protobuf:"varint,7,opt,name=refresh_expired_at,json=refreshExpiredAt,proto3" json:"refresh_expired_at,omitempty" bson:"refresh_expired_at"`
	Domain           string         `protobuf:"bytes,8,opt,name=domain,proto3" json:"domain,omitempty" bson:"domain"`
	UserType         types.UserType `protobuf:"varint,9,opt,name=user_type,json=userType,proto3,enum=keyauth.user.UserType" json:"user_type,omitempty" bson:"user_type"`
	Account          string         `protobuf:"bytes,10,opt,name=account,proto3" json:"account,omitempty" bson:"account"`
	ApplicationId    string         `protobuf:"bytes,11,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty" bson:"application_id"`
	ApplicationName  string         `protobuf:"bytes,12,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty" bson:"application_name"`
	ClientId         string         `protobuf:"bytes,13,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty" bson:"client_id"`
	StartGrantType   GrantType      `protobuf:"varint,14,opt,name=start_grant_type,json=startGrantType,proto3,enum=keyauth.token.GrantType" json:"start_grant_type,omitempty" bson:"start_grant_type"`
	GrantType        GrantType      `protobuf:"varint,15,opt,name=grant_type,json=grantType,proto3,enum=keyauth.token.GrantType" json:"grant_type,omitempty" bson:"grant_type"`
	Type             TokenType      `protobuf:"varint,16,opt,name=type,proto3,enum=keyauth.token.TokenType" json:"type,omitempty" bson:"type"`
	Scope            string         `protobuf:"bytes,17,opt,name=scope,proto3" json:"scope,omitempty" bson:"scope"`
	Description      string         `protobuf:"bytes,18,opt,name=description,proto3" json:"description,omitempty" bson:"description"`
	IsBlock          bool           `protobuf:"varint,19,opt,name=is_block,json=isBlock,proto3" json:"is_block" bson:"is_block"`
	BlockType        BlockType      `protobuf:"varint,20,opt,name=block_type,json=blockType,proto3,enum=keyauth.token.BlockType" json:"block_type" bson:"block_type"`
	BlockAt          int64          `protobuf:"varint,21,opt,name=block_at,json=blockAt,proto3" json:"block_at" bson:"block_at"`
	BlockReason      string         `protobuf:"bytes,22,opt,name=block_reason,json=blockReason,proto3" json:"block_reason,omitempty" bson:"block_reason"`
	RemoteIp         string         `protobuf:"bytes,23,opt,name=remote_ip,json=remoteIp,proto3" json:"-" bson:"-"`
	UserAgent        string         `protobuf:"bytes,24,opt,name=user_agent,json=userAgent,proto3" json:"-" bson:"-"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_token_pb_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_token_pb_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_pkg_token_pb_token_proto_rawDescGZIP(), []int{0}
}

func (x *Token) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *Token) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Token) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *Token) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Token) GetAccessExpiredAt() int64 {
	if x != nil {
		return x.AccessExpiredAt
	}
	return 0
}

func (x *Token) GetRefreshExpiredAt() int64 {
	if x != nil {
		return x.RefreshExpiredAt
	}
	return 0
}

func (x *Token) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Token) GetUserType() types.UserType {
	if x != nil {
		return x.UserType
	}
	return types.UserType_SUB
}

func (x *Token) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *Token) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *Token) GetApplicationName() string {
	if x != nil {
		return x.ApplicationName
	}
	return ""
}

func (x *Token) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Token) GetStartGrantType() GrantType {
	if x != nil {
		return x.StartGrantType
	}
	return GrantType_NULL
}

func (x *Token) GetGrantType() GrantType {
	if x != nil {
		return x.GrantType
	}
	return GrantType_NULL
}

func (x *Token) GetType() TokenType {
	if x != nil {
		return x.Type
	}
	return TokenType_BEARER
}

func (x *Token) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *Token) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Token) GetIsBlock() bool {
	if x != nil {
		return x.IsBlock
	}
	return false
}

func (x *Token) GetBlockType() BlockType {
	if x != nil {
		return x.BlockType
	}
	return BlockType_SESSION_TERMINATED
}

func (x *Token) GetBlockAt() int64 {
	if x != nil {
		return x.BlockAt
	}
	return 0
}

func (x *Token) GetBlockReason() string {
	if x != nil {
		return x.BlockReason
	}
	return ""
}

func (x *Token) GetRemoteIp() string {
	if x != nil {
		return x.RemoteIp
	}
	return ""
}

func (x *Token) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

type Set struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64    `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	Items []*Token `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *Set) Reset() {
	*x = Set{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_token_pb_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_token_pb_token_proto_msgTypes[1]
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
	return file_pkg_token_pb_token_proto_rawDescGZIP(), []int{1}
}

func (x *Set) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Set) GetItems() []*Token {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_pkg_token_pb_token_proto protoreflect.FileDescriptor

var file_pkg_token_pb_token_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6b, 0x65, 0x79, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x65,
	0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x67,
	0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x6b, 0x67, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x0f, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x48, 0x0a,
	0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x29, 0xc2, 0xde, 0x1f, 0x25, 0x0a, 0x23, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x22, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x24, 0xc2,
	0xde, 0x1f, 0x20, 0x0a, 0x1e, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64, 0x22, 0x20,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x5e, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x39, 0xc2, 0xde, 0x1f, 0x35, 0x0a, 0x33, 0x62,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x4e, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x31, 0xc2, 0xde, 0x1f, 0x2d, 0x0a, 0x2b, 0x62, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x2c, 0x6f, 0x6d, 0x69, 0x74,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x12, 0x6d, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x41, 0xc2, 0xde, 0x1f,
	0x3d, 0x0a, 0x3b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x0f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x71, 0x0a, 0x12, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x42, 0x43, 0xc2, 0xde, 0x1f,
	0x3f, 0x0a, 0x3d, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x52, 0x10, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x43, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x2b, 0xc2, 0xde, 0x1f, 0x27, 0x0a, 0x25, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x66, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6b, 0x65, 0x79,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x31, 0xc2, 0xde, 0x1f, 0x2d, 0x0a, 0x2b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x47, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2d, 0xc2, 0xde, 0x1f, 0x29, 0x0a, 0x27, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x62, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x3b, 0xc2, 0xde, 0x1f, 0x37, 0x0a, 0x35, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x0d, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x6a, 0x0a, 0x10,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3f, 0xc2, 0xde, 0x1f, 0x3b, 0x0a, 0x39, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x6f, 0x6d, 0x69,
	0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x42, 0x31, 0xc2, 0xde, 0x1f,
	0x2d, 0x0a, 0x2b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x83, 0x01, 0x0a, 0x10, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x3f, 0xc2,
	0xde, 0x1f, 0x3b, 0x0a, 0x39, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x0e,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x6c,
	0x0a, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x33, 0xc2, 0xde,
	0x1f, 0x2f, 0x0a, 0x2d, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x52, 0x09, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x55, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6b, 0x65, 0x79,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79,
	0x70, 0x65, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x3f, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x29, 0xc2, 0xde, 0x1f, 0x25, 0x0a, 0x23, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x05, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x12, 0x57, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x42, 0x35, 0xc2, 0xde, 0x1f, 0x31, 0x0a,
	0x2f, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a,
	0x08, 0x69, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x13, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x25, 0xc2, 0xde, 0x1f, 0x21, 0x0a, 0x1f, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x73, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x73, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x52, 0x07, 0x69, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x62, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x42, 0x29, 0xc2,
	0xde, 0x1f, 0x25, 0x0a, 0x23, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x61, 0x74, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x03, 0x42, 0x25, 0xc2, 0xde, 0x1f, 0x21, 0x0a, 0x1f, 0x62, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x61, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x61, 0x74, 0x22, 0x52, 0x07, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x41, 0x74, 0x12, 0x5a, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x42, 0x37, 0xc2, 0xde, 0x1f,
	0x33, 0x0a, 0x31, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x34, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x70, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xc2, 0xde, 0x1f, 0x13, 0x0a, 0x11, 0x62, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x2d, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2d, 0x22, 0x52, 0x08, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x70, 0x12, 0x36, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xc2, 0xde, 0x1f,
	0x13, 0x0a, 0x11, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2d, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x2d, 0x22, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x22,
	0x89, 0x01, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x1f, 0xc2, 0xde, 0x1f, 0x1b, 0x0a, 0x19, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x4b,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x42, 0x1f, 0xc2, 0xde, 0x1f, 0x1b, 0x0a, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2a, 0x7c, 0x0a, 0x09, 0x47,
	0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x55, 0x4c, 0x4c,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x02, 0x12, 0x08, 0x0a,
	0x04, 0x4c, 0x44, 0x41, 0x50, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x46, 0x52, 0x45,
	0x53, 0x48, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x05,
	0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09,
	0x41, 0x55, 0x54, 0x48, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x49,
	0x4d, 0x50, 0x4c, 0x49, 0x43, 0x49, 0x54, 0x10, 0x08, 0x2a, 0x29, 0x0a, 0x09, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x45, 0x41, 0x52, 0x45, 0x52,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x43, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4a,
	0x57, 0x54, 0x10, 0x02, 0x2a, 0x72, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x45, 0x52,
	0x4d, 0x49, 0x4e, 0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x4f, 0x54, 0x48,
	0x45, 0x52, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x4c, 0x4f, 0x47, 0x47, 0x45, 0x44,
	0x5f, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f, 0x50,
	0x4c, 0x41, 0x43, 0x45, 0x5f, 0x4c, 0x4f, 0x47, 0x47, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x10, 0x02,
	0x12, 0x16, 0x0a, 0x12, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f, 0x49, 0x50, 0x5f, 0x4c, 0x4f, 0x47,
	0x47, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x10, 0x03, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_token_pb_token_proto_rawDescOnce sync.Once
	file_pkg_token_pb_token_proto_rawDescData = file_pkg_token_pb_token_proto_rawDesc
)

func file_pkg_token_pb_token_proto_rawDescGZIP() []byte {
	file_pkg_token_pb_token_proto_rawDescOnce.Do(func() {
		file_pkg_token_pb_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_token_pb_token_proto_rawDescData)
	})
	return file_pkg_token_pb_token_proto_rawDescData
}

var file_pkg_token_pb_token_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_pkg_token_pb_token_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_token_pb_token_proto_goTypes = []interface{}{
	(GrantType)(0),      // 0: keyauth.token.GrantType
	(TokenType)(0),      // 1: keyauth.token.TokenType
	(BlockType)(0),      // 2: keyauth.token.BlockType
	(*Token)(nil),       // 3: keyauth.token.Token
	(*Set)(nil),         // 4: keyauth.token.Set
	(types.UserType)(0), // 5: keyauth.user.UserType
}
var file_pkg_token_pb_token_proto_depIdxs = []int32{
	5, // 0: keyauth.token.Token.user_type:type_name -> keyauth.user.UserType
	0, // 1: keyauth.token.Token.start_grant_type:type_name -> keyauth.token.GrantType
	0, // 2: keyauth.token.Token.grant_type:type_name -> keyauth.token.GrantType
	1, // 3: keyauth.token.Token.type:type_name -> keyauth.token.TokenType
	2, // 4: keyauth.token.Token.block_type:type_name -> keyauth.token.BlockType
	3, // 5: keyauth.token.Set.items:type_name -> keyauth.token.Token
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_token_pb_token_proto_init() }
func file_pkg_token_pb_token_proto_init() {
	if File_pkg_token_pb_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_token_pb_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_pkg_token_pb_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_token_pb_token_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_token_pb_token_proto_goTypes,
		DependencyIndexes: file_pkg_token_pb_token_proto_depIdxs,
		EnumInfos:         file_pkg_token_pb_token_proto_enumTypes,
		MessageInfos:      file_pkg_token_pb_token_proto_msgTypes,
	}.Build()
	File_pkg_token_pb_token_proto = out.File
	file_pkg_token_pb_token_proto_rawDesc = nil
	file_pkg_token_pb_token_proto_goTypes = nil
	file_pkg_token_pb_token_proto_depIdxs = nil
}