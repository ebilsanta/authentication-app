// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: api/proto/authentication.proto

package authentication

import (
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

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Company   string `protobuf:"bytes,1,opt,name=company,proto3" json:"company,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Birthdate string `protobuf:"bytes,5,opt,name=birthdate,proto3" json:"birthdate,omitempty"`
	Password  string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_authentication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_authentication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_authentication_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *RegisterRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *RegisterRequest) GetBirthdate() string {
	if x != nil {
		return x.Birthdate
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email           string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Message         string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	VerificationKey string `protobuf:"bytes,3,opt,name=verificationKey,proto3" json:"verificationKey,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_authentication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_authentication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_authentication_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RegisterResponse) GetVerificationKey() string {
	if x != nil {
		return x.VerificationKey
	}
	return ""
}

type VerifyEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VerificationKey string `protobuf:"bytes,1,opt,name=verificationKey,proto3" json:"verificationKey,omitempty"`
	Otp             string `protobuf:"bytes,2,opt,name=otp,proto3" json:"otp,omitempty"`
	Email           string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *VerifyEmailRequest) Reset() {
	*x = VerifyEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_authentication_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyEmailRequest) ProtoMessage() {}

func (x *VerifyEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_authentication_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyEmailRequest.ProtoReflect.Descriptor instead.
func (*VerifyEmailRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_authentication_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyEmailRequest) GetVerificationKey() string {
	if x != nil {
		return x.VerificationKey
	}
	return ""
}

func (x *VerifyEmailRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

func (x *VerifyEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type VerifyEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Details string `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	Email   string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *VerifyEmailResponse) Reset() {
	*x = VerifyEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_authentication_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyEmailResponse) ProtoMessage() {}

func (x *VerifyEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_authentication_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyEmailResponse.ProtoReflect.Descriptor instead.
func (*VerifyEmailResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_authentication_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyEmailResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *VerifyEmailResponse) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *VerifyEmailResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Company  string `protobuf:"bytes,1,opt,name=company,proto3" json:"company,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_authentication_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_authentication_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_authentication_proto_rawDescGZIP(), []int{4}
}

func (x *LoginRequest) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IdToken string `protobuf:"bytes,2,opt,name=idToken,proto3" json:"idToken,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_authentication_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_authentication_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_authentication_proto_rawDescGZIP(), []int{5}
}

func (x *LoginResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *LoginResponse) GetIdToken() string {
	if x != nil {
		return x.IdToken
	}
	return ""
}

type NewOTPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Company string `protobuf:"bytes,1,opt,name=company,proto3" json:"company,omitempty"`
	Email   string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *NewOTPRequest) Reset() {
	*x = NewOTPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_authentication_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewOTPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewOTPRequest) ProtoMessage() {}

func (x *NewOTPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_authentication_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewOTPRequest.ProtoReflect.Descriptor instead.
func (*NewOTPRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_authentication_proto_rawDescGZIP(), []int{6}
}

func (x *NewOTPRequest) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *NewOTPRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type NewOTPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message         string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	VerificationKey string `protobuf:"bytes,2,opt,name=verificationKey,proto3" json:"verificationKey,omitempty"`
}

func (x *NewOTPResponse) Reset() {
	*x = NewOTPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_authentication_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewOTPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewOTPResponse) ProtoMessage() {}

func (x *NewOTPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_authentication_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewOTPResponse.ProtoReflect.Descriptor instead.
func (*NewOTPResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_authentication_proto_rawDescGZIP(), []int{7}
}

func (x *NewOTPResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *NewOTPResponse) GetVerificationKey() string {
	if x != nil {
		return x.VerificationKey
	}
	return ""
}

type ValidTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VerificationKey string `protobuf:"bytes,1,opt,name=verificationKey,proto3" json:"verificationKey,omitempty"`
	Otp             string `protobuf:"bytes,2,opt,name=otp,proto3" json:"otp,omitempty"`
	Email           string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *ValidTokenRequest) Reset() {
	*x = ValidTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_authentication_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidTokenRequest) ProtoMessage() {}

func (x *ValidTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_authentication_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidTokenRequest.ProtoReflect.Descriptor instead.
func (*ValidTokenRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_authentication_proto_rawDescGZIP(), []int{8}
}

func (x *ValidTokenRequest) GetVerificationKey() string {
	if x != nil {
		return x.VerificationKey
	}
	return ""
}

func (x *ValidTokenRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

func (x *ValidTokenRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ValidTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Token   string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ValidTokenResponse) Reset() {
	*x = ValidTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_authentication_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidTokenResponse) ProtoMessage() {}

func (x *ValidTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_authentication_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidTokenResponse.ProtoReflect.Descriptor instead.
func (*ValidTokenResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_authentication_proto_rawDescGZIP(), []int{9}
}

func (x *ValidTokenResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ValidTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ChangePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Company  string `protobuf:"bytes,2,opt,name=company,proto3" json:"company,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ChangePasswordRequest) Reset() {
	*x = ChangePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_authentication_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordRequest) ProtoMessage() {}

func (x *ChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_authentication_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangePasswordRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_authentication_proto_rawDescGZIP(), []int{10}
}

func (x *ChangePasswordRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ChangePasswordRequest) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *ChangePasswordRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ChangePasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ChangePasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Details string `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	Email   string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *ChangePasswordResponse) Reset() {
	*x = ChangePasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_authentication_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordResponse) ProtoMessage() {}

func (x *ChangePasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_authentication_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordResponse.ProtoReflect.Descriptor instead.
func (*ChangePasswordResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_authentication_proto_rawDescGZIP(), []int{11}
}

func (x *ChangePasswordResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ChangePasswordResponse) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *ChangePasswordResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_api_proto_authentication_proto protoreflect.FileDescriptor

var file_api_proto_authentication_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb5, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x6c, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x22, 0x66, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x5d,
	0x0a, 0x13, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x5a, 0x0a,
	0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x41, 0x0a, 0x0d, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3f, 0x0a, 0x0d,
	0x4e, 0x65, 0x77, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x54, 0x0a,
	0x0e, 0x4e, 0x65, 0x77, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4b, 0x65, 0x79, 0x22, 0x65, 0x0a, 0x11, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b,
	0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6f, 0x74, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x44, 0x0a, 0x12, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x79, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x60, 0x0a, 0x16, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x32, 0xc8, 0x02,
	0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2f, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x38, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x13, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0d, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x4e, 0x65, 0x77, 0x4f, 0x54, 0x50, 0x12, 0x0e, 0x2e,
	0x4e, 0x65, 0x77, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x4e, 0x65, 0x77, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35,
	0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x2e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x50, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x73, 0x33, 0x30, 0x31, 0x2d, 0x69, 0x74, 0x73,
	0x61, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x32, 0x30, 0x32, 0x33, 0x2d, 0x32,
	0x34, 0x74, 0x31, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x32, 0x30, 0x32, 0x33,
	0x2d, 0x32, 0x34, 0x74, 0x31, 0x2d, 0x67, 0x32, 0x2d, 0x74, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_proto_authentication_proto_rawDescOnce sync.Once
	file_api_proto_authentication_proto_rawDescData = file_api_proto_authentication_proto_rawDesc
)

func file_api_proto_authentication_proto_rawDescGZIP() []byte {
	file_api_proto_authentication_proto_rawDescOnce.Do(func() {
		file_api_proto_authentication_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_authentication_proto_rawDescData)
	})
	return file_api_proto_authentication_proto_rawDescData
}

var file_api_proto_authentication_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_proto_authentication_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),        // 0: RegisterRequest
	(*RegisterResponse)(nil),       // 1: RegisterResponse
	(*VerifyEmailRequest)(nil),     // 2: VerifyEmailRequest
	(*VerifyEmailResponse)(nil),    // 3: VerifyEmailResponse
	(*LoginRequest)(nil),           // 4: LoginRequest
	(*LoginResponse)(nil),          // 5: LoginResponse
	(*NewOTPRequest)(nil),          // 6: NewOTPRequest
	(*NewOTPResponse)(nil),         // 7: NewOTPResponse
	(*ValidTokenRequest)(nil),      // 8: ValidTokenRequest
	(*ValidTokenResponse)(nil),     // 9: ValidTokenResponse
	(*ChangePasswordRequest)(nil),  // 10: ChangePasswordRequest
	(*ChangePasswordResponse)(nil), // 11: ChangePasswordResponse
}
var file_api_proto_authentication_proto_depIdxs = []int32{
	0,  // 0: Authentication.Register:input_type -> RegisterRequest
	2,  // 1: Authentication.VerifyEmail:input_type -> VerifyEmailRequest
	4,  // 2: Authentication.Login:input_type -> LoginRequest
	6,  // 3: Authentication.NewOTP:input_type -> NewOTPRequest
	8,  // 4: Authentication.ValidToken:input_type -> ValidTokenRequest
	10, // 5: Authentication.ChangePassword:input_type -> ChangePasswordRequest
	1,  // 6: Authentication.Register:output_type -> RegisterResponse
	3,  // 7: Authentication.VerifyEmail:output_type -> VerifyEmailResponse
	5,  // 8: Authentication.Login:output_type -> LoginResponse
	7,  // 9: Authentication.NewOTP:output_type -> NewOTPResponse
	9,  // 10: Authentication.ValidToken:output_type -> ValidTokenResponse
	11, // 11: Authentication.ChangePassword:output_type -> ChangePasswordResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_authentication_proto_init() }
func file_api_proto_authentication_proto_init() {
	if File_api_proto_authentication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_authentication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_api_proto_authentication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_api_proto_authentication_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyEmailRequest); i {
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
		file_api_proto_authentication_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyEmailResponse); i {
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
		file_api_proto_authentication_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_api_proto_authentication_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_api_proto_authentication_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewOTPRequest); i {
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
		file_api_proto_authentication_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewOTPResponse); i {
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
		file_api_proto_authentication_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidTokenRequest); i {
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
		file_api_proto_authentication_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidTokenResponse); i {
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
		file_api_proto_authentication_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordRequest); i {
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
		file_api_proto_authentication_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordResponse); i {
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
			RawDescriptor: file_api_proto_authentication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_authentication_proto_goTypes,
		DependencyIndexes: file_api_proto_authentication_proto_depIdxs,
		MessageInfos:      file_api_proto_authentication_proto_msgTypes,
	}.Build()
	File_api_proto_authentication_proto = out.File
	file_api_proto_authentication_proto_rawDesc = nil
	file_api_proto_authentication_proto_goTypes = nil
	file_api_proto_authentication_proto_depIdxs = nil
}
