// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: proto_comment/comment.proto

package proto_comment

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
	CommentService_CreateComment_FullMethodName       = "/comment.CommentService/CreateComment"
	CommentService_GetComment_FullMethodName          = "/comment.CommentService/GetComment"
	CommentService_GetOrderComment_FullMethodName     = "/comment.CommentService/GetOrderComment"
	CommentService_GetVehicleComments_FullMethodName  = "/comment.CommentService/GetVehicleComments"
	CommentService_GetUserComments_FullMethodName     = "/comment.CommentService/GetUserComments"
	CommentService_GetVehicleStats_FullMethodName     = "/comment.CommentService/GetVehicleStats"
	CommentService_UpdateComment_FullMethodName       = "/comment.CommentService/UpdateComment"
	CommentService_DeleteComment_FullMethodName       = "/comment.CommentService/DeleteComment"
	CommentService_ReplyComment_FullMethodName        = "/comment.CommentService/ReplyComment"
	CommentService_CheckOrderCommented_FullMethodName = "/comment.CommentService/CheckOrderCommented"
)

// CommentServiceClient is the client API for CommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 评论服务
type CommentServiceClient interface {
	// 创建评论
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
	// 获取评论详情
	GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentResponse, error)
	// 获取订单评论
	GetOrderComment(ctx context.Context, in *GetOrderCommentRequest, opts ...grpc.CallOption) (*GetCommentResponse, error)
	// 获取车辆评论列表
	GetVehicleComments(ctx context.Context, in *GetVehicleCommentsRequest, opts ...grpc.CallOption) (*GetVehicleCommentsResponse, error)
	// 获取用户评论列表
	GetUserComments(ctx context.Context, in *GetUserCommentsRequest, opts ...grpc.CallOption) (*GetUserCommentsResponse, error)
	// 获取车辆评论统计
	GetVehicleStats(ctx context.Context, in *GetVehicleStatsRequest, opts ...grpc.CallOption) (*GetVehicleStatsResponse, error)
	// 更新评论
	UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error)
	// 删除评论
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error)
	// 商家回复评论
	ReplyComment(ctx context.Context, in *ReplyCommentRequest, opts ...grpc.CallOption) (*ReplyCommentResponse, error)
	// 检查订单是否已评论
	CheckOrderCommented(ctx context.Context, in *CheckOrderCommentedRequest, opts ...grpc.CallOption) (*CheckOrderCommentedResponse, error)
}

type commentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentServiceClient(cc grpc.ClientConnInterface) CommentServiceClient {
	return &commentServiceClient{cc}
}

