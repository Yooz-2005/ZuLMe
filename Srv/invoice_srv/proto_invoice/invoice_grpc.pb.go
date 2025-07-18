// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: invoice.proto

package invoice

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
	Invoice_GenerateInvoice_FullMethodName     = "/invoice.Invoice/GenerateInvoice"
	Invoice_ApplyInvoiceForUser_FullMethodName = "/invoice.Invoice/ApplyInvoiceForUser"
)

// InvoiceClient is the client API for Invoice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InvoiceClient interface {
	GenerateInvoice(ctx context.Context, in *GenerateInvoiceRequest, opts ...grpc.CallOption) (*GenerateInvoiceResponse, error)
	ApplyInvoiceForUser(ctx context.Context, in *ApplyInvoiceForUserRequest, opts ...grpc.CallOption) (*GenerateInvoiceResponse, error)
}

type invoiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInvoiceClient(cc grpc.ClientConnInterface) InvoiceClient {
	return &invoiceClient{cc}
}

func (c *invoiceClient) GenerateInvoice(ctx context.Context, in *GenerateInvoiceRequest, opts ...grpc.CallOption) (*GenerateInvoiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateInvoiceResponse)
	err := c.cc.Invoke(ctx, Invoice_GenerateInvoice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceClient) ApplyInvoiceForUser(ctx context.Context, in *ApplyInvoiceForUserRequest, opts ...grpc.CallOption) (*GenerateInvoiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateInvoiceResponse)
	err := c.cc.Invoke(ctx, Invoice_ApplyInvoiceForUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvoiceServer is the server API for Invoice service.
// All implementations must embed UnimplementedInvoiceServer
// for forward compatibility.
type InvoiceServer interface {
	GenerateInvoice(context.Context, *GenerateInvoiceRequest) (*GenerateInvoiceResponse, error)
	ApplyInvoiceForUser(context.Context, *ApplyInvoiceForUserRequest) (*GenerateInvoiceResponse, error)
	mustEmbedUnimplementedInvoiceServer()
}

// UnimplementedInvoiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInvoiceServer struct{}

func (UnimplementedInvoiceServer) GenerateInvoice(context.Context, *GenerateInvoiceRequest) (*GenerateInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateInvoice not implemented")
}
func (UnimplementedInvoiceServer) ApplyInvoiceForUser(context.Context, *ApplyInvoiceForUserRequest) (*GenerateInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyInvoiceForUser not implemented")
}
func (UnimplementedInvoiceServer) mustEmbedUnimplementedInvoiceServer() {}
func (UnimplementedInvoiceServer) testEmbeddedByValue()                 {}

// UnsafeInvoiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InvoiceServer will
// result in compilation errors.
type UnsafeInvoiceServer interface {
	mustEmbedUnimplementedInvoiceServer()
}

func RegisterInvoiceServer(s grpc.ServiceRegistrar, srv InvoiceServer) {
	// If the following call pancis, it indicates UnimplementedInvoiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Invoice_ServiceDesc, srv)
}

func _Invoice_GenerateInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServer).GenerateInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Invoice_GenerateInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServer).GenerateInvoice(ctx, req.(*GenerateInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invoice_ApplyInvoiceForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyInvoiceForUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServer).ApplyInvoiceForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Invoice_ApplyInvoiceForUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServer).ApplyInvoiceForUser(ctx, req.(*ApplyInvoiceForUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Invoice_ServiceDesc is the grpc.ServiceDesc for Invoice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Invoice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "invoice.Invoice",
	HandlerType: (*InvoiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateInvoice",
			Handler:    _Invoice_GenerateInvoice_Handler,
		},
		{
			MethodName: "ApplyInvoiceForUser",
			Handler:    _Invoice_ApplyInvoiceForUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "invoice.proto",
}
