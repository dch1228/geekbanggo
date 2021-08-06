// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoServiceClient interface {
	CreateTodo(ctx context.Context, in *CreateTodoReq, opts ...grpc.CallOption) (*CreateTodoReply, error)
	UpdateTodo(ctx context.Context, in *UpdateTodoReq, opts ...grpc.CallOption) (*UpdateTodoReply, error)
	ListTodo(ctx context.Context, in *ListTodoReq, opts ...grpc.CallOption) (*ListTodoReply, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) CreateTodo(ctx context.Context, in *CreateTodoReq, opts ...grpc.CallOption) (*CreateTodoReply, error) {
	out := new(CreateTodoReply)
	err := c.cc.Invoke(ctx, "/api.todo.v1.TodoService/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) UpdateTodo(ctx context.Context, in *UpdateTodoReq, opts ...grpc.CallOption) (*UpdateTodoReply, error) {
	out := new(UpdateTodoReply)
	err := c.cc.Invoke(ctx, "/api.todo.v1.TodoService/UpdateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) ListTodo(ctx context.Context, in *ListTodoReq, opts ...grpc.CallOption) (*ListTodoReply, error) {
	out := new(ListTodoReply)
	err := c.cc.Invoke(ctx, "/api.todo.v1.TodoService/ListTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
// All implementations must embed UnimplementedTodoServiceServer
// for forward compatibility
type TodoServiceServer interface {
	CreateTodo(context.Context, *CreateTodoReq) (*CreateTodoReply, error)
	UpdateTodo(context.Context, *UpdateTodoReq) (*UpdateTodoReply, error)
	ListTodo(context.Context, *ListTodoReq) (*ListTodoReply, error)
	mustEmbedUnimplementedTodoServiceServer()
}

// UnimplementedTodoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (UnimplementedTodoServiceServer) CreateTodo(context.Context, *CreateTodoReq) (*CreateTodoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (UnimplementedTodoServiceServer) UpdateTodo(context.Context, *UpdateTodoReq) (*UpdateTodoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}
func (UnimplementedTodoServiceServer) ListTodo(context.Context, *ListTodoReq) (*ListTodoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTodo not implemented")
}
func (UnimplementedTodoServiceServer) mustEmbedUnimplementedTodoServiceServer() {}

// UnsafeTodoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServiceServer will
// result in compilation errors.
type UnsafeTodoServiceServer interface {
	mustEmbedUnimplementedTodoServiceServer()
}

func RegisterTodoServiceServer(s grpc.ServiceRegistrar, srv TodoServiceServer) {
	s.RegisterService(&TodoService_ServiceDesc, srv)
}

func _TodoService_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.todo.v1.TodoService/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateTodo(ctx, req.(*CreateTodoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_UpdateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTodoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).UpdateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.todo.v1.TodoService/UpdateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).UpdateTodo(ctx, req.(*UpdateTodoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_ListTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTodoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).ListTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.todo.v1.TodoService/ListTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).ListTodo(ctx, req.(*ListTodoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoService_ServiceDesc is the grpc.ServiceDesc for TodoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.todo.v1.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _TodoService_CreateTodo_Handler,
		},
		{
			MethodName: "UpdateTodo",
			Handler:    _TodoService_UpdateTodo_Handler,
		},
		{
			MethodName: "ListTodo",
			Handler:    _TodoService_ListTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/todo/v1/todo.proto",
}
