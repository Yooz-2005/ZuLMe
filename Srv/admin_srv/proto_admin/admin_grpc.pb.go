// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: admin.proto

package admin

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
	Admin_MerchantApprove_FullMethodName = "/admin.Admin/MerchantApprove"
	Admin_MerchantUpdate_FullMethodName  = "/admin.Admin/MerchantUpdate"
	Admin_MerchantDelete_FullMethodName  = "/admin.Admin/MerchantDelete"
	Admin_MerchantList_FullMethodName    = "/admin.Admin/MerchantList"
	Admin_MerchantDetail_FullMethodName  = "/admin.Admin/MerchantDetail"
)

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	// 商户管理
	MerchantApprove(ctx context.Context, in *MerchantApproveRequest, opts ...grpc.CallOption) (*MerchantApproveResponse, error)
	MerchantUpdate(ctx context.Context, in *MerchantUpdateRequest, opts ...grpc.CallOption) (*MerchantUpdateResponse, error)
	MerchantDelete(ctx context.Context, in *MerchantDeleteRequest, opts ...grpc.CallOption) (*MerchantDeleteResponse, error)
	MerchantList(ctx context.Context, in *MerchantListRequest, opts ...grpc.CallOption) (*MerchantListResponse, error)
	MerchantDetail(ctx context.Context, in *MerchantDetailRequest, opts ...grpc.CallOption) (*MerchantDetailResponse, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) MerchantApprove(ctx context.Context, in *MerchantApproveRequest, opts ...grpc.CallOption) (*MerchantApproveResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MerchantApproveResponse)
	err := c.cc.Invoke(ctx, Admin_MerchantApprove_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) MerchantUpdate(ctx context.Context, in *MerchantUpdateRequest, opts ...grpc.CallOption) (*MerchantUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MerchantUpdateResponse)
	err := c.cc.Invoke(ctx, Admin_MerchantUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) MerchantDelete(ctx context.Context, in *MerchantDeleteRequest, opts ...grpc.CallOption) (*MerchantDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MerchantDeleteResponse)
	err := c.cc.Invoke(ctx, Admin_MerchantDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) MerchantList(ctx context.Context, in *MerchantListRequest, opts ...grpc.CallOption) (*MerchantListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MerchantListResponse)
	err := c.cc.Invoke(ctx, Admin_MerchantList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) MerchantDetail(ctx context.Context, in *MerchantDetailRequest, opts ...grpc.CallOption) (*MerchantDetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MerchantDetailResponse)
	err := c.cc.Invoke(ctx, Admin_MerchantDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility.
type AdminServer interface {
	// 商户管理
	MerchantApprove(context.Context, *MerchantApproveRequest) (*MerchantApproveResponse, error)
	MerchantUpdate(context.Context, *MerchantUpdateRequest) (*MerchantUpdateResponse, error)
	MerchantDelete(context.Context, *MerchantDeleteRequest) (*MerchantDeleteResponse, error)
	MerchantList(context.Context, *MerchantListRequest) (*MerchantListResponse, error)
	MerchantDetail(context.Context, *MerchantDetailRequest) (*MerchantDetailResponse, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAdminServer struct{}

func (UnimplementedAdminServer) MerchantApprove(context.Context, *MerchantApproveRequest) (*MerchantApproveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MerchantApprove not implemented")
}
func (UnimplementedAdminServer) MerchantUpdate(context.Context, *MerchantUpdateRequest) (*MerchantUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MerchantUpdate not implemented")
}
func (UnimplementedAdminServer) MerchantDelete(context.Context, *MerchantDeleteRequest) (*MerchantDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MerchantDelete not implemented")
}
func (UnimplementedAdminServer) MerchantList(context.Context, *MerchantListRequest) (*MerchantListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MerchantList not implemented")
}
func (UnimplementedAdminServer) MerchantDetail(context.Context, *MerchantDetailRequest) (*MerchantDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MerchantDetail not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}
func (UnimplementedAdminServer) testEmbeddedByValue()               {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	// If the following call pancis, it indicates UnimplementedAdminServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_MerchantApprove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantApproveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).MerchantApprove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_MerchantApprove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).MerchantApprove(ctx, req.(*MerchantApproveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_MerchantUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).MerchantUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_MerchantUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).MerchantUpdate(ctx, req.(*MerchantUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_MerchantDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).MerchantDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_MerchantDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).MerchantDelete(ctx, req.(*MerchantDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_MerchantList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).MerchantList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_MerchantList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).MerchantList(ctx, req.(*MerchantListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_MerchantDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).MerchantDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_MerchantDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).MerchantDetail(ctx, req.(*MerchantDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MerchantApprove",
			Handler:    _Admin_MerchantApprove_Handler,
		},
		{
			MethodName: "MerchantUpdate",
			Handler:    _Admin_MerchantUpdate_Handler,
		},
		{
			MethodName: "MerchantDelete",
			Handler:    _Admin_MerchantDelete_Handler,
		},
		{
			MethodName: "MerchantList",
			Handler:    _Admin_MerchantList_Handler,
		},
		{
			MethodName: "MerchantDetail",
			Handler:    _Admin_MerchantDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}
