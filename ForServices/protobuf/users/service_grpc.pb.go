// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.0
// source: users/service.proto

package users

import (
	common "protobuf/common"
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
	Users_SendTransaction_FullMethodName       = "/users.Users/SendTransaction"
	Users_ChangeAutoBuy_FullMethodName         = "/users.Users/ChangeAutoBuy"
	Users_Create_FullMethodName                = "/users.Users/Create"
	Users_GetUser_FullMethodName               = "/users.Users/GetUser"
	Users_GetTop_FullMethodName                = "/users.Users/GetTop"
	Users_GetAll_FullMethodName                = "/users.Users/GetAll"
	Users_GetTransaction_FullMethodName        = "/users.Users/GetTransaction"
	Users_GetTransactionHistory_FullMethodName = "/users.Users/GetTransactionHistory"
	Users_GetAllTransactions_FullMethodName    = "/users.Users/GetAllTransactions"
	Users_Farm_FullMethodName                  = "/users.Users/Farm"
)

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	SendTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
	ChangeAutoBuy(ctx context.Context, in *Id, opts ...grpc.CallOption) (*common.Response, error)
	Create(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*User, error)
	GetUser(ctx context.Context, in *Id, opts ...grpc.CallOption) (*User, error)
	GetTop(ctx context.Context, in *GetTopUsers, opts ...grpc.CallOption) (*RepeatedUsers, error)
	GetAll(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*RepeatedUsers, error)
	GetTransaction(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Transaction, error)
	GetTransactionHistory(ctx context.Context, in *Id, opts ...grpc.CallOption) (*TransactionHistory, error)
	GetAllTransactions(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*TransactionHistory, error)
	Farm(ctx context.Context, in *Id, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) SendTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, Users_SendTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ChangeAutoBuy(ctx context.Context, in *Id, opts ...grpc.CallOption) (*common.Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Response)
	err := c.cc.Invoke(ctx, Users_ChangeAutoBuy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Create(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, Users_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUser(ctx context.Context, in *Id, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, Users_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetTop(ctx context.Context, in *GetTopUsers, opts ...grpc.CallOption) (*RepeatedUsers, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RepeatedUsers)
	err := c.cc.Invoke(ctx, Users_GetTop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetAll(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*RepeatedUsers, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RepeatedUsers)
	err := c.cc.Invoke(ctx, Users_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetTransaction(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Transaction, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Transaction)
	err := c.cc.Invoke(ctx, Users_GetTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetTransactionHistory(ctx context.Context, in *Id, opts ...grpc.CallOption) (*TransactionHistory, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionHistory)
	err := c.cc.Invoke(ctx, Users_GetTransactionHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetAllTransactions(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*TransactionHistory, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionHistory)
	err := c.cc.Invoke(ctx, Users_GetAllTransactions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Farm(ctx context.Context, in *Id, opts ...grpc.CallOption) (*TransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, Users_Farm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
// All implementations must embed UnimplementedUsersServer
// for forward compatibility.
type UsersServer interface {
	SendTransaction(context.Context, *TransactionRequest) (*TransactionResponse, error)
	ChangeAutoBuy(context.Context, *Id) (*common.Response, error)
	Create(context.Context, *common.Void) (*User, error)
	GetUser(context.Context, *Id) (*User, error)
	GetTop(context.Context, *GetTopUsers) (*RepeatedUsers, error)
	GetAll(context.Context, *common.Void) (*RepeatedUsers, error)
	GetTransaction(context.Context, *Id) (*Transaction, error)
	GetTransactionHistory(context.Context, *Id) (*TransactionHistory, error)
	GetAllTransactions(context.Context, *common.Void) (*TransactionHistory, error)
	Farm(context.Context, *Id) (*TransactionResponse, error)
	mustEmbedUnimplementedUsersServer()
}

// UnimplementedUsersServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUsersServer struct{}

func (UnimplementedUsersServer) SendTransaction(context.Context, *TransactionRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTransaction not implemented")
}
func (UnimplementedUsersServer) ChangeAutoBuy(context.Context, *Id) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeAutoBuy not implemented")
}
func (UnimplementedUsersServer) Create(context.Context, *common.Void) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUsersServer) GetUser(context.Context, *Id) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUsersServer) GetTop(context.Context, *GetTopUsers) (*RepeatedUsers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTop not implemented")
}
func (UnimplementedUsersServer) GetAll(context.Context, *common.Void) (*RepeatedUsers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedUsersServer) GetTransaction(context.Context, *Id) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedUsersServer) GetTransactionHistory(context.Context, *Id) (*TransactionHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionHistory not implemented")
}
func (UnimplementedUsersServer) GetAllTransactions(context.Context, *common.Void) (*TransactionHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTransactions not implemented")
}
func (UnimplementedUsersServer) Farm(context.Context, *Id) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Farm not implemented")
}
func (UnimplementedUsersServer) mustEmbedUnimplementedUsersServer() {}
func (UnimplementedUsersServer) testEmbeddedByValue()               {}

// UnsafeUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServer will
// result in compilation errors.
type UnsafeUsersServer interface {
	mustEmbedUnimplementedUsersServer()
}

func RegisterUsersServer(s grpc.ServiceRegistrar, srv UsersServer) {
	// If the following call pancis, it indicates UnimplementedUsersServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Users_ServiceDesc, srv)
}

func _Users_SendTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).SendTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_SendTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).SendTransaction(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ChangeAutoBuy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ChangeAutoBuy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_ChangeAutoBuy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ChangeAutoBuy(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Create(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUser(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetTop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopUsers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetTop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_GetTop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetTop(ctx, req.(*GetTopUsers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetAll(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_GetTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetTransaction(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetTransactionHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetTransactionHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_GetTransactionHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetTransactionHistory(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetAllTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetAllTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_GetAllTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetAllTransactions(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Farm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Farm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_Farm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Farm(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// Users_ServiceDesc is the grpc.ServiceDesc for Users service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Users_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendTransaction",
			Handler:    _Users_SendTransaction_Handler,
		},
		{
			MethodName: "ChangeAutoBuy",
			Handler:    _Users_ChangeAutoBuy_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Users_Create_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Users_GetUser_Handler,
		},
		{
			MethodName: "GetTop",
			Handler:    _Users_GetTop_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _Users_GetAll_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _Users_GetTransaction_Handler,
		},
		{
			MethodName: "GetTransactionHistory",
			Handler:    _Users_GetTransactionHistory_Handler,
		},
		{
			MethodName: "GetAllTransactions",
			Handler:    _Users_GetAllTransactions_Handler,
		},
		{
			MethodName: "Farm",
			Handler:    _Users_Farm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users/service.proto",
}
