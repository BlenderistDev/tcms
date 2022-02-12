// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tcms

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

// TcmsClient is the client API for Tcms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TcmsClient interface {
	AddAutomation(ctx context.Context, in *Automation, opts ...grpc.CallOption) (*Result, error)
	GetList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Result, error)
}

type tcmsClient struct {
	cc grpc.ClientConnInterface
}

func NewTcmsClient(cc grpc.ClientConnInterface) TcmsClient {
	return &tcmsClient{cc}
}

func (c *tcmsClient) AddAutomation(ctx context.Context, in *Automation, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/tcms.Tcms/AddAutomation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tcmsClient) GetList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/tcms.Tcms/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TcmsServer is the server API for Tcms service.
// All implementations must embed UnimplementedTcmsServer
// for forward compatibility
type TcmsServer interface {
	AddAutomation(context.Context, *Automation) (*Result, error)
	GetList(context.Context, *emptypb.Empty) (*Result, error)
	mustEmbedUnimplementedTcmsServer()
}

// UnimplementedTcmsServer must be embedded to have forward compatible implementations.
type UnimplementedTcmsServer struct {
}

func (UnimplementedTcmsServer) AddAutomation(context.Context, *Automation) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAutomation not implemented")
}
func (UnimplementedTcmsServer) GetList(context.Context, *emptypb.Empty) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedTcmsServer) mustEmbedUnimplementedTcmsServer() {}

// UnsafeTcmsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TcmsServer will
// result in compilation errors.
type UnsafeTcmsServer interface {
	mustEmbedUnimplementedTcmsServer()
}

func RegisterTcmsServer(s grpc.ServiceRegistrar, srv TcmsServer) {
	s.RegisterService(&Tcms_ServiceDesc, srv)
}

func _Tcms_AddAutomation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Automation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TcmsServer).AddAutomation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tcms.Tcms/AddAutomation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TcmsServer).AddAutomation(ctx, req.(*Automation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tcms_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TcmsServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tcms.Tcms/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TcmsServer).GetList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Tcms_ServiceDesc is the grpc.ServiceDesc for Tcms service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tcms_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tcms.Tcms",
	HandlerType: (*TcmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddAutomation",
			Handler:    _Tcms_AddAutomation_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _Tcms_GetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/tcms.proto",
}
