package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	protectorpb "github.com/zaz600/brute-force-protector/api"
	"github.com/zaz600/brute-force-protector/internal/bruteforceprotector"
	"github.com/zaz600/brute-force-protector/internal/handler"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}

	protector := bruteforceprotector.NewBruteForceProtector(
		bruteforceprotector.WithLoginLimit(10),
		bruteforceprotector.WithPasswordLimit(100),
		bruteforceprotector.WithIPLimit(1000),
	)

	srv := handler.NewServer(protector)
	grpcServer := grpc.NewServer()
	protectorpb.RegisterBruteforceProtectorServiceServer(grpcServer, srv)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
