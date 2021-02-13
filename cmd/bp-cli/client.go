package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"

	protectorpb "github.com/zaz600/brute-force-protector/api"
)

func main() {
	os.Exit(CLI(os.Args))
}

func CLI(args []string) int {
	app := createApp()
	if err := app.Run(args); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 10
	}
	return 0
}

func createApp() *cli.App {
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
		Commands: []*cli.Command{{
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

						service := bpService{host: host}
						if err := service.addAccessList(c.Args().First(), Black); err != nil {
							return cli.Exit(err, 9)
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

						service := bpService{host: host}
						if err := service.removeAccessList(c.Args().First(), Black); err != nil {
							return cli.Exit(err, 9)
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

						service := bpService{host: host}
						items, err := service.getAccessListItems(Black)
						if err != nil {
							return cli.Exit(err, 9)
						}

						fmt.Println("Blacklist items:")
						for _, item := range items {
							fmt.Printf("- %s\n", item)
						}
						return nil
					},
				},
			},
		}, {
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

						service := bpService{host: host}
						if err := service.addAccessList(c.Args().First(), White); err != nil {
							return cli.Exit(err, 9)
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

						service := bpService{host: host}
						if err := service.removeAccessList(c.Args().First(), White); err != nil {
							return cli.Exit(err, 9)
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

						service := bpService{host: host}
						items, err := service.getAccessListItems(White)
						if err != nil {
							return cli.Exit(err, 9)
						}

						fmt.Println("Whitelist items:")
						for _, item := range items {
							fmt.Printf("- %s\n", item)
						}
						return nil
					},
				},
			},
		},
		},
	}
	return app
}

type bpClient struct {
	conn      *grpc.ClientConn
	rpcClient protectorpb.BruteforceProtectorServiceClient
}

func newBpClient(server string) (*bpClient, error) {
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := protectorpb.NewBruteforceProtectorServiceClient(conn)

	return &bpClient{
		conn:      conn,
		rpcClient: client,
	}, nil
}

type ListType int

const (
	Black ListType = iota
	White
)

type bpService struct {
	host string
}

func (s bpService) addAccessList(item string, listType ListType) error {
	client, err := newBpClient(s.host)
	if err != nil {
		return fmt.Errorf("error connect to host %s: %w", s.host, err)
	}
	defer client.conn.Close()

	req := &protectorpb.AddAccessListRequest{NetworkCIDR: item}

	var result *protectorpb.AddAccessListResponse
	switch listType {
	case Black:
		result, err = client.rpcClient.AddBlackListItem(context.TODO(), req)
	case White:
		result, err = client.rpcClient.AddWhiteListItem(context.TODO(), req)
	}

	if err != nil {
		return fmt.Errorf("error add item to access list: %w", err)
	}
	if !result.Result {
		return fmt.Errorf("error add item to access list: %s", result.Error)
	}
	return nil
}

func (s bpService) removeAccessList(item string, listType ListType) error {
	client, err := newBpClient(s.host)
	if err != nil {
		return fmt.Errorf("error connect to host %s: %w", s.host, err)
	}
	defer client.conn.Close()

	req := &protectorpb.RemoveAccessListRequest{NetworkCIDR: item}
	switch listType {
	case Black:
		_, err = client.rpcClient.RemoveBlackListItem(context.TODO(), req)
	case White:
		_, err = client.rpcClient.RemoveWhiteListItem(context.TODO(), req)
	}

	if err != nil {
		return fmt.Errorf("error remove item from access list: %w", err)
	}
	return nil
}

func (s bpService) getAccessListItems(listType ListType) ([]string, error) {
	client, err := newBpClient(s.host)
	if err != nil {
		return nil, fmt.Errorf("error connect to host %s: %w", s.host, err)
	}
	defer client.conn.Close()
	req := &protectorpb.GetAccessListItemsRequest{}

	var result *protectorpb.GetAccessListItemsResponse
	switch listType {
	case Black:
		result, err = client.rpcClient.GetBlackListItems(context.TODO(), req)
	case White:
		result, err = client.rpcClient.GetWhiteListItems(context.TODO(), req)
	}

	if err != nil {
		return nil, fmt.Errorf("error get items from access list: %w", err)
	}
	return result.Items, nil
}
