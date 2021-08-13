// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ruby_vagrant

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RubyVagrantClient is the client API for RubyVagrant service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RubyVagrantClient interface {
	// Gets available ruby plugins
	GetPlugins(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetPluginsResponse, error)
	ParseVagrantfile(ctx context.Context, in *ParseVagrantfileRequest, opts ...grpc.CallOption) (*ParseVagrantfileResponse, error)
}

type rubyVagrantClient struct {
	cc grpc.ClientConnInterface
}

func NewRubyVagrantClient(cc grpc.ClientConnInterface) RubyVagrantClient {
	return &rubyVagrantClient{cc}
}

func (c *rubyVagrantClient) GetPlugins(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetPluginsResponse, error) {
	out := new(GetPluginsResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.vagrant.RubyVagrant/GetPlugins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rubyVagrantClient) ParseVagrantfile(ctx context.Context, in *ParseVagrantfileRequest, opts ...grpc.CallOption) (*ParseVagrantfileResponse, error) {
	out := new(ParseVagrantfileResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.vagrant.RubyVagrant/ParseVagrantfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RubyVagrantServer is the server API for RubyVagrant service.
// All implementations must embed UnimplementedRubyVagrantServer
// for forward compatibility
type RubyVagrantServer interface {
	// Gets available ruby plugins
	GetPlugins(context.Context, *emptypb.Empty) (*GetPluginsResponse, error)
	ParseVagrantfile(context.Context, *ParseVagrantfileRequest) (*ParseVagrantfileResponse, error)
	mustEmbedUnimplementedRubyVagrantServer()
}

// UnimplementedRubyVagrantServer must be embedded to have forward compatible implementations.
type UnimplementedRubyVagrantServer struct {
}

func (UnimplementedRubyVagrantServer) GetPlugins(context.Context, *emptypb.Empty) (*GetPluginsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlugins not implemented")
}
func (UnimplementedRubyVagrantServer) ParseVagrantfile(context.Context, *ParseVagrantfileRequest) (*ParseVagrantfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseVagrantfile not implemented")
}
func (UnimplementedRubyVagrantServer) mustEmbedUnimplementedRubyVagrantServer() {}

// UnsafeRubyVagrantServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RubyVagrantServer will
// result in compilation errors.
type UnsafeRubyVagrantServer interface {
	mustEmbedUnimplementedRubyVagrantServer()
}

func RegisterRubyVagrantServer(s grpc.ServiceRegistrar, srv RubyVagrantServer) {
	s.RegisterService(&RubyVagrant_ServiceDesc, srv)
}

func _RubyVagrant_GetPlugins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RubyVagrantServer).GetPlugins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.vagrant.RubyVagrant/GetPlugins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RubyVagrantServer).GetPlugins(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RubyVagrant_ParseVagrantfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseVagrantfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RubyVagrantServer).ParseVagrantfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.vagrant.RubyVagrant/ParseVagrantfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RubyVagrantServer).ParseVagrantfile(ctx, req.(*ParseVagrantfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RubyVagrant_ServiceDesc is the grpc.ServiceDesc for RubyVagrant service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RubyVagrant_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.vagrant.RubyVagrant",
	HandlerType: (*RubyVagrantServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlugins",
			Handler:    _RubyVagrant_GetPlugins_Handler,
		},
		{
			MethodName: "ParseVagrantfile",
			Handler:    _RubyVagrant_ParseVagrantfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ruby_vagrant/ruby-server.proto",
}
