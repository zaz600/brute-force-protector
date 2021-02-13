package handler

import (
	"context"
	"fmt"
	"log"

	protectorpb "github.com/zaz600/brute-force-protector/api"
	"github.com/zaz600/brute-force-protector/internal/bruteforceprotector"
)

type Server struct {
	protectorpb.UnimplementedBruteforceProtectorServiceServer
	protector *bruteforceprotector.BruteForceProtector
}

func (s *Server) Verify(ctx context.Context, req *protectorpb.VerifyRequest) (*protectorpb.VerifyResponse, error) {
	log.Printf("verify with params: %v\n", req)
	ip := req.VerifyParams.GetIp()
	login := req.VerifyParams.GetLogin()
	password := req.VerifyParams.GetPassword()

	resp := &protectorpb.VerifyResponse{
		Ok: s.protector.Verify(ctx, login, password, ip),
	}
	log.Printf("verify result: %#v\n", resp.Ok)
	return resp, nil
}

func (s *Server) ResetLogin(ctx context.Context, req *protectorpb.ResetLoginLimitRequest) (*protectorpb.ResetLoginLimitResponse, error) {
	log.Printf("ResetLogin with params: %v\n", req)
	s.protector.ResetLogin(ctx, req.GetLogin())
	return &protectorpb.ResetLoginLimitResponse{}, nil
}

func (s *Server) ResetIP(ctx context.Context, req *protectorpb.ResetIPLimitRequest) (*protectorpb.ResetIPLimitResponse, error) {
	log.Printf("ResetIP with params: %v\n", req)
	s.protector.ResetLogin(ctx, req.GetLogin())
	return &protectorpb.ResetIPLimitResponse{}, nil
}

func (s *Server) AddBlackListItem(ctx context.Context, req *protectorpb.AddAccessListRequest) (*protectorpb.AddAccessListResponse, error) {
	resp := &protectorpb.AddAccessListResponse{Result: true}
	log.Printf("AddBlackList with params: %v\n", req)
	err := s.protector.AddBlackList(ctx, req.NetworkCIDR)
	if err != nil {
		resp.Result = false
		resp.Error = fmt.Sprintf("error add item to black list: %s", err)
	}
	return resp, nil
}

func (s *Server) AddWhiteListItem(ctx context.Context, req *protectorpb.AddAccessListRequest) (*protectorpb.AddAccessListResponse, error) {
	resp := &protectorpb.AddAccessListResponse{Result: true}
	log.Printf("AddWhiteList with params: %v\n", req)
	err := s.protector.AddWhiteList(ctx, req.NetworkCIDR)
	if err != nil {
		resp.Result = false
		resp.Error = fmt.Sprintf("error add item to white list: %s", err)
	}
	return resp, nil
}

func (s *Server) RemoveBlackList(ctx context.Context, req *protectorpb.RemoveAccessListRequest) (*protectorpb.RemoveAccessListResponse, error) {
	log.Printf("RemoveBlackList with params: %v\n", req)
	s.protector.RemoveBlackList(ctx, req.NetworkCIDR)
	return &protectorpb.RemoveAccessListResponse{}, nil
}

func (s *Server) RemoveWhiteList(ctx context.Context, req *protectorpb.RemoveAccessListRequest) (*protectorpb.RemoveAccessListResponse, error) {
	log.Printf("RemoveWhiteList with params: %v\n", req)
	s.protector.RemoveWhiteList(ctx, req.NetworkCIDR)
	return &protectorpb.RemoveAccessListResponse{}, nil
}

func (s *Server) GetBlackListItems(ctx context.Context, req *protectorpb.GetAccessListItemsRequest) (*protectorpb.GetAccessListItemsResponse, error) {
	return &protectorpb.GetAccessListItemsResponse{Items: s.protector.BlackListItems(ctx)}, nil
}

func (s *Server) GetWhiteListItems(ctx context.Context, req *protectorpb.GetAccessListItemsRequest) (*protectorpb.GetAccessListItemsResponse, error) {
	return &protectorpb.GetAccessListItemsResponse{Items: s.protector.WhiteListItems(ctx)}, nil
}

func NewServer(protector *bruteforceprotector.BruteForceProtector) *Server {
	return &Server{
		protector: protector,
	}
}
