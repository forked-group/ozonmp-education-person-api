// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package education_person_api

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

// EducationPersonApiServiceClient is the client API for EducationPersonApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EducationPersonApiServiceClient interface {
	// DescribePersonV1 - Describe a person
	DescribePersonV1(ctx context.Context, in *DescribePersonV1Request, opts ...grpc.CallOption) (*DescribePersonV1Response, error)
	ListPersonV1(ctx context.Context, in *ListPersonV1Request, opts ...grpc.CallOption) (*ListPersonV1Response, error)
	CreatePersonV1(ctx context.Context, in *CreatePersonV1Request, opts ...grpc.CallOption) (*CreatePersonV1Response, error)
	UpdatePersonV1(ctx context.Context, in *UpdatePersonV1Request, opts ...grpc.CallOption) (*UpdatePersonV1Response, error)
	RemovePersonV1(ctx context.Context, in *RemovePersonV1Request, opts ...grpc.CallOption) (*RemovePersonV1Response, error)
}

type educationPersonApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEducationPersonApiServiceClient(cc grpc.ClientConnInterface) EducationPersonApiServiceClient {
	return &educationPersonApiServiceClient{cc}
}

func (c *educationPersonApiServiceClient) DescribePersonV1(ctx context.Context, in *DescribePersonV1Request, opts ...grpc.CallOption) (*DescribePersonV1Response, error) {
	out := new(DescribePersonV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.education_person_api.v1.EducationPersonApiService/DescribePersonV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *educationPersonApiServiceClient) ListPersonV1(ctx context.Context, in *ListPersonV1Request, opts ...grpc.CallOption) (*ListPersonV1Response, error) {
	out := new(ListPersonV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.education_person_api.v1.EducationPersonApiService/ListPersonV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *educationPersonApiServiceClient) CreatePersonV1(ctx context.Context, in *CreatePersonV1Request, opts ...grpc.CallOption) (*CreatePersonV1Response, error) {
	out := new(CreatePersonV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.education_person_api.v1.EducationPersonApiService/CreatePersonV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *educationPersonApiServiceClient) UpdatePersonV1(ctx context.Context, in *UpdatePersonV1Request, opts ...grpc.CallOption) (*UpdatePersonV1Response, error) {
	out := new(UpdatePersonV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.education_person_api.v1.EducationPersonApiService/UpdatePersonV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *educationPersonApiServiceClient) RemovePersonV1(ctx context.Context, in *RemovePersonV1Request, opts ...grpc.CallOption) (*RemovePersonV1Response, error) {
	out := new(RemovePersonV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.education_person_api.v1.EducationPersonApiService/RemovePersonV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EducationPersonApiServiceServer is the server API for EducationPersonApiService service.
// All implementations must embed UnimplementedEducationPersonApiServiceServer
// for forward compatibility
type EducationPersonApiServiceServer interface {
	// DescribePersonV1 - Describe a person
	DescribePersonV1(context.Context, *DescribePersonV1Request) (*DescribePersonV1Response, error)
	ListPersonV1(context.Context, *ListPersonV1Request) (*ListPersonV1Response, error)
	CreatePersonV1(context.Context, *CreatePersonV1Request) (*CreatePersonV1Response, error)
	UpdatePersonV1(context.Context, *UpdatePersonV1Request) (*UpdatePersonV1Response, error)
	RemovePersonV1(context.Context, *RemovePersonV1Request) (*RemovePersonV1Response, error)
	mustEmbedUnimplementedEducationPersonApiServiceServer()
}

// UnimplementedEducationPersonApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEducationPersonApiServiceServer struct {
}

func (UnimplementedEducationPersonApiServiceServer) DescribePersonV1(context.Context, *DescribePersonV1Request) (*DescribePersonV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePersonV1 not implemented")
}
func (UnimplementedEducationPersonApiServiceServer) ListPersonV1(context.Context, *ListPersonV1Request) (*ListPersonV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPersonV1 not implemented")
}
func (UnimplementedEducationPersonApiServiceServer) CreatePersonV1(context.Context, *CreatePersonV1Request) (*CreatePersonV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePersonV1 not implemented")
}
func (UnimplementedEducationPersonApiServiceServer) UpdatePersonV1(context.Context, *UpdatePersonV1Request) (*UpdatePersonV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePersonV1 not implemented")
}
func (UnimplementedEducationPersonApiServiceServer) RemovePersonV1(context.Context, *RemovePersonV1Request) (*RemovePersonV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePersonV1 not implemented")
}
func (UnimplementedEducationPersonApiServiceServer) mustEmbedUnimplementedEducationPersonApiServiceServer() {
}

// UnsafeEducationPersonApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EducationPersonApiServiceServer will
// result in compilation errors.
type UnsafeEducationPersonApiServiceServer interface {
	mustEmbedUnimplementedEducationPersonApiServiceServer()
}

func RegisterEducationPersonApiServiceServer(s grpc.ServiceRegistrar, srv EducationPersonApiServiceServer) {
	s.RegisterService(&EducationPersonApiService_ServiceDesc, srv)
}

func _EducationPersonApiService_DescribePersonV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePersonV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EducationPersonApiServiceServer).DescribePersonV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.education_person_api.v1.EducationPersonApiService/DescribePersonV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EducationPersonApiServiceServer).DescribePersonV1(ctx, req.(*DescribePersonV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _EducationPersonApiService_ListPersonV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPersonV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EducationPersonApiServiceServer).ListPersonV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.education_person_api.v1.EducationPersonApiService/ListPersonV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EducationPersonApiServiceServer).ListPersonV1(ctx, req.(*ListPersonV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _EducationPersonApiService_CreatePersonV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePersonV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EducationPersonApiServiceServer).CreatePersonV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.education_person_api.v1.EducationPersonApiService/CreatePersonV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EducationPersonApiServiceServer).CreatePersonV1(ctx, req.(*CreatePersonV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _EducationPersonApiService_UpdatePersonV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePersonV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EducationPersonApiServiceServer).UpdatePersonV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.education_person_api.v1.EducationPersonApiService/UpdatePersonV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EducationPersonApiServiceServer).UpdatePersonV1(ctx, req.(*UpdatePersonV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _EducationPersonApiService_RemovePersonV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePersonV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EducationPersonApiServiceServer).RemovePersonV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.education_person_api.v1.EducationPersonApiService/RemovePersonV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EducationPersonApiServiceServer).RemovePersonV1(ctx, req.(*RemovePersonV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// EducationPersonApiService_ServiceDesc is the grpc.ServiceDesc for EducationPersonApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EducationPersonApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ozonmp.education_person_api.v1.EducationPersonApiService",
	HandlerType: (*EducationPersonApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribePersonV1",
			Handler:    _EducationPersonApiService_DescribePersonV1_Handler,
		},
		{
			MethodName: "ListPersonV1",
			Handler:    _EducationPersonApiService_ListPersonV1_Handler,
		},
		{
			MethodName: "CreatePersonV1",
			Handler:    _EducationPersonApiService_CreatePersonV1_Handler,
		},
		{
			MethodName: "UpdatePersonV1",
			Handler:    _EducationPersonApiService_UpdatePersonV1_Handler,
		},
		{
			MethodName: "RemovePersonV1",
			Handler:    _EducationPersonApiService_RemovePersonV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ozonmp/education_person_api/v1/education_person_api.proto",
}
