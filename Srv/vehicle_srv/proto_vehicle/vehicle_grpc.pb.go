// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: vehicle.proto

package vehicle

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
	Vehicle_Ping_FullMethodName              = "/vehicle.Vehicle/Ping"
	Vehicle_CreateVehicle_FullMethodName     = "/vehicle.Vehicle/CreateVehicle"
	Vehicle_UpdateVehicle_FullMethodName     = "/vehicle.Vehicle/UpdateVehicle"
	Vehicle_DeleteVehicle_FullMethodName     = "/vehicle.Vehicle/DeleteVehicle"
	Vehicle_GetVehicle_FullMethodName        = "/vehicle.Vehicle/GetVehicle"
	Vehicle_ListVehicles_FullMethodName      = "/vehicle.Vehicle/ListVehicles"
	Vehicle_CreateVehicleType_FullMethodName = "/vehicle.Vehicle/CreateVehicleType"
	Vehicle_UpdateVehicleType_FullMethodName = "/vehicle.Vehicle/UpdateVehicleType"
	Vehicle_DeleteVehicleType_FullMethodName = "/vehicle.Vehicle/DeleteVehicleType"
	Vehicle_GetVehicleType_FullMethodName    = "/vehicle.Vehicle/GetVehicleType"
	Vehicle_ListVehicleTypes_FullMethodName  = "/vehicle.Vehicle/ListVehicleTypes"
)

// VehicleClient is the client API for Vehicle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VehicleClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	// 车辆管理
	CreateVehicle(ctx context.Context, in *CreateVehicleRequest, opts ...grpc.CallOption) (*CreateVehicleResponse, error)
	UpdateVehicle(ctx context.Context, in *UpdateVehicleRequest, opts ...grpc.CallOption) (*UpdateVehicleResponse, error)
	DeleteVehicle(ctx context.Context, in *DeleteVehicleRequest, opts ...grpc.CallOption) (*DeleteVehicleResponse, error)
	GetVehicle(ctx context.Context, in *GetVehicleRequest, opts ...grpc.CallOption) (*GetVehicleResponse, error)
	ListVehicles(ctx context.Context, in *ListVehiclesRequest, opts ...grpc.CallOption) (*ListVehiclesResponse, error)
	// 车辆类型管理
	CreateVehicleType(ctx context.Context, in *CreateVehicleTypeRequest, opts ...grpc.CallOption) (*CreateVehicleTypeResponse, error)
	UpdateVehicleType(ctx context.Context, in *UpdateVehicleTypeRequest, opts ...grpc.CallOption) (*UpdateVehicleTypeResponse, error)
	DeleteVehicleType(ctx context.Context, in *DeleteVehicleTypeRequest, opts ...grpc.CallOption) (*DeleteVehicleTypeResponse, error)
	GetVehicleType(ctx context.Context, in *GetVehicleTypeRequest, opts ...grpc.CallOption) (*GetVehicleTypeResponse, error)
	ListVehicleTypes(ctx context.Context, in *ListVehicleTypesRequest, opts ...grpc.CallOption) (*ListVehicleTypesResponse, error)
}

type vehicleClient struct {
	cc grpc.ClientConnInterface
}

func NewVehicleClient(cc grpc.ClientConnInterface) VehicleClient {
	return &vehicleClient{cc}
}