func (c *commentServiceClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCommentResponse)
	err := c.cc.Invoke(ctx, CommentService_CreateComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommentResponse)
	err := c.cc.Invoke(ctx, CommentService_GetComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetOrderComment(ctx context.Context, in *GetOrderCommentRequest, opts ...grpc.CallOption) (*GetCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommentResponse)
	err := c.cc.Invoke(ctx, CommentService_GetOrderComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetVehicleComments(ctx context.Context, in *GetVehicleCommentsRequest, opts ...grpc.CallOption) (*GetVehicleCommentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVehicleCommentsResponse)
	err := c.cc.Invoke(ctx, CommentService_GetVehicleComments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetUserComments(ctx context.Context, in *GetUserCommentsRequest, opts ...grpc.CallOption) (*GetUserCommentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserCommentsResponse)
	err := c.cc.Invoke(ctx, CommentService_GetUserComments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetVehicleStats(ctx context.Context, in *GetVehicleStatsRequest, opts ...grpc.CallOption) (*GetVehicleStatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVehicleStatsResponse)
	err := c.cc.Invoke(ctx, CommentService_GetVehicleStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCommentResponse)
	err := c.cc.Invoke(ctx, CommentService_UpdateComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCommentResponse)
	err := c.cc.Invoke(ctx, CommentService_DeleteComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) ReplyComment(ctx context.Context, in *ReplyCommentRequest, opts ...grpc.CallOption) (*ReplyCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReplyCommentResponse)
	err := c.cc.Invoke(ctx, CommentService_ReplyComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) CheckOrderCommented(ctx context.Context, in *CheckOrderCommentedRequest, opts ...grpc.CallOption) (*CheckOrderCommentedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckOrderCommentedResponse)
	err := c.cc.Invoke(ctx, CommentService_CheckOrderCommented_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServiceServer is the server API for CommentService service.
// All implementations must embed UnimplementedCommentServiceServer
// for forward compatibility.
//
// 评论服务
type CommentServiceServer interface {
	// 创建评论
	CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error)
	// 获取评论详情
	GetComment(context.Context, *GetCommentRequest) (*GetCommentResponse, error)
	// 获取订单评论
	GetOrderComment(context.Context, *GetOrderCommentRequest) (*GetCommentResponse, error)
	// 获取车辆评论列表
	GetVehicleComments(context.Context, *GetVehicleCommentsRequest) (*GetVehicleCommentsResponse, error)
	// 获取用户评论列表
	GetUserComments(context.Context, *GetUserCommentsRequest) (*GetUserCommentsResponse, error)
	// 获取车辆评论统计
	GetVehicleStats(context.Context, *GetVehicleStatsRequest) (*GetVehicleStatsResponse, error)
	// 更新评论
	UpdateComment(context.Context, *UpdateCommentRequest) (*UpdateCommentResponse, error)
	// 删除评论
	DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error)
	// 商家回复评论
	ReplyComment(context.Context, *ReplyCommentRequest) (*ReplyCommentResponse, error)
	// 检查订单是否已评论
	CheckOrderCommented(context.Context, *CheckOrderCommentedRequest) (*CheckOrderCommentedResponse, error)
	mustEmbedUnimplementedCommentServiceServer()
}

// UnimplementedCommentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCommentServiceServer struct{}

func (UnimplementedCommentServiceServer) CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedCommentServiceServer) GetComment(context.Context, *GetCommentRequest) (*GetCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComment not implemented")
}
func (UnimplementedCommentServiceServer) GetOrderComment(context.Context, *GetOrderCommentRequest) (*GetCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderComment not implemented")
}
func (UnimplementedCommentServiceServer) GetVehicleComments(context.Context, *GetVehicleCommentsRequest) (*GetVehicleCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVehicleComments not implemented")
}
func (UnimplementedCommentServiceServer) GetUserComments(context.Context, *GetUserCommentsRequest) (*GetUserCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserComments not implemented")
}
func (UnimplementedCommentServiceServer) GetVehicleStats(context.Context, *GetVehicleStatsRequest) (*GetVehicleStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVehicleStats not implemented")
}
func (UnimplementedCommentServiceServer) UpdateComment(context.Context, *UpdateCommentRequest) (*UpdateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedCommentServiceServer) DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedCommentServiceServer) ReplyComment(context.Context, *ReplyCommentRequest) (*ReplyCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplyComment not implemented")
}
func (UnimplementedCommentServiceServer) CheckOrderCommented(context.Context, *CheckOrderCommentedRequest) (*CheckOrderCommentedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckOrderCommented not implemented")
}
func (UnimplementedCommentServiceServer) mustEmbedUnimplementedCommentServiceServer() {}
func (UnimplementedCommentServiceServer) testEmbeddedByValue()                        {}

// UnsafeCommentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServiceServer will
// result in compilation errors.
type UnsafeCommentServiceServer interface {
	mustEmbedUnimplementedCommentServiceServer()
}

func RegisterCommentServiceServer(s grpc.ServiceRegistrar, srv CommentServiceServer) {
	// If the following call pancis, it indicates UnimplementedCommentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CommentService_ServiceDesc, srv)
}

func _CommentService_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_CreateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetComment(ctx, req.(*GetCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetOrderComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetOrderComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetOrderComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetOrderComment(ctx, req.(*GetOrderCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetVehicleComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVehicleCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetVehicleComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetVehicleComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetVehicleComments(ctx, req.(*GetVehicleCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetUserComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetUserComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetUserComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetUserComments(ctx, req.(*GetUserCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetVehicleStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVehicleStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetVehicleStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetVehicleStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetVehicleStats(ctx, req.(*GetVehicleStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_UpdateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).UpdateComment(ctx, req.(*UpdateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_DeleteComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_ReplyComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplyCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).ReplyComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_ReplyComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).ReplyComment(ctx, req.(*ReplyCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_CheckOrderCommented_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckOrderCommentedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).CheckOrderCommented(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_CheckOrderCommented_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).CheckOrderCommented(ctx, req.(*CheckOrderCommentedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentService_ServiceDesc is the grpc.ServiceDesc for CommentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comment.CommentService",
	HandlerType: (*CommentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateComment",
			Handler:    _CommentService_CreateComment_Handler,
		},
		{
			MethodName: "GetComment",
			Handler:    _CommentService_GetComment_Handler,
		},
		{
			MethodName: "GetOrderComment",
			Handler:    _CommentService_GetOrderComment_Handler,
		},
		{
			MethodName: "GetVehicleComments",
			Handler:    _CommentService_GetVehicleComments_Handler,
		},
		{
			MethodName: "GetUserComments",
			Handler:    _CommentService_GetUserComments_Handler,
		},
		{
			MethodName: "GetVehicleStats",
			Handler:    _CommentService_GetVehicleStats_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _CommentService_UpdateComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _CommentService_DeleteComment_Handler,
		},
		{
			MethodName: "ReplyComment",
			Handler:    _CommentService_ReplyComment_Handler,
		},
		{
			MethodName: "CheckOrderCommented",
			Handler:    _CommentService_CheckOrderCommented_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto_comment/comment.proto",
}
