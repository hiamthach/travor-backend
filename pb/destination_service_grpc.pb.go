// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: destination_service.proto

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
	DestinationService_IsValid_FullMethodName            = "/pb.DestinationService/isValid"
	DestinationService_GetDestinations_FullMethodName    = "/pb.DestinationService/GetDestinations"
	DestinationService_GetDestinationById_FullMethodName = "/pb.DestinationService/GetDestinationById"
	DestinationService_CreateDestination_FullMethodName  = "/pb.DestinationService/CreateDestination"
	DestinationService_UpdateDestination_FullMethodName  = "/pb.DestinationService/UpdateDestination"
	DestinationService_DeleteDestination_FullMethodName  = "/pb.DestinationService/DeleteDestination"
)

// DestinationServiceClient is the client API for DestinationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DestinationServiceClient interface {
	IsValid(ctx context.Context, in *ValidRequest, opts ...grpc.CallOption) (*DestinationValid, error)
	GetDestinations(ctx context.Context, in *GetDestinationsRequest, opts ...grpc.CallOption) (*GetDestinationsResponse, error)
	GetDestinationById(ctx context.Context, in *GetDestinationByIdRequest, opts ...grpc.CallOption) (*Destination, error)
	CreateDestination(ctx context.Context, in *CreateDestinationRequest, opts ...grpc.CallOption) (*CreateDestinationResponse, error)
	UpdateDestination(ctx context.Context, in *UpdateDestinationRequest, opts ...grpc.CallOption) (*UpdateDestinationResponse, error)
	DeleteDestination(ctx context.Context, in *DeleteDestinationRequest, opts ...grpc.CallOption) (*DeleteDestinationResponse, error)
}

type destinationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDestinationServiceClient(cc grpc.ClientConnInterface) DestinationServiceClient {
	return &destinationServiceClient{cc}
}

