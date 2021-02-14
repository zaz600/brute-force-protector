package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	protectorpb "github.com/zaz600/brute-force-protector/internal/grpc/api"
)

const ITEST_SERVER = "localhost:50051"

func newClient() (protectorpb.BruteforceProtectorServiceClient, error) {
	conn, err := grpc.Dial(ITEST_SERVER, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	return protectorpb.NewBruteforceProtectorServiceClient(conn), nil
}

func TestBruteforceProtectorServer_Connect(t *testing.T) {
	_, err := newClient()
	require.NoError(t, err)
}

func TestBruteforceProtectorServer_AccessList(t *testing.T) {
	type test struct {
		name      string
		blackList bool
	}
	for _, tt := range [...]test{
		{
			name:      "black list",
			blackList: true,
		},
		{
			name:      "white list",
			blackList: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			client, err := newClient()
			require.NoError(t, err)
			ip := faker.IPv4()
			net := fmt.Sprintf("%s/24", ip)

			// добавляем в список
			resp, err := addAccessList(net, tt.blackList, client)
			require.NoError(t, err)
			require.True(t, resp.Result)

			// проверяем, что в списке есть
			respItems, err := getAccessListItems(tt.blackList, client)
			require.NoError(t, err)
			require.Contains(t, respItems.Items, net)

			// удаляем из списка
			_, err = removeAccessListItem(net, tt.blackList, client)
			require.NoError(t, err)

			// проверяем, что больше не в списке
			respItems, err = getAccessListItems(tt.blackList, client)
			require.NoError(t, err)
			require.NotContains(t, respItems.Items, net)
		})
	}
}

func TestBruteforceProtectorServer_Verify(t *testing.T) {
	type test struct {
		name          string
		count         int
		limitType     string
		expectAllowed bool
	}
	for _, tt := range [...]test{
		{
			name:          "not limit reached",
			count:         10,
			limitType:     "login",
			expectAllowed: true,
		},
		{
			name:          "limit reached",
			count:         11,
			limitType:     "login",
			expectAllowed: false,
		},
		{
			name:          "not limit reached",
			count:         100,
			limitType:     "password",
			expectAllowed: true,
		},
		{
			name:          "limit reached",
			count:         101,
			limitType:     "password",
			expectAllowed: false,
		},
		{
			name:          "not limit reached",
			count:         1000,
			limitType:     "ip",
			expectAllowed: true,
		},
		{
			name:          "limit reached",
			count:         1001,
			limitType:     "ip",
			expectAllowed: false,
		},
	} {
		t.Run(fmt.Sprintf("%s: %s", tt.limitType, tt.name), func(t *testing.T) {
			client, err := newClient()
			require.NoError(t, err)

			login := faker.Username()
			password := faker.Password()
			ip := faker.IPv4()
			for i := 0; i < tt.count; i++ {
				switch tt.limitType {
				case "login":
					password = faker.Password()
					ip = faker.IPv4()
				case "password":
					login = faker.Username()
					ip = faker.IPv4()
				case "ip":
					password = faker.Password()
					login = faker.Username()
				}

				resp, err := verify(login, password, ip, client)
				require.NoError(t, err)

				if i == tt.count-1 {
					require.Equal(t, tt.expectAllowed, resp.Ok)
				} else {
					require.True(t, resp.Ok)
				}
			}
		})
	}
}
func TestBruteforceProtectorServer_ResetLimit(t *testing.T) {
	type test struct {
		name      string
		count     int
		limitType string
	}
	for _, tt := range [...]test{
		{
			name:      "reset limit",
			count:     11,
			limitType: "login",
		},
		{
			name:      "reset limit",
			count:     1001,
			limitType: "ip",
		},
	} {
		t.Run(fmt.Sprintf("%s: %s", tt.limitType, tt.name), func(t *testing.T) {
			client, err := newClient()
			require.NoError(t, err)

			login := faker.Username()
			password := faker.Password()
			ip := faker.IPv4()

			for i := 0; i < tt.count; i++ {
				switch tt.limitType {
				case "login":
					password = faker.Password()
					ip = faker.IPv4()
				case "ip":
					password = faker.Password()
					login = faker.Username()
				}

				resp, err := verify(login, password, ip, client)
				require.NoError(t, err)

				if i == tt.count-1 {
					require.False(t, resp.Ok)
				} else {
					require.True(t, resp.Ok)
				}
			}

			// сброс
			if tt.limitType == "login" {
				_, err := resetLimit(tt.limitType, login, client)
				require.NoError(t, err)
			} else {
				_, err := resetLimit(tt.limitType, ip, client)
				require.NoError(t, err)
			}
			// повтор
			resp, err := verify(login, password, ip, client)
			require.NoError(t, err)
			require.True(t, resp.Ok)
		})
	}
}

func addAccessList(net string, blacklist bool, client protectorpb.BruteforceProtectorServiceClient) (*protectorpb.AddAccessListResponse, error) {
	req := &protectorpb.AddAccessListRequest{NetworkCIDR: net}
	if blacklist {
		return client.AddBlackListItem(context.Background(), req)
	}
	return client.AddWhiteListItem(context.Background(), req)
}

func getAccessListItems(blacklist bool, client protectorpb.BruteforceProtectorServiceClient) (*protectorpb.GetAccessListItemsResponse, error) {
	req := &protectorpb.GetAccessListItemsRequest{}
	if blacklist {
		return client.GetBlackListItems(context.Background(), req)
	}
	return client.GetWhiteListItems(context.Background(), req)
}

func removeAccessListItem(net string, blacklist bool, client protectorpb.BruteforceProtectorServiceClient) (*protectorpb.RemoveAccessListResponse, error) {
	req := &protectorpb.RemoveAccessListRequest{NetworkCIDR: net}
	if blacklist {
		return client.RemoveBlackListItem(context.Background(), req)
	}
	return client.RemoveWhiteListItem(context.Background(), req)
}

func verify(login string, password string, ip string, client protectorpb.BruteforceProtectorServiceClient) (*protectorpb.VerifyResponse, error) {
	params := &protectorpb.VerifyParams{
		Login:    login,
		Password: password,
		Ip:       ip,
	}
	return client.Verify(context.Background(), &protectorpb.VerifyRequest{
		VerifyParams: params,
	})
}

func resetLimit(limitType string, value string, client protectorpb.BruteforceProtectorServiceClient) (*protectorpb.ResetLimitResponse, error) {
	if limitType == "login" {
		return client.ResetLogin(context.Background(), &protectorpb.ResetLoginLimitRequest{
			Login: value,
		})
	} else {
		return client.ResetIP(context.Background(), &protectorpb.ResetIPLimitRequest{
			Ip: value,
		})
	}
}
