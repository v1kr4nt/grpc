// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.2
// source: init_org.proto

package __

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
	InitOrgService_InitOrganization_FullMethodName = "/InitOrgService/InitOrganization"
)

// InitOrgServiceClient is the client API for InitOrgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InitOrgServiceClient interface {
	InitOrganization(ctx context.Context, in *InitOrgRequest, opts ...grpc.CallOption) (*InitOrgResponse, error)
}

type initOrgServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInitOrgServiceClient(cc grpc.ClientConnInterface) InitOrgServiceClient {
	return &initOrgServiceClient{cc}
}

func (c *initOrgServiceClient) InitOrganization(ctx context.Context, in *InitOrgRequest, opts ...grpc.CallOption) (*InitOrgResponse, error) {
	out := new(InitOrgResponse)
	err := c.cc.Invoke(ctx, InitOrgService_InitOrganization_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InitOrgServiceServer is the server API for InitOrgService service.
// All implementations must embed UnimplementedInitOrgServiceServer
// for forward compatibility
type InitOrgServiceServer interface {
	InitOrganization(context.Context, *InitOrgRequest) (*InitOrgResponse, error)
	mustEmbedUnimplementedInitOrgServiceServer()
}

// UnimplementedInitOrgServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInitOrgServiceServer struct {
}

func (UnimplementedInitOrgServiceServer) InitOrganization(context.Context, *InitOrgRequest) (*InitOrgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitOrganization not implemented")
}
func (UnimplementedInitOrgServiceServer) mustEmbedUnimplementedInitOrgServiceServer() {}

// UnsafeInitOrgServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InitOrgServiceServer will
// result in compilation errors.
type UnsafeInitOrgServiceServer interface {
	mustEmbedUnimplementedInitOrgServiceServer()
}

func RegisterInitOrgServiceServer(s grpc.ServiceRegistrar, srv InitOrgServiceServer) {
	s.RegisterService(&InitOrgService_ServiceDesc, srv)
}

func _InitOrgService_InitOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitOrgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InitOrgServiceServer).InitOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InitOrgService_InitOrganization_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InitOrgServiceServer).InitOrganization(ctx, req.(*InitOrgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InitOrgService_ServiceDesc is the grpc.ServiceDesc for InitOrgService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InitOrgService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "InitOrgService",
	HandlerType: (*InitOrgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitOrganization",
			Handler:    _InitOrgService_InitOrganization_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "init_org.proto",
}