func (c *destinationServiceClient) IsValid(ctx context.Context, in *ValidRequest, opts ...grpc.CallOption) (*DestinationValid, error) {
	out := new(DestinationValid)
	err := c.cc.Invoke(ctx, DestinationService_IsValid_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationServiceClient) GetDestinations(ctx context.Context, in *GetDestinationsRequest, opts ...grpc.CallOption) (*GetDestinationsResponse, error) {
	out := new(GetDestinationsResponse)
	err := c.cc.Invoke(ctx, DestinationService_GetDestinations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationServiceClient) GetDestinationById(ctx context.Context, in *GetDestinationByIdRequest, opts ...grpc.CallOption) (*Destination, error) {
	out := new(Destination)
	err := c.cc.Invoke(ctx, DestinationService_GetDestinationById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationServiceClient) CreateDestination(ctx context.Context, in *CreateDestinationRequest, opts ...grpc.CallOption) (*CreateDestinationResponse, error) {
	out := new(CreateDestinationResponse)
	err := c.cc.Invoke(ctx, DestinationService_CreateDestination_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationServiceClient) UpdateDestination(ctx context.Context, in *UpdateDestinationRequest, opts ...grpc.CallOption) (*UpdateDestinationResponse, error) {
	out := new(UpdateDestinationResponse)
	err := c.cc.Invoke(ctx, DestinationService_UpdateDestination_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationServiceClient) DeleteDestination(ctx context.Context, in *DeleteDestinationRequest, opts ...grpc.CallOption) (*DeleteDestinationResponse, error) {
	out := new(DeleteDestinationResponse)
	err := c.cc.Invoke(ctx, DestinationService_DeleteDestination_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DestinationServiceServer is the server API for DestinationService service.
// All implementations must embed UnimplementedDestinationServiceServer
// for forward compatibility
type DestinationServiceServer interface {
	IsValid(context.Context, *ValidRequest) (*DestinationValid, error)
	GetDestinations(context.Context, *GetDestinationsRequest) (*GetDestinationsResponse, error)
	GetDestinationById(context.Context, *GetDestinationByIdRequest) (*Destination, error)
	CreateDestination(context.Context, *CreateDestinationRequest) (*CreateDestinationResponse, error)
	UpdateDestination(context.Context, *UpdateDestinationRequest) (*UpdateDestinationResponse, error)
	DeleteDestination(context.Context, *DeleteDestinationRequest) (*DeleteDestinationResponse, error)
	mustEmbedUnimplementedDestinationServiceServer()
}

// UnimplementedDestinationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDestinationServiceServer struct {
}

func (UnimplementedDestinationServiceServer) IsValid(context.Context, *ValidRequest) (*DestinationValid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsValid not implemented")
}
func (UnimplementedDestinationServiceServer) GetDestinations(context.Context, *GetDestinationsRequest) (*GetDestinationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDestinations not implemented")
}
func (UnimplementedDestinationServiceServer) GetDestinationById(context.Context, *GetDestinationByIdRequest) (*Destination, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDestinationById not implemented")
}
func (UnimplementedDestinationServiceServer) CreateDestination(context.Context, *CreateDestinationRequest) (*CreateDestinationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDestination not implemented")
}
func (UnimplementedDestinationServiceServer) UpdateDestination(context.Context, *UpdateDestinationRequest) (*UpdateDestinationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDestination not implemented")
}
func (UnimplementedDestinationServiceServer) DeleteDestination(context.Context, *DeleteDestinationRequest) (*DeleteDestinationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDestination not implemented")
}
func (UnimplementedDestinationServiceServer) mustEmbedUnimplementedDestinationServiceServer() {}

// UnsafeDestinationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DestinationServiceServer will
// result in compilation errors.
type UnsafeDestinationServiceServer interface {
	mustEmbedUnimplementedDestinationServiceServer()
}

func RegisterDestinationServiceServer(s grpc.ServiceRegistrar, srv DestinationServiceServer) {
	s.RegisterService(&DestinationService_ServiceDesc, srv)
}

func _DestinationService_IsValid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationServiceServer).IsValid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DestinationService_IsValid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationServiceServer).IsValid(ctx, req.(*ValidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DestinationService_GetDestinations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDestinationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationServiceServer).GetDestinations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DestinationService_GetDestinations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationServiceServer).GetDestinations(ctx, req.(*GetDestinationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DestinationService_GetDestinationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDestinationByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationServiceServer).GetDestinationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DestinationService_GetDestinationById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationServiceServer).GetDestinationById(ctx, req.(*GetDestinationByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DestinationService_CreateDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDestinationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationServiceServer).CreateDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DestinationService_CreateDestination_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationServiceServer).CreateDestination(ctx, req.(*CreateDestinationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DestinationService_UpdateDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDestinationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationServiceServer).UpdateDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DestinationService_UpdateDestination_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationServiceServer).UpdateDestination(ctx, req.(*UpdateDestinationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DestinationService_DeleteDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDestinationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationServiceServer).DeleteDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DestinationService_DeleteDestination_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationServiceServer).DeleteDestination(ctx, req.(*DeleteDestinationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DestinationService_ServiceDesc is the grpc.ServiceDesc for DestinationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DestinationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DestinationService",
	HandlerType: (*DestinationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "isValid",
			Handler:    _DestinationService_IsValid_Handler,
		},
		{
			MethodName: "GetDestinations",
			Handler:    _DestinationService_GetDestinations_Handler,
		},
		{
			MethodName: "GetDestinationById",
			Handler:    _DestinationService_GetDestinationById_Handler,
		},
		{
			MethodName: "CreateDestination",
			Handler:    _DestinationService_CreateDestination_Handler,
		},
		{
			MethodName: "UpdateDestination",
			Handler:    _DestinationService_UpdateDestination_Handler,
		},
		{
			MethodName: "DeleteDestination",
			Handler:    _DestinationService_DeleteDestination_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "destination_service.proto",
}
