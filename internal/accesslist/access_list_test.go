package accesslist_test

import (
	"github.com/stretchr/testify/require"
	accesslist2 "github.com/zaz600/brute-force-protector/internal/accesslist"
	"testing"
)

const defaultCIDR = "192.168.1.1/24"

func memoryAccessListWithData(t *testing.T) accesslist2.AccessList {
	t.Helper()
	list := accesslist2.NewMemoryAccessList()
	err := list.Add(defaultCIDR)
	require.NoError(t, err)
	return list
}

func TestMemoryAccessList_Add(t *testing.T) {
	type test struct {
		name string
		cidr string
		err  bool
	}

	for _, tst := range [...]test{
		{
			name: "valid CIDR",
			cidr: "192.0.2.1/24",
			err:  false,
		},
		{
			name: "valid CIDR small range",
			cidr: "192.168.0.1/28",
			err:  false,
		},
		{
			name: "not a CIDR",
			cidr: "a4bc2d5e",
			err:  true,
		},
		{
			name: "invalid CIDR",
			cidr: "192.0.2.256/24",
			err:  true,
		},
		{
			name: "empty CIDR",
			cidr: "",
			err:  true,
		},
	} {
		t.Run(tst.name, func(t *testing.T) {
			list := memoryAccessListWithData(t)
			err := list.Add(tst.cidr)
			if tst.err {
				require.Error(t, err)
				require.Equal(t, list.Len(), 1)
			} else {
				require.NoError(t, err)
				require.Equal(t, list.Len(), 2)
			}
		})
	}
}

func TestMemoryAccessList_IsInList(t *testing.T) {
	type test struct {
		name   string
		cidr   string
		ip     string
		inList bool
	}

	for _, tst := range [...]test{
		{
			name:   "ip in the list",
			cidr:   "192.0.2.1/24",
			ip:     "192.0.2.2",
			inList: true,
		},
		{
			name:   "ip in the small list",
			cidr:   "192.168.0.1/28",
			ip:     "192.168.0.5",
			inList: true,
		},
		{
			name:   "ip in the small list",
			cidr:   "192.168.0.1/28", //192.168.0.1 - 192.168.0.14, Broadcast: 192.168.0.15
			ip:     "192.168.0.15",
			inList: true,
		},
		{
			name:   "ip not in the list",
			cidr:   "192.0.2.1/24",
			ip:     "1.1.1.1",
			inList: false,
		},
		{
			name:   "ip in the small list",
			cidr:   "192.168.0.1/28", //192.168.0.1 - 192.168.0.14, Broadcast: 192.168.0.15
			ip:     "192.168.0.16",
			inList: false,
		},
		{
			name:   "invalid ip",
			cidr:   "192.0.2.1/24",
			ip:     "1.1.1.256",
			inList: false,
		},
		{
			name:   "empty ip",
			cidr:   "192.0.2.1/24",
			ip:     "",
			inList: false,
		},
	} {
		t.Run(tst.name, func(t *testing.T) {
			list := memoryAccessListWithData(t)
			err := list.Add(tst.cidr)
			require.NoError(t, err)

			actual := list.IsInList(tst.ip)
			require.Equal(t, tst.inList, actual)

			// find old record
			actual = list.IsInList("192.168.1.2")
			require.Equal(t, true, actual)
		})
	}
}
