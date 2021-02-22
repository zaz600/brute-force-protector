package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/zaz600/brute-force-protector/internal/bruteforceprotector"
	protectorpb "github.com/zaz600/brute-force-protector/internal/grpc/api"
)

type BPServer struct {
	protectorpb.UnimplementedBruteforceProtectorServiceServer
	bp         *bruteforceprotector.BruteForceProtector
	grpcServer *grpc.Server
}

func NewBPServer(bp *bruteforceprotector.BruteForceProtector) *BPServer {
	return &BPServer{
		bp:         bp,
		grpcServer: grpc.NewServer(),
	}
}

func (b *BPServer) GracefulStop() {
	b.grpcServer.GracefulStop()
}

func (b *BPServer) ListenAndServe(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	log.Println("start BP Server on", addr)

	protectorpb.RegisterBruteforceProtectorServiceServer(b.grpcServer, b)
	return b.grpcServer.Serve(listener)
}

func (b *BPServer) Verify(ctx context.Context, req *protectorpb.VerifyRequest) (*protectorpb.VerifyResponse, error) {
	log.Printf("verify with params: %v\n", req)
	ip := req.VerifyParams.GetIp()
	login := req.VerifyParams.GetLogin()
	password := req.VerifyParams.GetPassword()

	resp := &protectorpb.VerifyResponse{
		Ok: b.bp.Verify(ctx, login, password, ip),
	}
	log.Printf("verify result: %#v\n", resp.Ok)
	return resp, nil
}

func (b *BPServer) ResetLogin(ctx context.Context, req *protectorpb.ResetLoginLimitRequest) (*protectorpb.ResetLimitResponse, error) {
	log.Printf("ResetLogin with params: %v\n", req)
	b.bp.ResetLogin(ctx, req.GetLogin())
	return &protectorpb.ResetLimitResponse{}, nil
}

func (b *BPServer) ResetIP(ctx context.Context, req *protectorpb.ResetIPLimitRequest) (*protectorpb.ResetLimitResponse, error) {
	log.Printf("ResetIP with params: %v\n", req)
	b.bp.ResetIP(ctx, req.GetIp())
	return &protectorpb.ResetLimitResponse{}, nil
}

func (b *BPServer) AddBlackListItem(ctx context.Context, req *protectorpb.AddAccessListRequest) (*protectorpb.AddAccessListResponse, error) {
	resp := &protectorpb.AddAccessListResponse{Result: true}
	log.Printf("AddBlackList with params: %v\n", req)
	err := b.bp.AddBlackList(ctx, req.NetworkCIDR)
	if err != nil {
		resp.Result = false
		resp.Error = fmt.Sprintf("error add item to black list: %s", err)
	}
	return resp, nil
}

func (b *BPServer) AddWhiteListItem(ctx context.Context, req *protectorpb.AddAccessListRequest) (*protectorpb.AddAccessListResponse, error) {
	resp := &protectorpb.AddAccessListResponse{Result: true}
	log.Printf("AddWhiteList with params: %v\n", req)
	err := b.bp.AddWhiteList(ctx, req.NetworkCIDR)
	if err != nil {
		resp.Result = false
		resp.Error = fmt.Sprintf("error add item to white list: %s", err)
	}
	return resp, nil
}

func (b *BPServer) RemoveBlackListItem(ctx context.Context, req *protectorpb.RemoveAccessListRequest) (*protectorpb.RemoveAccessListResponse, error) {
	log.Printf("RemoveBlackList with params: %v\n", req)
	resp := &protectorpb.RemoveAccessListResponse{}
	err := b.bp.RemoveBlackList(ctx, req.NetworkCIDR)
	if err != nil {
		resp.Result = false
		resp.Error = fmt.Sprintf("error remove item from black list: %s", err)
	}
	return resp, nil
}

func (b *BPServer) RemoveWhiteListItem(ctx context.Context, req *protectorpb.RemoveAccessListRequest) (*protectorpb.RemoveAccessListResponse, error) {
	log.Printf("RemoveWhiteList with params: %v\n", req)
	resp := &protectorpb.RemoveAccessListResponse{}
	err := b.bp.RemoveWhiteList(ctx, req.NetworkCIDR)
	if err != nil {
		resp.Result = false
		resp.Error = fmt.Sprintf("error remove item from white list: %s", err)
	}
	return resp, nil
}

func (b *BPServer) GetBlackListItems(ctx context.Context, req *protectorpb.GetAccessListItemsRequest) (*protectorpb.GetAccessListItemsResponse, error) {
	return &protectorpb.GetAccessListItemsResponse{Items: b.bp.BlackListItems(ctx)}, nil
}

func (b *BPServer) GetWhiteListItems(ctx context.Context, req *protectorpb.GetAccessListItemsRequest) (*protectorpb.GetAccessListItemsResponse, error) {
	return &protectorpb.GetAccessListItemsResponse{Items: b.bp.WhiteListItems(ctx)}, nil
}
