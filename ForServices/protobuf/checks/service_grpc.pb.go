// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.0
// source: checks/service.proto

package checks

import (
	common "protobuf/common"
	users "protobuf/users"
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
	Checks_Create_FullMethodName        = "/checks.Checks/Create"
	Checks_Remove_FullMethodName        = "/checks.Checks/Remove"
	Checks_Use_FullMethodName           = "/checks.Checks/Use"
	Checks_GetUserChecks_FullMethodName = "/checks.Checks/GetUserChecks"
)

// ChecksClient is the client API for Checks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChecksClient interface {
	// Create check
	Create(ctx context.Context, in *CheckCreate, opts ...grpc.CallOption) (*CheckFailure, error)
	// Make check used
	Remove(ctx context.Context, in *CheckId, opts ...grpc.CallOption) (*common.Response, error)
	// Use check
	Use(ctx context.Context, in *CheckUse, opts ...grpc.CallOption) (*common.Response, error)
	// Get check created by user
	GetUserChecks(ctx context.Context, in *users.Id, opts ...grpc.CallOption) (*AllChecksFailure, error)
}

type checksClient struct {
	cc grpc.ClientConnInterface
}

func NewChecksClient(cc grpc.ClientConnInterface) ChecksClient {
	return &checksClient{cc}
}

func (c *checksClient) Create(ctx context.Context, in *CheckCreate, opts ...grpc.CallOption) (*CheckFailure, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckFailure)
	err := c.cc.Invoke(ctx, Checks_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checksClient) Remove(ctx context.Context, in *CheckId, opts ...grpc.CallOption) (*common.Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Response)
	err := c.cc.Invoke(ctx, Checks_Remove_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checksClient) Use(ctx context.Context, in *CheckUse, opts ...grpc.CallOption) (*common.Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Response)
	err := c.cc.Invoke(ctx, Checks_Use_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checksClient) GetUserChecks(ctx context.Context, in *users.Id, opts ...grpc.CallOption) (*AllChecksFailure, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AllChecksFailure)
	err := c.cc.Invoke(ctx, Checks_GetUserChecks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChecksServer is the server API for Checks service.
// All implementations must embed UnimplementedChecksServer
// for forward compatibility.
type ChecksServer interface {
	// Create check
	Create(context.Context, *CheckCreate) (*CheckFailure, error)
	// Make check used
	Remove(context.Context, *CheckId) (*common.Response, error)
	// Use check
	Use(context.Context, *CheckUse) (*common.Response, error)
	// Get check created by user
	GetUserChecks(context.Context, *users.Id) (*AllChecksFailure, error)
	mustEmbedUnimplementedChecksServer()
}

// UnimplementedChecksServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChecksServer struct{}

func (UnimplementedChecksServer) Create(context.Context, *CheckCreate) (*CheckFailure, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedChecksServer) Remove(context.Context, *CheckId) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedChecksServer) Use(context.Context, *CheckUse) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Use not implemented")
}
func (UnimplementedChecksServer) GetUserChecks(context.Context, *users.Id) (*AllChecksFailure, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserChecks not implemented")
}
func (UnimplementedChecksServer) mustEmbedUnimplementedChecksServer() {}
func (UnimplementedChecksServer) testEmbeddedByValue()                {}

// UnsafeChecksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChecksServer will
// result in compilation errors.
type UnsafeChecksServer interface {
	mustEmbedUnimplementedChecksServer()
}

func RegisterChecksServer(s grpc.ServiceRegistrar, srv ChecksServer) {
	// If the following call pancis, it indicates UnimplementedChecksServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Checks_ServiceDesc, srv)
}

func _Checks_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChecksServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Checks_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChecksServer).Create(ctx, req.(*CheckCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checks_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChecksServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Checks_Remove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChecksServer).Remove(ctx, req.(*CheckId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checks_Use_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChecksServer).Use(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Checks_Use_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChecksServer).Use(ctx, req.(*CheckUse))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checks_GetUserChecks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(users.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChecksServer).GetUserChecks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Checks_GetUserChecks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChecksServer).GetUserChecks(ctx, req.(*users.Id))
	}
	return interceptor(ctx, in, info, handler)
}

// Checks_ServiceDesc is the grpc.ServiceDesc for Checks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Checks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "checks.Checks",
	HandlerType: (*ChecksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Checks_Create_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Checks_Remove_Handler,
		},
		{
			MethodName: "Use",
			Handler:    _Checks_Use_Handler,
		},
		{
			MethodName: "GetUserChecks",
			Handler:    _Checks_GetUserChecks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "checks/service.proto",
}
