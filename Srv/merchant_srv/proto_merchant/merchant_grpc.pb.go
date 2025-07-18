// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: merchant.proto

package merchant

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MerchantService_MerchantRegister_FullMethodName = "/merchant.MerchantService/MerchantRegister"
	MerchantService_MerchantLogin_FullMethodName    = "/merchant.MerchantService/MerchantLogin"
	MerchantService_GetLocationList_FullMethodName  = "/merchant.MerchantService/GetLocationList"
)

// MerchantServiceClient is the client API for MerchantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 商家服务
type MerchantServiceClient interface {
	// 商家注册
	MerchantRegister(ctx context.Context, in *MerchantRegisterRequest, opts ...grpc.CallOption) (*MerchantRegisterResponse, error)
	// 商家登录
	MerchantLogin(ctx context.Context, in *MerchantLoginRequest, opts ...grpc.CallOption) (*MerchantLoginResponse, error)
	// 获取网点列表
	GetLocationList(ctx context.Context, in *GetLocationListRequest, opts ...grpc.CallOption) (*GetLocationListResponse, error)
}

type merchantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMerchantServiceClient(cc grpc.ClientConnInterface) MerchantServiceClient {
	return &merchantServiceClient{cc}
}

func (c *merchantServiceClient) MerchantRegister(ctx context.Context, in *MerchantRegisterRequest, opts ...grpc.CallOption) (*MerchantRegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MerchantRegisterResponse)
	err := c.cc.Invoke(ctx, MerchantService_MerchantRegister_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) MerchantLogin(ctx context.Context, in *MerchantLoginRequest, opts ...grpc.CallOption) (*MerchantLoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MerchantLoginResponse)
	err := c.cc.Invoke(ctx, MerchantService_MerchantLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) GetLocationList(ctx context.Context, in *GetLocationListRequest, opts ...grpc.CallOption) (*GetLocationListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLocationListResponse)
	err := c.cc.Invoke(ctx, MerchantService_GetLocationList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MerchantServiceServer is the server API for MerchantService service.
// All implementations must embed UnimplementedMerchantServiceServer
// for forward compatibility.
//
// 商家服务
type MerchantServiceServer interface {
	// 商家注册
	MerchantRegister(context.Context, *MerchantRegisterRequest) (*MerchantRegisterResponse, error)
	// 商家登录
	MerchantLogin(context.Context, *MerchantLoginRequest) (*MerchantLoginResponse, error)
	// 获取网点列表
	GetLocationList(context.Context, *GetLocationListRequest) (*GetLocationListResponse, error)
	mustEmbedUnimplementedMerchantServiceServer()
}

// UnimplementedMerchantServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMerchantServiceServer struct{}

func (UnimplementedMerchantServiceServer) MerchantRegister(context.Context, *MerchantRegisterRequest) (*MerchantRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MerchantRegister not implemented")
}
func (UnimplementedMerchantServiceServer) MerchantLogin(context.Context, *MerchantLoginRequest) (*MerchantLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MerchantLogin not implemented")
}
func (UnimplementedMerchantServiceServer) GetLocationList(context.Context, *GetLocationListRequest) (*GetLocationListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocationList not implemented")
}
func (UnimplementedMerchantServiceServer) mustEmbedUnimplementedMerchantServiceServer() {}
func (UnimplementedMerchantServiceServer) testEmbeddedByValue()                         {}

// UnsafeMerchantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MerchantServiceServer will
// result in compilation errors.
type UnsafeMerchantServiceServer interface {
	mustEmbedUnimplementedMerchantServiceServer()
}

func RegisterMerchantServiceServer(s grpc.ServiceRegistrar, srv MerchantServiceServer) {
	// If the following call pancis, it indicates UnimplementedMerchantServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MerchantService_ServiceDesc, srv)
}

func _MerchantService_MerchantRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).MerchantRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MerchantService_MerchantRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).MerchantRegister(ctx, req.(*MerchantRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_MerchantLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).MerchantLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MerchantService_MerchantLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).MerchantLogin(ctx, req.(*MerchantLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_GetLocationList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocationListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).GetLocationList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MerchantService_GetLocationList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).GetLocationList(ctx, req.(*GetLocationListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MerchantService_ServiceDesc is the grpc.ServiceDesc for MerchantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MerchantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "merchant.MerchantService",
	HandlerType: (*MerchantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MerchantRegister",
			Handler:    _MerchantService_MerchantRegister_Handler,
		},
		{
			MethodName: "MerchantLogin",
			Handler:    _MerchantService_MerchantLogin_Handler,
		},
		{
			MethodName: "GetLocationList",
			Handler:    _MerchantService_GetLocationList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "merchant.proto",
}
