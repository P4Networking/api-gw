// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// PiscServiceClient is the client API for PiscService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PiscServiceClient interface {
	LoginAuth(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	GetSystemInformation(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SystemInformationResponse, error)
	GetSystemCPUInformation(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SystemCPUInformationResponse, error)
	GetSystemRAMInformation(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SystemRAMInformationResponse, error)
	GetSystemDiskInformation(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SystemDiskInformationResponse, error)
	GetFeature(ctx context.Context, in *NameMessage, opts ...grpc.CallOption) (*FeatureResponse, error)
	GetPort(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*PortResponse, error)
	GetLicense(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*LicenseResponse, error)
}

type piscServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPiscServiceClient(cc grpc.ClientConnInterface) PiscServiceClient {
	return &piscServiceClient{cc}
}

func (c *piscServiceClient) LoginAuth(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/proto.PiscService/LoginAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piscServiceClient) GetSystemInformation(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SystemInformationResponse, error) {
	out := new(SystemInformationResponse)
	err := c.cc.Invoke(ctx, "/proto.PiscService/GetSystemInformation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piscServiceClient) GetSystemCPUInformation(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SystemCPUInformationResponse, error) {
	out := new(SystemCPUInformationResponse)
	err := c.cc.Invoke(ctx, "/proto.PiscService/GetSystemCPUInformation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piscServiceClient) GetSystemRAMInformation(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SystemRAMInformationResponse, error) {
	out := new(SystemRAMInformationResponse)
	err := c.cc.Invoke(ctx, "/proto.PiscService/GetSystemRAMInformation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piscServiceClient) GetSystemDiskInformation(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SystemDiskInformationResponse, error) {
	out := new(SystemDiskInformationResponse)
	err := c.cc.Invoke(ctx, "/proto.PiscService/GetSystemDiskInformation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piscServiceClient) GetFeature(ctx context.Context, in *NameMessage, opts ...grpc.CallOption) (*FeatureResponse, error) {
	out := new(FeatureResponse)
	err := c.cc.Invoke(ctx, "/proto.PiscService/GetFeature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piscServiceClient) GetPort(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*PortResponse, error) {
	out := new(PortResponse)
	err := c.cc.Invoke(ctx, "/proto.PiscService/GetPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piscServiceClient) GetLicense(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*LicenseResponse, error) {
	out := new(LicenseResponse)
	err := c.cc.Invoke(ctx, "/proto.PiscService/GetLicense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PiscServiceServer is the server API for PiscService service.
// All implementations must embed UnimplementedPiscServiceServer
// for forward compatibility
type PiscServiceServer interface {
	LoginAuth(context.Context, *LoginRequest) (*LoginResponse, error)
	GetSystemInformation(context.Context, *emptypb.Empty) (*SystemInformationResponse, error)
	GetSystemCPUInformation(context.Context, *emptypb.Empty) (*SystemCPUInformationResponse, error)
	GetSystemRAMInformation(context.Context, *emptypb.Empty) (*SystemRAMInformationResponse, error)
	GetSystemDiskInformation(context.Context, *emptypb.Empty) (*SystemDiskInformationResponse, error)
	GetFeature(context.Context, *NameMessage) (*FeatureResponse, error)
	GetPort(context.Context, *IdMessage) (*PortResponse, error)
	GetLicense(context.Context, *emptypb.Empty) (*LicenseResponse, error)
	mustEmbedUnimplementedPiscServiceServer()
}

// UnimplementedPiscServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPiscServiceServer struct {
}

func (UnimplementedPiscServiceServer) LoginAuth(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginAuth not implemented")
}
func (UnimplementedPiscServiceServer) GetSystemInformation(context.Context, *emptypb.Empty) (*SystemInformationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemInformation not implemented")
}
func (UnimplementedPiscServiceServer) GetSystemCPUInformation(context.Context, *emptypb.Empty) (*SystemCPUInformationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemCPUInformation not implemented")
}
func (UnimplementedPiscServiceServer) GetSystemRAMInformation(context.Context, *emptypb.Empty) (*SystemRAMInformationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemRAMInformation not implemented")
}
func (UnimplementedPiscServiceServer) GetSystemDiskInformation(context.Context, *emptypb.Empty) (*SystemDiskInformationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemDiskInformation not implemented")
}
func (UnimplementedPiscServiceServer) GetFeature(context.Context, *NameMessage) (*FeatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeature not implemented")
}
func (UnimplementedPiscServiceServer) GetPort(context.Context, *IdMessage) (*PortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPort not implemented")
}
func (UnimplementedPiscServiceServer) GetLicense(context.Context, *emptypb.Empty) (*LicenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLicense not implemented")
}
func (UnimplementedPiscServiceServer) mustEmbedUnimplementedPiscServiceServer() {}

// UnsafePiscServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PiscServiceServer will
// result in compilation errors.
type UnsafePiscServiceServer interface {
	mustEmbedUnimplementedPiscServiceServer()
}

func RegisterPiscServiceServer(s grpc.ServiceRegistrar, srv PiscServiceServer) {
	s.RegisterService(&_PiscService_serviceDesc, srv)
}

func _PiscService_LoginAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiscServiceServer).LoginAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PiscService/LoginAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiscServiceServer).LoginAuth(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PiscService_GetSystemInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiscServiceServer).GetSystemInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PiscService/GetSystemInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiscServiceServer).GetSystemInformation(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PiscService_GetSystemCPUInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiscServiceServer).GetSystemCPUInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PiscService/GetSystemCPUInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiscServiceServer).GetSystemCPUInformation(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PiscService_GetSystemRAMInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiscServiceServer).GetSystemRAMInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PiscService/GetSystemRAMInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiscServiceServer).GetSystemRAMInformation(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PiscService_GetSystemDiskInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiscServiceServer).GetSystemDiskInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PiscService/GetSystemDiskInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiscServiceServer).GetSystemDiskInformation(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PiscService_GetFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiscServiceServer).GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PiscService/GetFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiscServiceServer).GetFeature(ctx, req.(*NameMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _PiscService_GetPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiscServiceServer).GetPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PiscService/GetPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiscServiceServer).GetPort(ctx, req.(*IdMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _PiscService_GetLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiscServiceServer).GetLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PiscService/GetLicense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiscServiceServer).GetLicense(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _PiscService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PiscService",
	HandlerType: (*PiscServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginAuth",
			Handler:    _PiscService_LoginAuth_Handler,
		},
		{
			MethodName: "GetSystemInformation",
			Handler:    _PiscService_GetSystemInformation_Handler,
		},
		{
			MethodName: "GetSystemCPUInformation",
			Handler:    _PiscService_GetSystemCPUInformation_Handler,
		},
		{
			MethodName: "GetSystemRAMInformation",
			Handler:    _PiscService_GetSystemRAMInformation_Handler,
		},
		{
			MethodName: "GetSystemDiskInformation",
			Handler:    _PiscService_GetSystemDiskInformation_Handler,
		},
		{
			MethodName: "GetFeature",
			Handler:    _PiscService_GetFeature_Handler,
		},
		{
			MethodName: "GetPort",
			Handler:    _PiscService_GetPort_Handler,
		},
		{
			MethodName: "GetLicense",
			Handler:    _PiscService_GetLicense_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pisc.proto",
}