// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: followaction.proto

package pb

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

// FollowActionClient is the client API for FollowAction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FollowActionClient interface {
	FollowAction(ctx context.Context, in *DouyinRelationActionRequest, opts ...grpc.CallOption) (*DouyinRelationActionResponse, error)
}

type followActionClient struct {
	cc grpc.ClientConnInterface
}

func NewFollowActionClient(cc grpc.ClientConnInterface) FollowActionClient {
	return &followActionClient{cc}
}

func (c *followActionClient) FollowAction(ctx context.Context, in *DouyinRelationActionRequest, opts ...grpc.CallOption) (*DouyinRelationActionResponse, error) {
	out := new(DouyinRelationActionResponse)
	err := c.cc.Invoke(ctx, "/douyin.extra.second.FollowAction/FollowAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FollowActionServer is the server API for FollowAction service.
// All implementations must embed UnimplementedFollowActionServer
// for forward compatibility
type FollowActionServer interface {
	FollowAction(context.Context, *DouyinRelationActionRequest) (*DouyinRelationActionResponse, error)
	mustEmbedUnimplementedFollowActionServer()
}

// UnimplementedFollowActionServer must be embedded to have forward compatible implementations.
type UnimplementedFollowActionServer struct {
}

func (UnimplementedFollowActionServer) FollowAction(context.Context, *DouyinRelationActionRequest) (*DouyinRelationActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowAction not implemented")
}
func (UnimplementedFollowActionServer) mustEmbedUnimplementedFollowActionServer() {}

// UnsafeFollowActionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FollowActionServer will
// result in compilation errors.
type UnsafeFollowActionServer interface {
	mustEmbedUnimplementedFollowActionServer()
}

func RegisterFollowActionServer(s grpc.ServiceRegistrar, srv FollowActionServer) {
	s.RegisterService(&FollowAction_ServiceDesc, srv)
}

func _FollowAction_FollowAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinRelationActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowActionServer).FollowAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/douyin.extra.second.FollowAction/FollowAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowActionServer).FollowAction(ctx, req.(*DouyinRelationActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FollowAction_ServiceDesc is the grpc.ServiceDesc for FollowAction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FollowAction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "douyin.extra.second.FollowAction",
	HandlerType: (*FollowActionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FollowAction",
			Handler:    _FollowAction_FollowAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "followaction.proto",
}
