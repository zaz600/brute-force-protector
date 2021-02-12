package handler

import (
	"context"
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

func (s *Server) AddBlackList(ctx context.Context, req *protectorpb.AddBlackListRequest) (*protectorpb.AddBlackListResponse, error) {
	// TODO логировать ошибку или вернуть ее в ответ
	log.Printf("AddBlackList with params: %v\n", req)
	_ = s.protector.AddBlackList(ctx, req.NetworkCIDR)
	return &protectorpb.AddBlackListResponse{}, nil
}

func (s *Server) RemoveBlackList(ctx context.Context, req *protectorpb.RemoveBlackListRequest) (*protectorpb.RemoveBlackListResponse, error) {
	log.Printf("RemoveBlackList with params: %v\n", req)
	s.protector.RemoveBlackList(ctx, req.NetworkCIDR)
	return &protectorpb.RemoveBlackListResponse{}, nil
}

func (s *Server) AddWhiteList(ctx context.Context, req *protectorpb.AddWhiteListRequest) (*protectorpb.AddWhiteListResponse, error) {
	// TODO логировать ошибку или вернуть ее в ответ
	log.Printf("AddWhiteList with params: %v\n", req)
	_ = s.protector.AddWhiteList(ctx, req.NetworkCIDR)
	return &protectorpb.AddWhiteListResponse{}, nil
}

func (s *Server) RemoveWhiteList(ctx context.Context, req *protectorpb.RemoveWhiteListRequest) (*protectorpb.RemoveWhiteListResponse, error) {
	log.Printf("RemoveWhiteList with params: %v\n", req)
	s.protector.RemoveWhiteList(ctx, req.NetworkCIDR)
	return &protectorpb.RemoveWhiteListResponse{}, nil
}

func (s *Server) GetBlackListItems(ctx context.Context, req *protectorpb.GetBlackListItemsRequest) (*protectorpb.GetBlackListItemsResponse, error) {
	return &protectorpb.GetBlackListItemsResponse{Items: s.protector.BlackListItems(ctx)}, nil
}

func (s *Server) GetWhiteListItems(ctx context.Context, req *protectorpb.GetWhiteListItemsRequest) (*protectorpb.GetWhiteListItemsResponse, error) {
	return &protectorpb.GetWhiteListItemsResponse{Items: s.protector.WhiteListItems(ctx)}, nil
}

func NewServer(protector *bruteforceprotector.BruteForceProtector) *Server {
	return &Server{
		protector: protector,
	}
}
