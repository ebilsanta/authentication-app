// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: external/proto/otp.proto

package otp

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OTP_GetOTP_FullMethodName    = "/OTP/GetOTP"
	OTP_VerifyOTP_FullMethodName = "/OTP/VerifyOTP"
)

// OTPClient is the client API for OTP service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OTPClient interface {
	GetOTP(ctx context.Context, in *OTPRequest, opts ...grpc.CallOption) (*OTPResponse, error)
	VerifyOTP(ctx context.Context, in *VerifyOTPRequest, opts ...grpc.CallOption) (*VerifyOTPResponse, error)
}

type oTPClient struct {
	cc grpc.ClientConnInterface
}

func NewOTPClient(cc grpc.ClientConnInterface) OTPClient {
	return &oTPClient{cc}
}

func (c *oTPClient) GetOTP(ctx context.Context, in *OTPRequest, opts ...grpc.CallOption) (*OTPResponse, error) {
	out := new(OTPResponse)
	err := c.cc.Invoke(ctx, OTP_GetOTP_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oTPClient) VerifyOTP(ctx context.Context, in *VerifyOTPRequest, opts ...grpc.CallOption) (*VerifyOTPResponse, error) {
	out := new(VerifyOTPResponse)
	err := c.cc.Invoke(ctx, OTP_VerifyOTP_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OTPServer is the server API for OTP service.
// All implementations must embed UnimplementedOTPServer
// for forward compatibility
type OTPServer interface {
	GetOTP(context.Context, *OTPRequest) (*OTPResponse, error)
	VerifyOTP(context.Context, *VerifyOTPRequest) (*VerifyOTPResponse, error)
	mustEmbedUnimplementedOTPServer()
}

// UnimplementedOTPServer must be embedded to have forward compatible implementations.
type UnimplementedOTPServer struct {
}

func (UnimplementedOTPServer) GetOTP(context.Context, *OTPRequest) (*OTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOTP not implemented")
}
func (UnimplementedOTPServer) VerifyOTP(context.Context, *VerifyOTPRequest) (*VerifyOTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyOTP not implemented")
}
func (UnimplementedOTPServer) mustEmbedUnimplementedOTPServer() {}

// UnsafeOTPServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OTPServer will
// result in compilation errors.
type UnsafeOTPServer interface {
	mustEmbedUnimplementedOTPServer()
}

func RegisterOTPServer(s grpc.ServiceRegistrar, srv OTPServer) {
	s.RegisterService(&OTP_ServiceDesc, srv)
}

func _OTP_GetOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OTPServer).GetOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OTP_GetOTP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OTPServer).GetOTP(ctx, req.(*OTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OTP_VerifyOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OTPServer).VerifyOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OTP_VerifyOTP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OTPServer).VerifyOTP(ctx, req.(*VerifyOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OTP_ServiceDesc is the grpc.ServiceDesc for OTP service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OTP_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OTP",
	HandlerType: (*OTPServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOTP",
			Handler:    _OTP_GetOTP_Handler,
		},
		{
			MethodName: "VerifyOTP",
			Handler:    _OTP_VerifyOTP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "external/proto/otp.proto",
}
