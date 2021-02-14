package main

import (
	"flag"
	"log"

	"github.com/zaz600/brute-force-protector/internal/bruteforceprotector"
	"github.com/zaz600/brute-force-protector/internal/grpc"
)

func main() {
	addr := flag.String("listen", "0.0.0.0:50051", "server host:port")
	loginLimit := flag.Int64("n", 10, "login limit per minute")
	passwordLimit := flag.Int64("m", 100, "password limit per minute")
	ipLimit := flag.Int64("k", 1000, "ip limit per minute")
	flag.Parse()

	protector := bruteforceprotector.NewBruteForceProtector(
		bruteforceprotector.WithLoginLimit(*loginLimit),
		bruteforceprotector.WithPasswordLimit(*passwordLimit),
		bruteforceprotector.WithIPLimit(*ipLimit),
	)
	bpServer := grpc.NewBPServer(protector)
	err := bpServer.ListenAndServe(*addr)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
