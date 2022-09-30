// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: ibc/applications/ibcswap/v1/tx.proto

package types

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

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	DelegateCreatePool(ctx context.Context, in *MsgCreatePoolRequest, opts ...grpc.CallOption) (*MsgCreatePoolResponse, error)
	DelegateSingleDeposit(ctx context.Context, in *MsgSingleDepositRequest, opts ...grpc.CallOption) (*MsgSingleDepositResponse, error)
	DelegateWithdraw(ctx context.Context, in *MsgWithdrawRequest, opts ...grpc.CallOption) (*MsgWithdrawResponse, error)
	DelegateLeftSwap(ctx context.Context, in *MsgLeftSwapRequest, opts ...grpc.CallOption) (*MsgSwapResponse, error)
	DelegateRightSwap(ctx context.Context, in *MsgRightSwapRequest, opts ...grpc.CallOption) (*MsgSwapResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) DelegateCreatePool(ctx context.Context, in *MsgCreatePoolRequest, opts ...grpc.CallOption) (*MsgCreatePoolResponse, error) {
	out := new(MsgCreatePoolResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.ibcswap.v1.Msg/DelegateCreatePool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DelegateSingleDeposit(ctx context.Context, in *MsgSingleDepositRequest, opts ...grpc.CallOption) (*MsgSingleDepositResponse, error) {
	out := new(MsgSingleDepositResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.ibcswap.v1.Msg/DelegateSingleDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DelegateWithdraw(ctx context.Context, in *MsgWithdrawRequest, opts ...grpc.CallOption) (*MsgWithdrawResponse, error) {
	out := new(MsgWithdrawResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.ibcswap.v1.Msg/DelegateWithdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DelegateLeftSwap(ctx context.Context, in *MsgLeftSwapRequest, opts ...grpc.CallOption) (*MsgSwapResponse, error) {
	out := new(MsgSwapResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.ibcswap.v1.Msg/DelegateLeftSwap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DelegateRightSwap(ctx context.Context, in *MsgRightSwapRequest, opts ...grpc.CallOption) (*MsgSwapResponse, error) {
	out := new(MsgSwapResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.ibcswap.v1.Msg/DelegateRightSwap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations should embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	DelegateCreatePool(context.Context, *MsgCreatePoolRequest) (*MsgCreatePoolResponse, error)
	DelegateSingleDeposit(context.Context, *MsgSingleDepositRequest) (*MsgSingleDepositResponse, error)
	DelegateWithdraw(context.Context, *MsgWithdrawRequest) (*MsgWithdrawResponse, error)
	DelegateLeftSwap(context.Context, *MsgLeftSwapRequest) (*MsgSwapResponse, error)
	DelegateRightSwap(context.Context, *MsgRightSwapRequest) (*MsgSwapResponse, error)
}

// UnimplementedMsgServer should be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) DelegateCreatePool(context.Context, *MsgCreatePoolRequest) (*MsgCreatePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegateCreatePool not implemented")
}
func (UnimplementedMsgServer) DelegateSingleDeposit(context.Context, *MsgSingleDepositRequest) (*MsgSingleDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegateSingleDeposit not implemented")
}
func (UnimplementedMsgServer) DelegateWithdraw(context.Context, *MsgWithdrawRequest) (*MsgWithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegateWithdraw not implemented")
}
func (UnimplementedMsgServer) DelegateLeftSwap(context.Context, *MsgLeftSwapRequest) (*MsgSwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegateLeftSwap not implemented")
}
func (UnimplementedMsgServer) DelegateRightSwap(context.Context, *MsgRightSwapRequest) (*MsgSwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegateRightSwap not implemented")
}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_DelegateCreatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DelegateCreatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.ibcswap.v1.Msg/DelegateCreatePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DelegateCreatePool(ctx, req.(*MsgCreatePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DelegateSingleDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSingleDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DelegateSingleDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.ibcswap.v1.Msg/DelegateSingleDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DelegateSingleDeposit(ctx, req.(*MsgSingleDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DelegateWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DelegateWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.ibcswap.v1.Msg/DelegateWithdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DelegateWithdraw(ctx, req.(*MsgWithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DelegateLeftSwap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLeftSwapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DelegateLeftSwap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.ibcswap.v1.Msg/DelegateLeftSwap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DelegateLeftSwap(ctx, req.(*MsgLeftSwapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DelegateRightSwap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRightSwapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DelegateRightSwap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.ibcswap.v1.Msg/DelegateRightSwap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DelegateRightSwap(ctx, req.(*MsgRightSwapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.applications.ibcswap.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DelegateCreatePool",
			Handler:    _Msg_DelegateCreatePool_Handler,
		},
		{
			MethodName: "DelegateSingleDeposit",
			Handler:    _Msg_DelegateSingleDeposit_Handler,
		},
		{
			MethodName: "DelegateWithdraw",
			Handler:    _Msg_DelegateWithdraw_Handler,
		},
		{
			MethodName: "DelegateLeftSwap",
			Handler:    _Msg_DelegateLeftSwap_Handler,
		},
		{
			MethodName: "DelegateRightSwap",
			Handler:    _Msg_DelegateRightSwap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/applications/ibcswap/v1/tx.proto",
}
