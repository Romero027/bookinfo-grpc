// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: proto/details/details.proto

package details

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

// DetailsClient is the client API for Details service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DetailsClient interface {
	GetRatings(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Result, error)
}

type detailsClient struct {
	cc grpc.ClientConnInterface
}

func NewDetailsClient(cc grpc.ClientConnInterface) DetailsClient {
	return &detailsClient{cc}
}

func (c *detailsClient) GetRatings(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/details.Details/getRatings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DetailsServer is the server API for Details service.
// All implementations must embed UnimplementedDetailsServer
// for forward compatibility
type DetailsServer interface {
	GetRatings(context.Context, *Product) (*Result, error)
	mustEmbedUnimplementedDetailsServer()
}

// UnimplementedDetailsServer must be embedded to have forward compatible implementations.
type UnimplementedDetailsServer struct {
}

func (UnimplementedDetailsServer) GetRatings(context.Context, *Product) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRatings not implemented")
}
func (UnimplementedDetailsServer) mustEmbedUnimplementedDetailsServer() {}

// UnsafeDetailsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DetailsServer will
// result in compilation errors.
type UnsafeDetailsServer interface {
	mustEmbedUnimplementedDetailsServer()
}

func RegisterDetailsServer(s grpc.ServiceRegistrar, srv DetailsServer) {
	s.RegisterService(&Details_ServiceDesc, srv)
}

func _Details_GetRatings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetailsServer).GetRatings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/details.Details/getRatings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetailsServer).GetRatings(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

// Details_ServiceDesc is the grpc.ServiceDesc for Details service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Details_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "details.Details",
	HandlerType: (*DetailsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getRatings",
			Handler:    _Details_GetRatings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/details/details.proto",
}