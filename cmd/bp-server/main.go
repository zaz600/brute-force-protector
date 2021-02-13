package main

import (
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"

	protectorpb "github.com/zaz600/brute-force-protector/api"
	"github.com/zaz600/brute-force-protector/internal/bruteforceprotector"
	"github.com/zaz600/brute-force-protector/internal/handler"
)

func main() {
	addr := flag.String("listen", "0.0.0.0:50051", "server host:port")
	loginLimit := flag.Int64("n", 10, "login limit per minute")
	passwordLimit := flag.Int64("m", 100, "password limit per minute")
	ipLimit := flag.Int64("k", 1000, "ip limit per minute")
	flag.Parse()

	listener, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}

	protector := bruteforceprotector.NewBruteForceProtector(
		bruteforceprotector.WithLoginLimit(*loginLimit),
		bruteforceprotector.WithPasswordLimit(*passwordLimit),
		bruteforceprotector.WithIPLimit(*ipLimit),
	)

	srv := handler.NewServer(protector)
	grpcServer := grpc.NewServer()
	protectorpb.RegisterBruteforceProtectorServiceServer(grpcServer, srv)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
