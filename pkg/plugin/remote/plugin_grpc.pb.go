// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package remote

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

// PluginInitializerClient is the client API for PluginInitializer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginInitializerClient interface {
	Register(ctx context.Context, in *PluginInfo, opts ...grpc.CallOption) (*Context, error)
	IsLeader(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*LeaderInfo, error)
}

type pluginInitializerClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginInitializerClient(cc grpc.ClientConnInterface) PluginInitializerClient {
	return &pluginInitializerClient{cc}
}

func (c *pluginInitializerClient) Register(ctx context.Context, in *PluginInfo, opts ...grpc.CallOption) (*Context, error) {
	out := new(Context)
	err := c.cc.Invoke(ctx, "/remote.PluginInitializer/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginInitializerClient) IsLeader(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*LeaderInfo, error) {
	out := new(LeaderInfo)
	err := c.cc.Invoke(ctx, "/remote.PluginInitializer/IsLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginInitializerServer is the server API for PluginInitializer service.
// All implementations must embed UnimplementedPluginInitializerServer
// for forward compatibility
type PluginInitializerServer interface {
	Register(context.Context, *PluginInfo) (*Context, error)
	IsLeader(context.Context, *Empty) (*LeaderInfo, error)
	mustEmbedUnimplementedPluginInitializerServer()
}

// UnimplementedPluginInitializerServer must be embedded to have forward compatible implementations.
type UnimplementedPluginInitializerServer struct {
}

func (UnimplementedPluginInitializerServer) Register(context.Context, *PluginInfo) (*Context, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedPluginInitializerServer) IsLeader(context.Context, *Empty) (*LeaderInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsLeader not implemented")
}
func (UnimplementedPluginInitializerServer) mustEmbedUnimplementedPluginInitializerServer() {}

// UnsafePluginInitializerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginInitializerServer will
// result in compilation errors.
type UnsafePluginInitializerServer interface {
	mustEmbedUnimplementedPluginInitializerServer()
}

func RegisterPluginInitializerServer(s grpc.ServiceRegistrar, srv PluginInitializerServer) {
	s.RegisterService(&PluginInitializer_ServiceDesc, srv)
}

func _PluginInitializer_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginInitializerServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.PluginInitializer/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginInitializerServer).Register(ctx, req.(*PluginInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginInitializer_IsLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginInitializerServer).IsLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.PluginInitializer/IsLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginInitializerServer).IsLeader(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// PluginInitializer_ServiceDesc is the grpc.ServiceDesc for PluginInitializer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PluginInitializer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "remote.PluginInitializer",
	HandlerType: (*PluginInitializerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _PluginInitializer_Register_Handler,
		},
		{
			MethodName: "IsLeader",
			Handler:    _PluginInitializer_IsLeader_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin.proto",
}