package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	bp "github.com/zaz600/brute-force-protector/internal/bruteforceprotector"
	"github.com/zaz600/brute-force-protector/internal/grpc"
)

func main() {
	os.Exit(CLI(os.Args))
}

func CLI(args []string) int {
	app := createApp()
	if err := app.Run(args); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}
	return 0
}

func createApp() *cli.App {
	app := &cli.App{
		Name:  "bp-server",
		Usage: "bruteforce protector server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "listen",
				Aliases: []string{"l"},
				Value:   "0.0.0.0:50051",
				Usage:   "server host:port",
				EnvVars: []string{"BP_ADDR"},
			},
			&cli.Int64Flag{
				Name:    "n",
				Value:   10,
				Usage:   "logins limit per minute",
				EnvVars: []string{"BP_LOGIN_LIMIT"},
			},
			&cli.Int64Flag{
				Name:    "m",
				Value:   100,
				Usage:   "passwords limit per minute",
				EnvVars: []string{"BP_PASSWORD_LIMIT"},
			},
			&cli.Int64Flag{
				Name:    "k",
				Value:   1000,
				Usage:   "IPs limit per minute",
				EnvVars: []string{"BP_IP_LIMIT"},
			},
			&cli.StringFlag{
				Name:    "redis",
				Value:   "",
				Usage:   "use redis to store access lists",
				EnvVars: []string{"BP_REDIS_HOST"},
			},
		},
		Action: func(c *cli.Context) error {
			protector := bp.NewBruteForceProtector(
				bp.WithLoginLimit(c.Int64("n")),
				bp.WithPasswordLimit(c.Int64("m")),
				bp.WithIPLimit(c.Int64("k")),
				bp.WithRedis(c.String("redis")),
			)
			bpServer := grpc.NewBPServer(protector)
			err := bpServer.ListenAndServe(c.String("listen"))
			if err != nil {
				return cli.Exit(fmt.Sprintf("Cannot start server: %v", err), 1)
			}
			return nil
		},
	}
	return app
}
