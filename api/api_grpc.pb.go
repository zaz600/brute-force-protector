// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protectorpb

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

// BruteforceProtectorServiceClient is the client API for BruteforceProtectorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BruteforceProtectorServiceClient interface {
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error)
}

type bruteforceProtectorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBruteforceProtectorServiceClient(cc grpc.ClientConnInterface) BruteforceProtectorServiceClient {
	return &bruteforceProtectorServiceClient{cc}
}

func (c *bruteforceProtectorServiceClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error) {
	out := new(VerifyResponse)
	err := c.cc.Invoke(ctx, "/bruteforceprotector.BruteforceProtectorService/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BruteforceProtectorServiceServer is the server API for BruteforceProtectorService service.
// All implementations must embed UnimplementedBruteforceProtectorServiceServer
// for forward compatibility
type BruteforceProtectorServiceServer interface {
	Verify(context.Context, *VerifyRequest) (*VerifyResponse, error)
	mustEmbedUnimplementedBruteforceProtectorServiceServer()
}

// UnimplementedBruteforceProtectorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBruteforceProtectorServiceServer struct {
}

func (UnimplementedBruteforceProtectorServiceServer) Verify(context.Context, *VerifyRequest) (*VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedBruteforceProtectorServiceServer) mustEmbedUnimplementedBruteforceProtectorServiceServer() {
}

// UnsafeBruteforceProtectorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BruteforceProtectorServiceServer will
// result in compilation errors.
type UnsafeBruteforceProtectorServiceServer interface {
	mustEmbedUnimplementedBruteforceProtectorServiceServer()
}

func RegisterBruteforceProtectorServiceServer(s grpc.ServiceRegistrar, srv BruteforceProtectorServiceServer) {
	s.RegisterService(&BruteforceProtectorService_ServiceDesc, srv)
}

func _BruteforceProtectorService_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BruteforceProtectorServiceServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bruteforceprotector.BruteforceProtectorService/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BruteforceProtectorServiceServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BruteforceProtectorService_ServiceDesc is the grpc.ServiceDesc for BruteforceProtectorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BruteforceProtectorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bruteforceprotector.BruteforceProtectorService",
	HandlerType: (*BruteforceProtectorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Verify",
			Handler:    _BruteforceProtectorService_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}
