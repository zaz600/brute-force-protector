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
	ResetLogin(ctx context.Context, in *ResetLoginLimitRequest, opts ...grpc.CallOption) (*ResetLoginLimitResponse, error)
	ResetIP(ctx context.Context, in *ResetIPLimitRequest, opts ...grpc.CallOption) (*ResetIPLimitResponse, error)
	AddBlackList(ctx context.Context, in *AddBlackListRequest, opts ...grpc.CallOption) (*AddBlackListResponse, error)
	RemoveBlackList(ctx context.Context, in *RemoveBlackListRequest, opts ...grpc.CallOption) (*RemoveBlackListResponse, error)
	AddWhiteList(ctx context.Context, in *AddWhiteListRequest, opts ...grpc.CallOption) (*AddWhiteListResponse, error)
	RemoveWhiteList(ctx context.Context, in *RemoveWhiteListRequest, opts ...grpc.CallOption) (*RemoveWhiteListResponse, error)
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

func (c *bruteforceProtectorServiceClient) ResetLogin(ctx context.Context, in *ResetLoginLimitRequest, opts ...grpc.CallOption) (*ResetLoginLimitResponse, error) {
	out := new(ResetLoginLimitResponse)
	err := c.cc.Invoke(ctx, "/bruteforceprotector.BruteforceProtectorService/ResetLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bruteforceProtectorServiceClient) ResetIP(ctx context.Context, in *ResetIPLimitRequest, opts ...grpc.CallOption) (*ResetIPLimitResponse, error) {
	out := new(ResetIPLimitResponse)
	err := c.cc.Invoke(ctx, "/bruteforceprotector.BruteforceProtectorService/ResetIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bruteforceProtectorServiceClient) AddBlackList(ctx context.Context, in *AddBlackListRequest, opts ...grpc.CallOption) (*AddBlackListResponse, error) {
	out := new(AddBlackListResponse)
	err := c.cc.Invoke(ctx, "/bruteforceprotector.BruteforceProtectorService/AddBlackList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bruteforceProtectorServiceClient) RemoveBlackList(ctx context.Context, in *RemoveBlackListRequest, opts ...grpc.CallOption) (*RemoveBlackListResponse, error) {
	out := new(RemoveBlackListResponse)
	err := c.cc.Invoke(ctx, "/bruteforceprotector.BruteforceProtectorService/RemoveBlackList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bruteforceProtectorServiceClient) AddWhiteList(ctx context.Context, in *AddWhiteListRequest, opts ...grpc.CallOption) (*AddWhiteListResponse, error) {
	out := new(AddWhiteListResponse)
	err := c.cc.Invoke(ctx, "/bruteforceprotector.BruteforceProtectorService/AddWhiteList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bruteforceProtectorServiceClient) RemoveWhiteList(ctx context.Context, in *RemoveWhiteListRequest, opts ...grpc.CallOption) (*RemoveWhiteListResponse, error) {
	out := new(RemoveWhiteListResponse)
	err := c.cc.Invoke(ctx, "/bruteforceprotector.BruteforceProtectorService/RemoveWhiteList", in, out, opts...)
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
	ResetLogin(context.Context, *ResetLoginLimitRequest) (*ResetLoginLimitResponse, error)
	ResetIP(context.Context, *ResetIPLimitRequest) (*ResetIPLimitResponse, error)
	AddBlackList(context.Context, *AddBlackListRequest) (*AddBlackListResponse, error)
	RemoveBlackList(context.Context, *RemoveBlackListRequest) (*RemoveBlackListResponse, error)
	AddWhiteList(context.Context, *AddWhiteListRequest) (*AddWhiteListResponse, error)
	RemoveWhiteList(context.Context, *RemoveWhiteListRequest) (*RemoveWhiteListResponse, error)
	mustEmbedUnimplementedBruteforceProtectorServiceServer()
}

// UnimplementedBruteforceProtectorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBruteforceProtectorServiceServer struct {
}

func (UnimplementedBruteforceProtectorServiceServer) Verify(context.Context, *VerifyRequest) (*VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedBruteforceProtectorServiceServer) ResetLogin(context.Context, *ResetLoginLimitRequest) (*ResetLoginLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetLogin not implemented")
}
func (UnimplementedBruteforceProtectorServiceServer) ResetIP(context.Context, *ResetIPLimitRequest) (*ResetIPLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetIP not implemented")
}
func (UnimplementedBruteforceProtectorServiceServer) AddBlackList(context.Context, *AddBlackListRequest) (*AddBlackListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBlackList not implemented")
}
func (UnimplementedBruteforceProtectorServiceServer) RemoveBlackList(context.Context, *RemoveBlackListRequest) (*RemoveBlackListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBlackList not implemented")
}
func (UnimplementedBruteforceProtectorServiceServer) AddWhiteList(context.Context, *AddWhiteListRequest) (*AddWhiteListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWhiteList not implemented")
}
func (UnimplementedBruteforceProtectorServiceServer) RemoveWhiteList(context.Context, *RemoveWhiteListRequest) (*RemoveWhiteListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveWhiteList not implemented")
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

func _BruteforceProtectorService_ResetLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetLoginLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BruteforceProtectorServiceServer).ResetLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bruteforceprotector.BruteforceProtectorService/ResetLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BruteforceProtectorServiceServer).ResetLogin(ctx, req.(*ResetLoginLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BruteforceProtectorService_ResetIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetIPLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BruteforceProtectorServiceServer).ResetIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bruteforceprotector.BruteforceProtectorService/ResetIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BruteforceProtectorServiceServer).ResetIP(ctx, req.(*ResetIPLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BruteforceProtectorService_AddBlackList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBlackListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BruteforceProtectorServiceServer).AddBlackList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bruteforceprotector.BruteforceProtectorService/AddBlackList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BruteforceProtectorServiceServer).AddBlackList(ctx, req.(*AddBlackListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BruteforceProtectorService_RemoveBlackList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveBlackListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BruteforceProtectorServiceServer).RemoveBlackList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bruteforceprotector.BruteforceProtectorService/RemoveBlackList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BruteforceProtectorServiceServer).RemoveBlackList(ctx, req.(*RemoveBlackListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BruteforceProtectorService_AddWhiteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddWhiteListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BruteforceProtectorServiceServer).AddWhiteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bruteforceprotector.BruteforceProtectorService/AddWhiteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BruteforceProtectorServiceServer).AddWhiteList(ctx, req.(*AddWhiteListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BruteforceProtectorService_RemoveWhiteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveWhiteListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BruteforceProtectorServiceServer).RemoveWhiteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bruteforceprotector.BruteforceProtectorService/RemoveWhiteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BruteforceProtectorServiceServer).RemoveWhiteList(ctx, req.(*RemoveWhiteListRequest))
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
		{
			MethodName: "ResetLogin",
			Handler:    _BruteforceProtectorService_ResetLogin_Handler,
		},
		{
			MethodName: "ResetIP",
			Handler:    _BruteforceProtectorService_ResetIP_Handler,
		},
		{
			MethodName: "AddBlackList",
			Handler:    _BruteforceProtectorService_AddBlackList_Handler,
		},
		{
			MethodName: "RemoveBlackList",
			Handler:    _BruteforceProtectorService_RemoveBlackList_Handler,
		},
		{
			MethodName: "AddWhiteList",
			Handler:    _BruteforceProtectorService_AddWhiteList_Handler,
		},
		{
			MethodName: "RemoveWhiteList",
			Handler:    _BruteforceProtectorService_RemoveWhiteList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}
