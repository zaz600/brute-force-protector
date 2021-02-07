package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	protectorpb "github.com/zaz600/brute-force-protector/api"
	"github.com/zaz600/brute-force-protector/internal/bruteforceprotector"
)

type server struct {
	protectorpb.UnimplementedBruteforceProtectorServiceServer
	protector *bruteforceprotector.BruteForceProtector
}

func (s *server) Verify(ctx context.Context, req *protectorpb.VerifyRequest) (*protectorpb.VerifyResponse, error) {
	log.Printf("verify with params: %v", req)
	ip := req.VerifyParams.GetIp()
	login := req.VerifyParams.GetIp()
	password := req.VerifyParams.GetIp()

	resp := &protectorpb.VerifyResponse{
		Ok: s.protector.Verify(ctx, login, password, ip),
	}
	return resp, nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}

	srv := &server{
		protector: bruteforceprotector.NewBruteForceProtector(
			bruteforceprotector.WithLoginLimit(10),
			bruteforceprotector.WithPasswordLimit(100),
			bruteforceprotector.WithIPdLimit(1000),
		),
	}
	grpcServer := grpc.NewServer()
	protectorpb.RegisterBruteforceProtectorServiceServer(grpcServer, srv)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}

	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < 90000; i++ {
	// 	result := bfProtector.Verify("foo", "password", "127.0.0.1")
	// 	log.Println(i, result)
	// 	time.Sleep(time.Duration(rand.Intn(5000-100+1)+100) * time.Millisecond)
	// }
}
