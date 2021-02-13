package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"

	protectorpb "github.com/zaz600/brute-force-protector/api"
)

func main() {
	var host string

	app := &cli.App{
		Name:  "bp-cli",
		Usage: "Bruteforce Protector Client",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "host",
				Value:       "127.0.0.1:50051",
				Usage:       "bruteforce protector host",
				Destination: &host,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "blacklist",
				Aliases: []string{"bl"},
				Usage:   "black list operations",
				Subcommands: []*cli.Command{
					{
						Name:      "add",
						Usage:     "add an item to blacklist",
						ArgsUsage: "networkCIDR",
						Action: func(c *cli.Context) error {
							if c.NArg() != 1 {
								return cli.Exit("missed networkCIDR arg", 10)
							}
							client, err := newBpClient(host)
							if err != nil {
								return cli.Exit(fmt.Sprintf("error connect to host %s: %v", host, err), 9)
							}
							defer client.conn.Close()

							req := &protectorpb.AddBlackListRequest{NetworkCIDR: c.Args().First()}
							result, err := client.client.AddBlackList(context.TODO(), req)
							if err != nil {
								return cli.Exit(fmt.Sprintf("error add item to blacklist: %v", err), 9)
							}
							if !result.Result {
								return cli.Exit(fmt.Sprintf("error add item to blacklist: %v", result.Error), 9)
							}
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "remove an item to blacklist",
						Action: func(c *cli.Context) error {
							if c.NArg() != 1 {
								return cli.Exit("missed networkCIDR arg", 10)
							}

							client, err := newBpClient(host)
							if err != nil {
								return cli.Exit(fmt.Sprintf("error connect to host %s: %v", host, err), 9)
							}
							defer client.conn.Close()

							req := &protectorpb.RemoveBlackListRequest{NetworkCIDR: c.Args().First()}
							_, err = client.client.RemoveBlackList(context.TODO(), req)
							if err != nil {
								return cli.Exit(fmt.Sprintf("error remove item from blacklist: %v", err), 9)
							}
							return nil
						},
					},
					{
						Name:  "show",
						Usage: "show blacklist items",
						Action: func(c *cli.Context) error {
							if c.NArg() != 0 {
								return cli.Exit("unknown argument", 10)
							}

							client, err := newBpClient(host)
							if err != nil {
								return cli.Exit(fmt.Sprintf("error connect to host %s: %v", host, err), 9)
							}
							defer client.conn.Close()

							req := &protectorpb.GetBlackListItemsRequest{}
							resp, err := client.client.GetBlackListItems(context.TODO(), req)
							if err != nil {
								return cli.Exit(fmt.Sprintf("error get items from blacklist: %v", err), 9)
							}

							fmt.Println("Blacklist items:")
							for _, item := range resp.Items {
								fmt.Printf("- %s\n", item)
							}
							return nil
						},
					},
				},
			},
			{
				Name:    "whitelist",
				Aliases: []string{"wl"},
				Usage:   "white list operations",
				Subcommands: []*cli.Command{
					{
						Name:      "add",
						Usage:     "add an item to whitelist",
						ArgsUsage: "networkCIDR",
						Action: func(c *cli.Context) error {
							if c.NArg() != 1 {
								return cli.Exit("missed networkCIDR arg", 10)
							}
							client, err := newBpClient(host)
							if err != nil {
								return cli.Exit(fmt.Sprintf("error connect to host %s: %v", host, err), 9)
							}
							defer client.conn.Close()

							req := &protectorpb.AddWhiteListRequest{NetworkCIDR: c.Args().First()}
							result, err := client.client.AddWhiteList(context.TODO(), req)
							if err != nil {
								return cli.Exit(fmt.Sprintf("error add item to whitelist: %v", err), 9)
							}
							if !result.Result {
								return cli.Exit(fmt.Sprintf("error add item to whitelist: %v", result.Error), 9)
							}
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "remove an item to whitelist",
						Action: func(c *cli.Context) error {
							if c.NArg() != 1 {
								return cli.Exit("missed networkCIDR arg", 10)
							}

							client, err := newBpClient(host)
							if err != nil {
								return cli.Exit(fmt.Sprintf("error connect to host %s: %v", host, err), 9)
							}
							defer client.conn.Close()

							req := &protectorpb.RemoveWhiteListRequest{NetworkCIDR: c.Args().First()}
							_, err = client.client.RemoveWhiteList(context.TODO(), req)
							if err != nil {
								return cli.Exit(fmt.Sprintf("error remove item from whitelist: %v", err), 9)
							}
							return nil
						},
					},
					{
						Name:  "show",
						Usage: "show whitelist items",
						Action: func(c *cli.Context) error {
							if c.NArg() != 0 {
								return cli.Exit("unknown argument", 10)
							}

							client, err := newBpClient(host)
							if err != nil {
								return cli.Exit(fmt.Sprintf("error connect to host %s: %v", host, err), 9)
							}
							defer client.conn.Close()

							req := &protectorpb.GetWhiteListItemsRequest{}
							resp, err := client.client.GetWhiteListItems(context.TODO(), req)
							if err != nil {
								return cli.Exit(fmt.Sprintf("error get items from whitelist: %v", err), 9)
							}

							fmt.Println("Whitelist items:")
							for _, item := range resp.Items {
								fmt.Printf("- %s\n", item)
							}
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

type bpClient struct {
	conn   *grpc.ClientConn
	client protectorpb.BruteforceProtectorServiceClient
}

func newBpClient(server string) (*bpClient, error) {
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := protectorpb.NewBruteforceProtectorServiceClient(conn)

	return &bpClient{
		conn:   conn,
		client: client,
	}, nil
}
