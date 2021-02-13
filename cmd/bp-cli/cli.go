package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

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
		},
	}
	return app
}

func createAddAccessListCommand(listType ListType) *cli.Command {
	return &cli.Command{
		Name:      "add",
		Usage:     fmt.Sprintf("add an item to %s", listType),
		ArgsUsage: "networkCIDR",
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				return cli.Exit("missed networkCIDR arg", 10)
			}

			service := bpService{host: c.String("host")}
			if err := service.addAccessList(c.Args().First(), listType); err != nil {
				return cli.Exit(err, 9)
			}
			return nil
		},
	}
}

func createRemoveAccessListCommand(listType ListType) *cli.Command {
	return &cli.Command{
		Name:  "remove",
		Usage: "remove an item to whitelist",
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				return cli.Exit("missed networkCIDR arg", 10)
			}

			service := bpService{host: c.String("host")}
			if err := service.removeAccessList(c.Args().First(), listType); err != nil {
				return cli.Exit(err, 9)
			}
			return nil
		},
	}
}

func createShowAccessListItemsCommand(listType ListType) *cli.Command {
	return &cli.Command{
		Name:  "show",
		Usage: "show whitelist items",
		Action: func(c *cli.Context) error {
			if c.NArg() != 0 {
				return cli.Exit("unknown argument", 10)
			}

			service := bpService{host: c.String("host")}
			items, err := service.getAccessListItems(listType)
			if err != nil {
				return cli.Exit(err, 9)
			}

			fmt.Printf("%s items:\n", listType)
			for _, item := range items {
				fmt.Printf("- %s\n", item)
			}
			return nil
		},
	}
}
