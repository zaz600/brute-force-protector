package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"

	protectorpb "github.com/zaz600/brute-force-protector/internal/grpc/api"
)

type ListType string

const (
	Black ListType = "BlackList"
	White ListType = "WhiteList"
)

var clientTimeout = 5 * time.Second // TODO - config

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
		Name:     "bp-cli",
		HelpName: "bp-cli",
		Usage:    "Bruteforce Protector Client",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "server",
				Value:       "127.0.0.1:50051",
				Usage:       "bruteforce protector server",
				Destination: &host,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "blacklist",
				Aliases: []string{"bl"},
				Usage:   "black list operations",
				Subcommands: []*cli.Command{
					createAddAccessListCommand(Black),
					createRemoveAccessListCommand(Black),
					createShowAccessListItemsCommand(Black),
				},
			},
			{
				Name:    "whitelist",
				Aliases: []string{"wl"},
				Usage:   "white list operations",
				Subcommands: []*cli.Command{
					createAddAccessListCommand(White),
					createRemoveAccessListCommand(White),
					createShowAccessListItemsCommand(White),
				},
			},
			{
				Name:  "reset",
				Usage: "reset limits",
				Subcommands: []*cli.Command{
					createResetLoginLimitCommand(),
					createResetIPLimitCommand(),
				},
			},
		},
	}
	return app
}

func createAddAccessListCommand(listType ListType) *cli.Command {
	return &cli.Command{
		Name:      "add",
		Usage:     fmt.Sprintf("add an item to %s", listType),
		ArgsUsage: "<networkCIDR>",
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				return cli.Exit("missed networkCIDR arg", 10)
			}

			ctx, cancel := context.WithTimeout(context.Background(), clientTimeout)
			defer cancel()
			client := newClient(ctx, c.String("server"))
			req := &protectorpb.AddAccessListRequest{NetworkCIDR: c.Args().First()}

			var result *protectorpb.AddAccessListResponse
			var err error
			switch listType {
			case Black:
				result, err = client.AddBlackListItem(ctx, req)
			case White:
				result, err = client.AddWhiteListItem(ctx, req)
			}

			if err != nil {
				return cli.Exit(err, 10)
			}
			log.Println("Done: ", result)
			return nil
		},
	}
}

func createRemoveAccessListCommand(listType ListType) *cli.Command {
	return &cli.Command{
		Name:      "remove",
		Usage:     fmt.Sprintf("remove an item from %s", listType),
		ArgsUsage: "<networkCIDR>",
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				return cli.Exit("missed networkCIDR arg", 10)
			}

			ctx, cancel := context.WithTimeout(context.Background(), clientTimeout)
			defer cancel()
			client := newClient(ctx, c.String("server"))
			req := &protectorpb.RemoveAccessListRequest{NetworkCIDR: c.Args().First()}

			var result *protectorpb.RemoveAccessListResponse
			var err error
			switch listType {
			case Black:
				result, err = client.RemoveBlackListItem(ctx, req)
			case White:
				result, err = client.RemoveWhiteListItem(ctx, req)
			}

			if err != nil {
				return cli.Exit(err, 10)
			}
			log.Println("Done: ", result)
			return nil
		},
	}
}

func createShowAccessListItemsCommand(listType ListType) *cli.Command {
	return &cli.Command{
		Name:  "show",
		Usage: fmt.Sprintf("show %s items", listType),
		Action: func(c *cli.Context) error {
			if c.NArg() != 0 {
				return cli.Exit("unknown argument", 10)
			}

			ctx, cancel := context.WithTimeout(context.Background(), clientTimeout)
			defer cancel()
			client := newClient(ctx, c.String("server"))
			req := &protectorpb.GetAccessListItemsRequest{}

			var result *protectorpb.GetAccessListItemsResponse
			var err error
			switch listType {
			case Black:
				result, err = client.GetBlackListItems(ctx, req)
			case White:
				result, err = client.GetWhiteListItems(ctx, req)
			}

			if err != nil {
				return cli.Exit(err, 10)
			}

			fmt.Printf("%s items:\n", listType)
			for _, item := range result.Items {
				fmt.Printf("- %s\n", item)
			}
			return nil
		},
	}
}

func createResetLoginLimitCommand() *cli.Command {
	return &cli.Command{
		Name:      "login",
		Usage:     "reset login limit",
		ArgsUsage: "<login>",
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				return cli.Exit("missed LOGIN arg", 10)
			}

			ctx, cancel := context.WithTimeout(context.Background(), clientTimeout)
			defer cancel()
			client := newClient(ctx, c.String("server"))
			req := &protectorpb.ResetLoginLimitRequest{Login: c.Args().First()}

			result, err := client.ResetLogin(ctx, req)

			if err != nil {
				return cli.Exit(err, 9)
			}
			log.Println("Done: ", result)
			return nil
		},
	}
}

func createResetIPLimitCommand() *cli.Command {
	return &cli.Command{
		Name:      "ip",
		Usage:     "reset ip limit",
		ArgsUsage: "<ip>",
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				return cli.Exit("missed IP arg", 10)
			}

			ctx, cancel := context.WithTimeout(context.Background(), clientTimeout)
			defer cancel()
			client := newClient(ctx, c.String("server"))
			req := &protectorpb.ResetIPLimitRequest{Ip: c.Args().First()}

			result, err := client.ResetIP(ctx, req)

			if err != nil {
				return cli.Exit(err, 9)
			}
			log.Println("Done: ", result)
			return nil
		},
	}
}

func newClient(ctx context.Context, server string) protectorpb.BruteforceProtectorServiceClient {
	conn, err := grpc.DialContext(ctx, server, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Cannot connect to server %s: %v", server, err)
	}

	return protectorpb.NewBruteforceProtectorServiceClient(conn)
}
