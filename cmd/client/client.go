package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	protectorpb "github.com/zaz600/brute-force-protector/api"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := protectorpb.NewBruteforceProtectorServiceClient(conn)
	resp, err := c.Verify(context.Background(), &protectorpb.VerifyRequest{VerifyParams: &protectorpb.VerifyParams{
		Login:    "admin",
		Password: "hash",
		Ip:       "127.0.0.1",
	}})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("resp: %v", resp)
}
