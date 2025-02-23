// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: Withdraw/withdraw.proto

package withdrawalv1

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
	Withdrawal_Withdraw_FullMethodName = "/withdrawal.Withdrawal/Withdraw"
)

// WithdrawalClient is the client API for Withdrawal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WithdrawalClient interface {
	Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawResponse, error)
}

type withdrawalClient struct {
	cc grpc.ClientConnInterface
}

func NewWithdrawalClient(cc grpc.ClientConnInterface) WithdrawalClient {
	return &withdrawalClient{cc}
}

func (c *withdrawalClient) Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WithdrawResponse)
	err := c.cc.Invoke(ctx, Withdrawal_Withdraw_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WithdrawalServer is the server API for Withdrawal service.
// All implementations must embed UnimplementedWithdrawalServer
// for forward compatibility.
type WithdrawalServer interface {
	Withdraw(context.Context, *WithdrawRequest) (*WithdrawResponse, error)
	mustEmbedUnimplementedWithdrawalServer()
}

// UnimplementedWithdrawalServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWithdrawalServer struct{}

func (UnimplementedWithdrawalServer) Withdraw(context.Context, *WithdrawRequest) (*WithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}
func (UnimplementedWithdrawalServer) mustEmbedUnimplementedWithdrawalServer() {}
func (UnimplementedWithdrawalServer) testEmbeddedByValue()                    {}

// UnsafeWithdrawalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WithdrawalServer will
// result in compilation errors.
type UnsafeWithdrawalServer interface {
	mustEmbedUnimplementedWithdrawalServer()
}

func RegisterWithdrawalServer(s grpc.ServiceRegistrar, srv WithdrawalServer) {
	// If the following call pancis, it indicates UnimplementedWithdrawalServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Withdrawal_ServiceDesc, srv)
}

func _Withdrawal_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WithdrawalServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Withdrawal_Withdraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WithdrawalServer).Withdraw(ctx, req.(*WithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Withdrawal_ServiceDesc is the grpc.ServiceDesc for Withdrawal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Withdrawal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "withdrawal.Withdrawal",
	HandlerType: (*WithdrawalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Withdraw",
			Handler:    _Withdrawal_Withdraw_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Withdraw/withdraw.proto",
}
