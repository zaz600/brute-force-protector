package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-redis/redis/v8"
	"github.com/urfave/cli/v2"

	"github.com/zaz600/brute-force-protector/internal/accesslist/redisaccesslist"
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

func createApp() *cli.App { // nolint
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
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			opts := []bp.ProtectorOption{
				bp.WithContext(ctx),
				bp.WithLoginLimit(c.Int64("n")),
				bp.WithPasswordLimit(c.Int64("m")),
				bp.WithIPLimit(c.Int64("k")),
			}

			if c.String("redis") != "" {
				redisClient, err := getRedisClient(ctx, c.String("redis"))
				if err != nil {
					return cli.Exit(fmt.Sprintf("can't connect to redis: %v", err), 1)
				}

				log.Println("use redis as access lists storage")
				opts = append(opts, bp.WithBlackList(redisaccesslist.NewRedisAccessList("blacklist", redisClient)))
				opts = append(opts, bp.WithWhiteList(redisaccesslist.NewRedisAccessList("whitelist", redisClient)))
			}

			bpServer := grpc.NewBPServer(bp.NewBruteForceProtector(opts...))
			go func() {
				termCh := make(chan os.Signal, 1)
				signal.Notify(termCh, os.Interrupt, syscall.SIGINT)
				<-termCh
				log.Println("Shutdown...")
				bpServer.GracefulStop()
			}()
			err := bpServer.ListenAndServe(c.String("listen"))
			if err != nil {
				return cli.Exit(fmt.Sprintf("Can't start server: %v", err), 1)
			}
			return nil
		},
	}
	return app
}

func getRedisClient(ctx context.Context, redisHost string) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "",
		DB:       0,
	})
	redisClient = redisClient.WithContext(ctx)

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("can't connect to redis: %w", err)
	}

	return redisClient, nil
}
