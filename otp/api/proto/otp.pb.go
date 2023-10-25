// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.0
// source: api/proto/otp.proto

package otp

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

type OTPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Company string `protobuf:"bytes,1,opt,name=company,proto3" json:"company,omitempty"`
	Email   string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *OTPRequest) Reset() {
	*x = OTPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_otp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OTPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OTPRequest) ProtoMessage() {}

func (x *OTPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_otp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OTPRequest.ProtoReflect.Descriptor instead.
func (*OTPRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_otp_proto_rawDescGZIP(), []int{0}
}

func (x *OTPRequest) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *OTPRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type OTPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VerificationKey string `protobuf:"bytes,1,opt,name=verificationKey,proto3" json:"verificationKey,omitempty"`
	Message         string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *OTPResponse) Reset() {
	*x = OTPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_otp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OTPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OTPResponse) ProtoMessage() {}

func (x *OTPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_otp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OTPResponse.ProtoReflect.Descriptor instead.
func (*OTPResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_otp_proto_rawDescGZIP(), []int{1}
}

func (x *OTPResponse) GetVerificationKey() string {
	if x != nil {
		return x.VerificationKey
	}
	return ""
}

func (x *OTPResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type VerifyOTPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VerificationKey string `protobuf:"bytes,1,opt,name=verificationKey,proto3" json:"verificationKey,omitempty"`
	Otp             string `protobuf:"bytes,2,opt,name=otp,proto3" json:"otp,omitempty"`
	Email           string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *VerifyOTPRequest) Reset() {
	*x = VerifyOTPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_otp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyOTPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyOTPRequest) ProtoMessage() {}

func (x *VerifyOTPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_otp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyOTPRequest.ProtoReflect.Descriptor instead.
func (*VerifyOTPRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_otp_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyOTPRequest) GetVerificationKey() string {
	if x != nil {
		return x.VerificationKey
	}
	return ""
}

func (x *VerifyOTPRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

func (x *VerifyOTPRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type VerifyOTPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Details string `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	Company string `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
	Email   string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *VerifyOTPResponse) Reset() {
	*x = VerifyOTPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_otp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyOTPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyOTPResponse) ProtoMessage() {}

func (x *VerifyOTPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_otp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyOTPResponse.ProtoReflect.Descriptor instead.
func (*VerifyOTPResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_otp_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyOTPResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *VerifyOTPResponse) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *VerifyOTPResponse) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *VerifyOTPResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_api_proto_otp_proto protoreflect.FileDescriptor

var file_api_proto_otp_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x74, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x0a, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0x51, 0x0a, 0x0b, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x64, 0x0a, 0x10, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x4f, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x75, 0x0a, 0x11,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x32, 0x5e, 0x0a, 0x03, 0x4f, 0x54, 0x50, 0x12, 0x23, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x4f, 0x54, 0x50, 0x12, 0x0b, 0x2e, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0c, 0x2e, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x09, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x54, 0x50, 0x12, 0x11, 0x2e, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x73, 0x33, 0x30, 0x31, 0x2d, 0x69, 0x74, 0x73, 0x61, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2d, 0x32, 0x30, 0x32, 0x33, 0x2d, 0x32, 0x34, 0x74, 0x31, 0x2d, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x32, 0x30, 0x32, 0x33, 0x2d, 0x32, 0x34, 0x74, 0x31,
	0x2d, 0x67, 0x32, 0x2d, 0x74, 0x31, 0x2f, 0x6f, 0x74, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_proto_otp_proto_rawDescOnce sync.Once
	file_api_proto_otp_proto_rawDescData = file_api_proto_otp_proto_rawDesc
)

func file_api_proto_otp_proto_rawDescGZIP() []byte {
	file_api_proto_otp_proto_rawDescOnce.Do(func() {
		file_api_proto_otp_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_otp_proto_rawDescData)
	})
	return file_api_proto_otp_proto_rawDescData
}

var file_api_proto_otp_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_proto_otp_proto_goTypes = []interface{}{
	(*OTPRequest)(nil),        // 0: OTPRequest
	(*OTPResponse)(nil),       // 1: OTPResponse
	(*VerifyOTPRequest)(nil),  // 2: VerifyOTPRequest
	(*VerifyOTPResponse)(nil), // 3: VerifyOTPResponse
}
var file_api_proto_otp_proto_depIdxs = []int32{
	0, // 0: OTP.GetOTP:input_type -> OTPRequest
	2, // 1: OTP.VerifyOTP:input_type -> VerifyOTPRequest
	1, // 2: OTP.GetOTP:output_type -> OTPResponse
	3, // 3: OTP.VerifyOTP:output_type -> VerifyOTPResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_otp_proto_init() }
func file_api_proto_otp_proto_init() {
	if File_api_proto_otp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_otp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OTPRequest); i {
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
		file_api_proto_otp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OTPResponse); i {
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
		file_api_proto_otp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyOTPRequest); i {
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
		file_api_proto_otp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyOTPResponse); i {
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
			RawDescriptor: file_api_proto_otp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_otp_proto_goTypes,
		DependencyIndexes: file_api_proto_otp_proto_depIdxs,
		MessageInfos:      file_api_proto_otp_proto_msgTypes,
	}.Build()
	File_api_proto_otp_proto = out.File
	file_api_proto_otp_proto_rawDesc = nil
	file_api_proto_otp_proto_goTypes = nil
	file_api_proto_otp_proto_depIdxs = nil
}