func (c *vehicleClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Vehicle_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleClient) CreateVehicle(ctx context.Context, in *CreateVehicleRequest, opts ...grpc.CallOption) (*CreateVehicleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateVehicleResponse)
	err := c.cc.Invoke(ctx, Vehicle_CreateVehicle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleClient) UpdateVehicle(ctx context.Context, in *UpdateVehicleRequest, opts ...grpc.CallOption) (*UpdateVehicleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateVehicleResponse)
	err := c.cc.Invoke(ctx, Vehicle_UpdateVehicle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleClient) DeleteVehicle(ctx context.Context, in *DeleteVehicleRequest, opts ...grpc.CallOption) (*DeleteVehicleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteVehicleResponse)
	err := c.cc.Invoke(ctx, Vehicle_DeleteVehicle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleClient) GetVehicle(ctx context.Context, in *GetVehicleRequest, opts ...grpc.CallOption) (*GetVehicleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVehicleResponse)
	err := c.cc.Invoke(ctx, Vehicle_GetVehicle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleClient) ListVehicles(ctx context.Context, in *ListVehiclesRequest, opts ...grpc.CallOption) (*ListVehiclesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListVehiclesResponse)
	err := c.cc.Invoke(ctx, Vehicle_ListVehicles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleClient) CreateVehicleType(ctx context.Context, in *CreateVehicleTypeRequest, opts ...grpc.CallOption) (*CreateVehicleTypeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateVehicleTypeResponse)
	err := c.cc.Invoke(ctx, Vehicle_CreateVehicleType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleClient) UpdateVehicleType(ctx context.Context, in *UpdateVehicleTypeRequest, opts ...grpc.CallOption) (*UpdateVehicleTypeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateVehicleTypeResponse)
	err := c.cc.Invoke(ctx, Vehicle_UpdateVehicleType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleClient) DeleteVehicleType(ctx context.Context, in *DeleteVehicleTypeRequest, opts ...grpc.CallOption) (*DeleteVehicleTypeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteVehicleTypeResponse)
	err := c.cc.Invoke(ctx, Vehicle_DeleteVehicleType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleClient) GetVehicleType(ctx context.Context, in *GetVehicleTypeRequest, opts ...grpc.CallOption) (*GetVehicleTypeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVehicleTypeResponse)
	err := c.cc.Invoke(ctx, Vehicle_GetVehicleType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleClient) ListVehicleTypes(ctx context.Context, in *ListVehicleTypesRequest, opts ...grpc.CallOption) (*ListVehicleTypesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListVehicleTypesResponse)
	err := c.cc.Invoke(ctx, Vehicle_ListVehicleTypes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VehicleServer is the server API for Vehicle service.
// All implementations must embed UnimplementedVehicleServer
// for forward compatibility.
type VehicleServer interface {
	Ping(context.Context, *Request) (*Response, error)
	// 车辆管理
	CreateVehicle(context.Context, *CreateVehicleRequest) (*CreateVehicleResponse, error)
	UpdateVehicle(context.Context, *UpdateVehicleRequest) (*UpdateVehicleResponse, error)
	DeleteVehicle(context.Context, *DeleteVehicleRequest) (*DeleteVehicleResponse, error)
	GetVehicle(context.Context, *GetVehicleRequest) (*GetVehicleResponse, error)
	ListVehicles(context.Context, *ListVehiclesRequest) (*ListVehiclesResponse, error)
	// 车辆类型管理
	CreateVehicleType(context.Context, *CreateVehicleTypeRequest) (*CreateVehicleTypeResponse, error)
	UpdateVehicleType(context.Context, *UpdateVehicleTypeRequest) (*UpdateVehicleTypeResponse, error)
	DeleteVehicleType(context.Context, *DeleteVehicleTypeRequest) (*DeleteVehicleTypeResponse, error)
	GetVehicleType(context.Context, *GetVehicleTypeRequest) (*GetVehicleTypeResponse, error)
	ListVehicleTypes(context.Context, *ListVehicleTypesRequest) (*ListVehicleTypesResponse, error)
	mustEmbedUnimplementedVehicleServer()
}

// UnimplementedVehicleServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVehicleServer struct{}

func (UnimplementedVehicleServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedVehicleServer) CreateVehicle(context.Context, *CreateVehicleRequest) (*CreateVehicleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVehicle not implemented")
}
func (UnimplementedVehicleServer) UpdateVehicle(context.Context, *UpdateVehicleRequest) (*UpdateVehicleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVehicle not implemented")
}
func (UnimplementedVehicleServer) DeleteVehicle(context.Context, *DeleteVehicleRequest) (*DeleteVehicleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVehicle not implemented")
}
func (UnimplementedVehicleServer) GetVehicle(context.Context, *GetVehicleRequest) (*GetVehicleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVehicle not implemented")
}
func (UnimplementedVehicleServer) ListVehicles(context.Context, *ListVehiclesRequest) (*ListVehiclesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVehicles not implemented")
}
func (UnimplementedVehicleServer) CreateVehicleType(context.Context, *CreateVehicleTypeRequest) (*CreateVehicleTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVehicleType not implemented")
}
func (UnimplementedVehicleServer) UpdateVehicleType(context.Context, *UpdateVehicleTypeRequest) (*UpdateVehicleTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVehicleType not implemented")
}
func (UnimplementedVehicleServer) DeleteVehicleType(context.Context, *DeleteVehicleTypeRequest) (*DeleteVehicleTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVehicleType not implemented")
}
func (UnimplementedVehicleServer) GetVehicleType(context.Context, *GetVehicleTypeRequest) (*GetVehicleTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVehicleType not implemented")
}
func (UnimplementedVehicleServer) ListVehicleTypes(context.Context, *ListVehicleTypesRequest) (*ListVehicleTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVehicleTypes not implemented")
}
func (UnimplementedVehicleServer) mustEmbedUnimplementedVehicleServer() {}
func (UnimplementedVehicleServer) testEmbeddedByValue()                 {}

// UnsafeVehicleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VehicleServer will
// result in compilation errors.
type UnsafeVehicleServer interface {
	mustEmbedUnimplementedVehicleServer()
}

func RegisterVehicleServer(s grpc.ServiceRegistrar, srv VehicleServer) {
	// If the following call pancis, it indicates UnimplementedVehicleServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Vehicle_ServiceDesc, srv)
}

func _Vehicle_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Vehicle_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vehicle_CreateVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVehicleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServer).CreateVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Vehicle_CreateVehicle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServer).CreateVehicle(ctx, req.(*CreateVehicleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vehicle_UpdateVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVehicleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServer).UpdateVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Vehicle_UpdateVehicle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServer).UpdateVehicle(ctx, req.(*UpdateVehicleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vehicle_DeleteVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVehicleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServer).DeleteVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Vehicle_DeleteVehicle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServer).DeleteVehicle(ctx, req.(*DeleteVehicleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vehicle_GetVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVehicleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServer).GetVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Vehicle_GetVehicle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServer).GetVehicle(ctx, req.(*GetVehicleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vehicle_ListVehicles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVehiclesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServer).ListVehicles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Vehicle_ListVehicles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServer).ListVehicles(ctx, req.(*ListVehiclesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vehicle_CreateVehicleType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVehicleTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServer).CreateVehicleType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Vehicle_CreateVehicleType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServer).CreateVehicleType(ctx, req.(*CreateVehicleTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vehicle_UpdateVehicleType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVehicleTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServer).UpdateVehicleType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Vehicle_UpdateVehicleType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServer).UpdateVehicleType(ctx, req.(*UpdateVehicleTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vehicle_DeleteVehicleType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVehicleTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServer).DeleteVehicleType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Vehicle_DeleteVehicleType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServer).DeleteVehicleType(ctx, req.(*DeleteVehicleTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vehicle_GetVehicleType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVehicleTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServer).GetVehicleType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Vehicle_GetVehicleType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServer).GetVehicleType(ctx, req.(*GetVehicleTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vehicle_ListVehicleTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVehicleTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServer).ListVehicleTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Vehicle_ListVehicleTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServer).ListVehicleTypes(ctx, req.(*ListVehicleTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Vehicle_ServiceDesc is the grpc.ServiceDesc for Vehicle service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Vehicle_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vehicle.Vehicle",
	HandlerType: (*VehicleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Vehicle_Ping_Handler,
		},
		{
			MethodName: "CreateVehicle",
			Handler:    _Vehicle_CreateVehicle_Handler,
		},
		{
			MethodName: "UpdateVehicle",
			Handler:    _Vehicle_UpdateVehicle_Handler,
		},
		{
			MethodName: "DeleteVehicle",
			Handler:    _Vehicle_DeleteVehicle_Handler,
		},
		{
			MethodName: "GetVehicle",
			Handler:    _Vehicle_GetVehicle_Handler,
		},
		{
			MethodName: "ListVehicles",
			Handler:    _Vehicle_ListVehicles_Handler,
		},
		{
			MethodName: "CreateVehicleType",
			Handler:    _Vehicle_CreateVehicleType_Handler,
		},
		{
			MethodName: "UpdateVehicleType",
			Handler:    _Vehicle_UpdateVehicleType_Handler,
		},
		{
			MethodName: "DeleteVehicleType",
			Handler:    _Vehicle_DeleteVehicleType_Handler,
		},
		{
			MethodName: "GetVehicleType",
			Handler:    _Vehicle_GetVehicleType_Handler,
		},
		{
			MethodName: "ListVehicleTypes",
			Handler:    _Vehicle_ListVehicleTypes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vehicle.proto",
}
