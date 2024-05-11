// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: authorization_server/authorization_server.proto

package authorization_server

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ResponseType int32

const (
	ResponseType_code  ResponseType = 0
	ResponseType_token ResponseType = 1
)

// Enum value maps for ResponseType.
var (
	ResponseType_name = map[int32]string{
		0: "code",
		1: "token",
	}
	ResponseType_value = map[string]int32{
		"code":  0,
		"token": 1,
	}
)

func (x ResponseType) Enum() *ResponseType {
	p := new(ResponseType)
	*p = x
	return p
}

func (x ResponseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseType) Descriptor() protoreflect.EnumDescriptor {
	return file_authorization_server_authorization_server_proto_enumTypes[0].Descriptor()
}

func (ResponseType) Type() protoreflect.EnumType {
	return &file_authorization_server_authorization_server_proto_enumTypes[0]
}

func (x ResponseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseType.Descriptor instead.
func (ResponseType) EnumDescriptor() ([]byte, []int) {
	return file_authorization_server_authorization_server_proto_rawDescGZIP(), []int{0}
}

type GrantType int32

const (
	GrantType_authorization_code GrantType = 0
	GrantType_refresh_token      GrantType = 1
)

// Enum value maps for GrantType.
var (
	GrantType_name = map[int32]string{
		0: "authorization_code",
		1: "refresh_token",
	}
	GrantType_value = map[string]int32{
		"authorization_code": 0,
		"refresh_token":      1,
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
	return file_authorization_server_authorization_server_proto_enumTypes[1].Descriptor()
}

func (GrantType) Type() protoreflect.EnumType {
	return &file_authorization_server_authorization_server_proto_enumTypes[1]
}

func (x GrantType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GrantType.Descriptor instead.
func (GrantType) EnumDescriptor() ([]byte, []int) {
	return file_authorization_server_authorization_server_proto_rawDescGZIP(), []int{1}
}

type RegisterClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url         string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	RedirectUri string `protobuf:"bytes,3,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
}

func (x *RegisterClientRequest) Reset() {
	*x = RegisterClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorization_server_authorization_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterClientRequest) ProtoMessage() {}

func (x *RegisterClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_server_authorization_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterClientRequest.ProtoReflect.Descriptor instead.
func (*RegisterClientRequest) Descriptor() ([]byte, []int) {
	return file_authorization_server_authorization_server_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterClientRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterClientRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *RegisterClientRequest) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

type RegisterClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
}

func (x *RegisterClientResponse) Reset() {
	*x = RegisterClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorization_server_authorization_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterClientResponse) ProtoMessage() {}

func (x *RegisterClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_server_authorization_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterClientResponse.ProtoReflect.Descriptor instead.
func (*RegisterClientResponse) Descriptor() ([]byte, []int) {
	return file_authorization_server_authorization_server_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterClientResponse) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *RegisterClientResponse) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

type GetAuthorizationCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseType ResponseType `protobuf:"varint,1,opt,name=response_type,json=responseType,proto3,enum=authorization_server.ResponseType" json:"response_type,omitempty"`
	ClientId     string       `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	RedirectUri  *string      `protobuf:"bytes,3,opt,name=redirect_uri,json=redirectUri,proto3,oneof" json:"redirect_uri,omitempty"`
	Scope        *string      `protobuf:"bytes,4,opt,name=scope,proto3,oneof" json:"scope,omitempty"`
	// state recommended to prevent CSRF attacks
	State *string `protobuf:"bytes,5,opt,name=state,proto3,oneof" json:"state,omitempty"`
}

func (x *GetAuthorizationCodeRequest) Reset() {
	*x = GetAuthorizationCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorization_server_authorization_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorizationCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorizationCodeRequest) ProtoMessage() {}

func (x *GetAuthorizationCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_server_authorization_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorizationCodeRequest.ProtoReflect.Descriptor instead.
func (*GetAuthorizationCodeRequest) Descriptor() ([]byte, []int) {
	return file_authorization_server_authorization_server_proto_rawDescGZIP(), []int{2}
}

func (x *GetAuthorizationCodeRequest) GetResponseType() ResponseType {
	if x != nil {
		return x.ResponseType
	}
	return ResponseType_code
}

func (x *GetAuthorizationCodeRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *GetAuthorizationCodeRequest) GetRedirectUri() string {
	if x != nil && x.RedirectUri != nil {
		return *x.RedirectUri
	}
	return ""
}

func (x *GetAuthorizationCodeRequest) GetScope() string {
	if x != nil && x.Scope != nil {
		return *x.Scope
	}
	return ""
}

func (x *GetAuthorizationCodeRequest) GetState() string {
	if x != nil && x.State != nil {
		return *x.State
	}
	return ""
}

type GetAuthorizationCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// state recommended to prevent CSRF attacks
	State *string `protobuf:"bytes,2,opt,name=state,proto3,oneof" json:"state,omitempty"`
}

func (x *GetAuthorizationCodeResponse) Reset() {
	*x = GetAuthorizationCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorization_server_authorization_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorizationCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorizationCodeResponse) ProtoMessage() {}

func (x *GetAuthorizationCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_server_authorization_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorizationCodeResponse.ProtoReflect.Descriptor instead.
func (*GetAuthorizationCodeResponse) Descriptor() ([]byte, []int) {
	return file_authorization_server_authorization_server_proto_rawDescGZIP(), []int{3}
}

func (x *GetAuthorizationCodeResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetAuthorizationCodeResponse) GetState() string {
	if x != nil && x.State != nil {
		return *x.State
	}
	return ""
}

type GetAccessTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GrantType GrantType `protobuf:"varint,1,opt,name=grant_type,json=grantType,proto3,enum=authorization_server.GrantType" json:"grant_type,omitempty"`
	Code      string    `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	// required if was included in GetAuthorizationCodeRequest
	// value must be identical to value included in GetAuthorizationCodeRequest
	RedirectUri *string `protobuf:"bytes,3,opt,name=redirect_uri,json=redirectUri,proto3,oneof" json:"redirect_uri,omitempty"`
	ClientId    string  `protobuf:"bytes,4,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// client_secret send if authorizing server-side client
	ClientSecret *string `protobuf:"bytes,5,opt,name=client_secret,json=clientSecret,proto3,oneof" json:"client_secret,omitempty"`
	// refresh_token is required if refresh token grant type specified
	// allows to get new access token
	RefreshToken *string `protobuf:"bytes,6,opt,name=refresh_token,json=refreshToken,proto3,oneof" json:"refresh_token,omitempty"`
}

func (x *GetAccessTokenRequest) Reset() {
	*x = GetAccessTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorization_server_authorization_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessTokenRequest) ProtoMessage() {}

func (x *GetAccessTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_server_authorization_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessTokenRequest.ProtoReflect.Descriptor instead.
func (*GetAccessTokenRequest) Descriptor() ([]byte, []int) {
	return file_authorization_server_authorization_server_proto_rawDescGZIP(), []int{4}
}

func (x *GetAccessTokenRequest) GetGrantType() GrantType {
	if x != nil {
		return x.GrantType
	}
	return GrantType_authorization_code
}

func (x *GetAccessTokenRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetAccessTokenRequest) GetRedirectUri() string {
	if x != nil && x.RedirectUri != nil {
		return *x.RedirectUri
	}
	return ""
}

func (x *GetAccessTokenRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *GetAccessTokenRequest) GetClientSecret() string {
	if x != nil && x.ClientSecret != nil {
		return *x.ClientSecret
	}
	return ""
}

func (x *GetAccessTokenRequest) GetRefreshToken() string {
	if x != nil && x.RefreshToken != nil {
		return *x.RefreshToken
	}
	return ""
}

type GetAccessTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// bearer/mac/etc.
	TokenType string `protobuf:"bytes,2,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	// expires_in could be omitted,
	// but should be provided to client by other means or be documented as default value
	ExpiresIn    int64   `protobuf:"varint,3,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	RefreshToken *string `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken,proto3,oneof" json:"refresh_token,omitempty"`
	// scope is optional if identical to the scope requested by the client, otherwise, required
	Scope *string `protobuf:"bytes,5,opt,name=scope,proto3,oneof" json:"scope,omitempty"`
}

func (x *GetAccessTokenResponse) Reset() {
	*x = GetAccessTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorization_server_authorization_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessTokenResponse) ProtoMessage() {}

func (x *GetAccessTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_server_authorization_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessTokenResponse.ProtoReflect.Descriptor instead.
func (*GetAccessTokenResponse) Descriptor() ([]byte, []int) {
	return file_authorization_server_authorization_server_proto_rawDescGZIP(), []int{5}
}

func (x *GetAccessTokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *GetAccessTokenResponse) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *GetAccessTokenResponse) GetExpiresIn() int64 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

func (x *GetAccessTokenResponse) GetRefreshToken() string {
	if x != nil && x.RefreshToken != nil {
		return *x.RefreshToken
	}
	return ""
}

func (x *GetAccessTokenResponse) GetScope() string {
	if x != nil && x.Scope != nil {
		return *x.Scope
	}
	return ""
}

var File_authorization_server_authorization_server_proto protoreflect.FileDescriptor

var file_authorization_server_authorization_server_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x5f, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x69, 0x22, 0x5a, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x22, 0x86, 0x02, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0c, 0x72, 0x65, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x69, 0x88, 0x01,
	0x01, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x72, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x57, 0x0a, 0x1c,
	0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x19, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0xb9, 0x02, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3e, 0x0a, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f,
	0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x72, 0x65, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x69, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0c, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d,
	0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x42, 0x10, 0x0a,
	0x0e, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0xda, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x12, 0x28, 0x0a,
	0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2a, 0x23,
	0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x10, 0x01, 0x2a, 0x36, 0x0a, 0x09, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x10, 0x01, 0x32, 0xcf, 0x03, 0x0a, 0x13,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x8f, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x9a, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x31,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x32, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x12, 0x88, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x13, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x5f, 0x5a,
	0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x64, 0x72,
	0x65, 0x79, 0x42, 0x4d, 0x57, 0x58, 0x36, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authorization_server_authorization_server_proto_rawDescOnce sync.Once
	file_authorization_server_authorization_server_proto_rawDescData = file_authorization_server_authorization_server_proto_rawDesc
)

func file_authorization_server_authorization_server_proto_rawDescGZIP() []byte {
	file_authorization_server_authorization_server_proto_rawDescOnce.Do(func() {
		file_authorization_server_authorization_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_authorization_server_authorization_server_proto_rawDescData)
	})
	return file_authorization_server_authorization_server_proto_rawDescData
}

var file_authorization_server_authorization_server_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_authorization_server_authorization_server_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_authorization_server_authorization_server_proto_goTypes = []interface{}{
	(ResponseType)(0),                    // 0: authorization_server.ResponseType
	(GrantType)(0),                       // 1: authorization_server.GrantType
	(*RegisterClientRequest)(nil),        // 2: authorization_server.RegisterClientRequest
	(*RegisterClientResponse)(nil),       // 3: authorization_server.RegisterClientResponse
	(*GetAuthorizationCodeRequest)(nil),  // 4: authorization_server.GetAuthorizationCodeRequest
	(*GetAuthorizationCodeResponse)(nil), // 5: authorization_server.GetAuthorizationCodeResponse
	(*GetAccessTokenRequest)(nil),        // 6: authorization_server.GetAccessTokenRequest
	(*GetAccessTokenResponse)(nil),       // 7: authorization_server.GetAccessTokenResponse
}
var file_authorization_server_authorization_server_proto_depIdxs = []int32{
	0, // 0: authorization_server.GetAuthorizationCodeRequest.response_type:type_name -> authorization_server.ResponseType
	1, // 1: authorization_server.GetAccessTokenRequest.grant_type:type_name -> authorization_server.GrantType
	2, // 2: authorization_server.AuthorizationServer.RegisterClient:input_type -> authorization_server.RegisterClientRequest
	4, // 3: authorization_server.AuthorizationServer.GetAuthorizationCode:input_type -> authorization_server.GetAuthorizationCodeRequest
	6, // 4: authorization_server.AuthorizationServer.GetAccessToken:input_type -> authorization_server.GetAccessTokenRequest
	3, // 5: authorization_server.AuthorizationServer.RegisterClient:output_type -> authorization_server.RegisterClientResponse
	5, // 6: authorization_server.AuthorizationServer.GetAuthorizationCode:output_type -> authorization_server.GetAuthorizationCodeResponse
	7, // 7: authorization_server.AuthorizationServer.GetAccessToken:output_type -> authorization_server.GetAccessTokenResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_authorization_server_authorization_server_proto_init() }
func file_authorization_server_authorization_server_proto_init() {
	if File_authorization_server_authorization_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authorization_server_authorization_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterClientRequest); i {
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
		file_authorization_server_authorization_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterClientResponse); i {
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
		file_authorization_server_authorization_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorizationCodeRequest); i {
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
		file_authorization_server_authorization_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorizationCodeResponse); i {
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
		file_authorization_server_authorization_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessTokenRequest); i {
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
		file_authorization_server_authorization_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessTokenResponse); i {
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
	file_authorization_server_authorization_server_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_authorization_server_authorization_server_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_authorization_server_authorization_server_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_authorization_server_authorization_server_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_authorization_server_authorization_server_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authorization_server_authorization_server_proto_goTypes,
		DependencyIndexes: file_authorization_server_authorization_server_proto_depIdxs,
		EnumInfos:         file_authorization_server_authorization_server_proto_enumTypes,
		MessageInfos:      file_authorization_server_authorization_server_proto_msgTypes,
	}.Build()
	File_authorization_server_authorization_server_proto = out.File
	file_authorization_server_authorization_server_proto_rawDesc = nil
	file_authorization_server_authorization_server_proto_goTypes = nil
	file_authorization_server_authorization_server_proto_depIdxs = nil
}
