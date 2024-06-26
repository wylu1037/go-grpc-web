// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: tblock/v1/tblock.proto

package tblockv1

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

// TBlockServiceClient is the client API for TBlockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TBlockServiceClient interface {
	Details(ctx context.Context, in *TBlockServiceDetailsRequest, opts ...grpc.CallOption) (*TBlockServiceDetailsResponse, error)
}

type tBlockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTBlockServiceClient(cc grpc.ClientConnInterface) TBlockServiceClient {
	return &tBlockServiceClient{cc}
}

func (c *tBlockServiceClient) Details(ctx context.Context, in *TBlockServiceDetailsRequest, opts ...grpc.CallOption) (*TBlockServiceDetailsResponse, error) {
	out := new(TBlockServiceDetailsResponse)
	err := c.cc.Invoke(ctx, "/tblock.v1.TBlockService/Details", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TBlockServiceServer is the server API for TBlockService service.
// All implementations must embed UnimplementedTBlockServiceServer
// for forward compatibility
type TBlockServiceServer interface {
	Details(context.Context, *TBlockServiceDetailsRequest) (*TBlockServiceDetailsResponse, error)
	mustEmbedUnimplementedTBlockServiceServer()
}

// UnimplementedTBlockServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTBlockServiceServer struct {
}

func (UnimplementedTBlockServiceServer) Details(context.Context, *TBlockServiceDetailsRequest) (*TBlockServiceDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Details not implemented")
}
func (UnimplementedTBlockServiceServer) mustEmbedUnimplementedTBlockServiceServer() {}

// UnsafeTBlockServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TBlockServiceServer will
// result in compilation errors.
type UnsafeTBlockServiceServer interface {
	mustEmbedUnimplementedTBlockServiceServer()
}

func RegisterTBlockServiceServer(s grpc.ServiceRegistrar, srv TBlockServiceServer) {
	s.RegisterService(&TBlockService_ServiceDesc, srv)
}

func _TBlockService_Details_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TBlockServiceDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TBlockServiceServer).Details(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tblock.v1.TBlockService/Details",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TBlockServiceServer).Details(ctx, req.(*TBlockServiceDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TBlockService_ServiceDesc is the grpc.ServiceDesc for TBlockService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TBlockService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tblock.v1.TBlockService",
	HandlerType: (*TBlockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Details",
			Handler:    _TBlockService_Details_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tblock/v1/tblock.proto",
}
