package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"

	protectorpb "github.com/zaz600/brute-force-protector/api"
)

type ListType string

const (
	Black ListType = "BlackList"
	White ListType = "WhiteList"
)

func main() {
	os.Exit(CLI(os.Args))
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

func (s bpService) resetLoginLimit(login string) error {
	client, err := newBpClient(s.host)
	if err != nil {
		return fmt.Errorf("error connect to host %s: %w", s.host, err)
	}
	defer client.conn.Close()

	req := &protectorpb.ResetLoginLimitRequest{Login: login}
	_, err = client.rpcClient.ResetLogin(context.TODO(), req)
	if err != nil {
		return fmt.Errorf("error reset login limit: %w", err)
	}
	return nil
}

func (s bpService) resetIPLimit(ip string) error {
	client, err := newBpClient(s.host)
	if err != nil {
		return fmt.Errorf("error connect to host %s: %w", s.host, err)
	}
	defer client.conn.Close()

	req := &protectorpb.ResetIPLimitRequest{Ip: ip}
	_, err = client.rpcClient.ResetIP(context.TODO(), req)
	if err != nil {
		return fmt.Errorf("error reset ip limit: %w", err)
	}
	return nil
}
