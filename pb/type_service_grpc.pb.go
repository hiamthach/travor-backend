// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: type_service.proto

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

const (
	TypeService_GetTypes_FullMethodName   = "/pb.TypeService/GetTypes"
	TypeService_CreateType_FullMethodName = "/pb.TypeService/CreateType"
	TypeService_DeleteType_FullMethodName = "/pb.TypeService/DeleteType"
)

// TypeServiceClient is the client API for TypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TypeServiceClient interface {
	GetTypes(ctx context.Context, in *GetTypesRequest, opts ...grpc.CallOption) (*GetTypesResponse, error)
	CreateType(ctx context.Context, in *CreateTypeRequest, opts ...grpc.CallOption) (*CreateTypeResponse, error)
	DeleteType(ctx context.Context, in *DeleteTypeRequest, opts ...grpc.CallOption) (*DeleteTypeResponse, error)
}

type typeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTypeServiceClient(cc grpc.ClientConnInterface) TypeServiceClient {
	return &typeServiceClient{cc}
}

func (c *typeServiceClient) GetTypes(ctx context.Context, in *GetTypesRequest, opts ...grpc.CallOption) (*GetTypesResponse, error) {
	out := new(GetTypesResponse)
	err := c.cc.Invoke(ctx, TypeService_GetTypes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typeServiceClient) CreateType(ctx context.Context, in *CreateTypeRequest, opts ...grpc.CallOption) (*CreateTypeResponse, error) {
	out := new(CreateTypeResponse)
	err := c.cc.Invoke(ctx, TypeService_CreateType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typeServiceClient) DeleteType(ctx context.Context, in *DeleteTypeRequest, opts ...grpc.CallOption) (*DeleteTypeResponse, error) {
	out := new(DeleteTypeResponse)
	err := c.cc.Invoke(ctx, TypeService_DeleteType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TypeServiceServer is the server API for TypeService service.
// All implementations must embed UnimplementedTypeServiceServer
// for forward compatibility
type TypeServiceServer interface {
	GetTypes(context.Context, *GetTypesRequest) (*GetTypesResponse, error)
	CreateType(context.Context, *CreateTypeRequest) (*CreateTypeResponse, error)
	DeleteType(context.Context, *DeleteTypeRequest) (*DeleteTypeResponse, error)
	mustEmbedUnimplementedTypeServiceServer()
}

// UnimplementedTypeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTypeServiceServer struct {
}

func (UnimplementedTypeServiceServer) GetTypes(context.Context, *GetTypesRequest) (*GetTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTypes not implemented")
}
func (UnimplementedTypeServiceServer) CreateType(context.Context, *CreateTypeRequest) (*CreateTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateType not implemented")
}
func (UnimplementedTypeServiceServer) DeleteType(context.Context, *DeleteTypeRequest) (*DeleteTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteType not implemented")
}
func (UnimplementedTypeServiceServer) mustEmbedUnimplementedTypeServiceServer() {}

// UnsafeTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TypeServiceServer will
// result in compilation errors.
type UnsafeTypeServiceServer interface {
	mustEmbedUnimplementedTypeServiceServer()
}

func RegisterTypeServiceServer(s grpc.ServiceRegistrar, srv TypeServiceServer) {
	s.RegisterService(&TypeService_ServiceDesc, srv)
}

func _TypeService_GetTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypeServiceServer).GetTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TypeService_GetTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypeServiceServer).GetTypes(ctx, req.(*GetTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TypeService_CreateType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypeServiceServer).CreateType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TypeService_CreateType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypeServiceServer).CreateType(ctx, req.(*CreateTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TypeService_DeleteType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypeServiceServer).DeleteType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TypeService_DeleteType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypeServiceServer).DeleteType(ctx, req.(*DeleteTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TypeService_ServiceDesc is the grpc.ServiceDesc for TypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TypeService",
	HandlerType: (*TypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTypes",
			Handler:    _TypeService_GetTypes_Handler,
		},
		{
			MethodName: "CreateType",
			Handler:    _TypeService_CreateType_Handler,
		},
		{
			MethodName: "DeleteType",
			Handler:    _TypeService_DeleteType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "type_service.proto",
}
