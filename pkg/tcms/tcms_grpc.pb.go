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
	UpdateAutomation(ctx context.Context, in *UpdateAutomationRequest, opts ...grpc.CallOption) (*Result, error)
	GetList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AutomationList, error)
	GetOne(ctx context.Context, in *AutomationRequest, opts ...grpc.CallOption) (*Automation, error)
	RemoveAutomation(ctx context.Context, in *RemoveAutomationRequest, opts ...grpc.CallOption) (*Result, error)
	GetConditionList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ConditionList, error)
	GetActionList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ActionList, error)
	GetTriggerList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TriggerList, error)
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

func (c *tcmsClient) UpdateAutomation(ctx context.Context, in *UpdateAutomationRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/tcms.Tcms/UpdateAutomation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tcmsClient) GetList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AutomationList, error) {
	out := new(AutomationList)
	err := c.cc.Invoke(ctx, "/tcms.Tcms/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tcmsClient) GetOne(ctx context.Context, in *AutomationRequest, opts ...grpc.CallOption) (*Automation, error) {
	out := new(Automation)
	err := c.cc.Invoke(ctx, "/tcms.Tcms/GetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tcmsClient) RemoveAutomation(ctx context.Context, in *RemoveAutomationRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/tcms.Tcms/RemoveAutomation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tcmsClient) GetConditionList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ConditionList, error) {
	out := new(ConditionList)
	err := c.cc.Invoke(ctx, "/tcms.Tcms/GetConditionList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tcmsClient) GetActionList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ActionList, error) {
	out := new(ActionList)
	err := c.cc.Invoke(ctx, "/tcms.Tcms/GetActionList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tcmsClient) GetTriggerList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TriggerList, error) {
	out := new(TriggerList)
	err := c.cc.Invoke(ctx, "/tcms.Tcms/GetTriggerList", in, out, opts...)
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
	UpdateAutomation(context.Context, *UpdateAutomationRequest) (*Result, error)
	GetList(context.Context, *emptypb.Empty) (*AutomationList, error)
	GetOne(context.Context, *AutomationRequest) (*Automation, error)
	RemoveAutomation(context.Context, *RemoveAutomationRequest) (*Result, error)
	GetConditionList(context.Context, *emptypb.Empty) (*ConditionList, error)
	GetActionList(context.Context, *emptypb.Empty) (*ActionList, error)
	GetTriggerList(context.Context, *emptypb.Empty) (*TriggerList, error)
	mustEmbedUnimplementedTcmsServer()
}

// UnimplementedTcmsServer must be embedded to have forward compatible implementations.
type UnimplementedTcmsServer struct {
}

func (UnimplementedTcmsServer) AddAutomation(context.Context, *Automation) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAutomation not implemented")
}
func (UnimplementedTcmsServer) UpdateAutomation(context.Context, *UpdateAutomationRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAutomation not implemented")
}
func (UnimplementedTcmsServer) GetList(context.Context, *emptypb.Empty) (*AutomationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedTcmsServer) GetOne(context.Context, *AutomationRequest) (*Automation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedTcmsServer) RemoveAutomation(context.Context, *RemoveAutomationRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAutomation not implemented")
}
func (UnimplementedTcmsServer) GetConditionList(context.Context, *emptypb.Empty) (*ConditionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConditionList not implemented")
}
func (UnimplementedTcmsServer) GetActionList(context.Context, *emptypb.Empty) (*ActionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActionList not implemented")
}
func (UnimplementedTcmsServer) GetTriggerList(context.Context, *emptypb.Empty) (*TriggerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTriggerList not implemented")
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

func _Tcms_UpdateAutomation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAutomationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TcmsServer).UpdateAutomation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tcms.Tcms/UpdateAutomation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TcmsServer).UpdateAutomation(ctx, req.(*UpdateAutomationRequest))
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

func _Tcms_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AutomationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TcmsServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tcms.Tcms/GetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TcmsServer).GetOne(ctx, req.(*AutomationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tcms_RemoveAutomation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveAutomationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TcmsServer).RemoveAutomation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tcms.Tcms/RemoveAutomation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TcmsServer).RemoveAutomation(ctx, req.(*RemoveAutomationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tcms_GetConditionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TcmsServer).GetConditionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tcms.Tcms/GetConditionList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TcmsServer).GetConditionList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tcms_GetActionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TcmsServer).GetActionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tcms.Tcms/GetActionList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TcmsServer).GetActionList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tcms_GetTriggerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TcmsServer).GetTriggerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tcms.Tcms/GetTriggerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TcmsServer).GetTriggerList(ctx, req.(*emptypb.Empty))
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
			MethodName: "UpdateAutomation",
			Handler:    _Tcms_UpdateAutomation_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _Tcms_GetList_Handler,
		},
		{
			MethodName: "GetOne",
			Handler:    _Tcms_GetOne_Handler,
		},
		{
			MethodName: "RemoveAutomation",
			Handler:    _Tcms_RemoveAutomation_Handler,
		},
		{
			MethodName: "GetConditionList",
			Handler:    _Tcms_GetConditionList_Handler,
		},
		{
			MethodName: "GetActionList",
			Handler:    _Tcms_GetActionList_Handler,
		},
		{
			MethodName: "GetTriggerList",
			Handler:    _Tcms_GetTriggerList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/tcms.proto",
}
