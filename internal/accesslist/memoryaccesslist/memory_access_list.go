package memoryaccesslist

import (
	"fmt"
	"net"
	"sync"

	"github.com/zaz600/brute-force-protector/internal/accesslist"
)

type ListValue struct {
	IP    net.IP
	IPNet *net.IPNet
}

type MemoryAccessList struct {
	*sync.RWMutex
	db map[string]ListValue
}

func (m *MemoryAccessList) Add(networkCIDR string) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.db[networkCIDR]; ok {
		return nil
	}

	ipv4Addr, ipv4Net, err := net.ParseCIDR(networkCIDR)
	if err != nil {
		return fmt.Errorf("can't add value to list: %w", err)
	}

	m.db[networkCIDR] = ListValue{
		IP:    ipv4Addr,
		IPNet: ipv4Net,
	}
	return nil
}

func (m *MemoryAccessList) Remove(networkCIDR string) {
	m.Lock()
	defer m.Unlock()
	delete(m.db, networkCIDR)
}

func (m *MemoryAccessList) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.db)
}

func (m *MemoryAccessList) Clear() {
	m.Lock()
	defer m.Unlock()
	m.db = make(map[string]ListValue)
}

func (m *MemoryAccessList) Exists(networkCIDR string) bool {
	m.RLock()
	defer m.RUnlock()
	if _, ok := m.db[networkCIDR]; ok {
		return true
	}
	return false
}

func (m *MemoryAccessList) IsInList(ip string) bool {
	m.RLock()
	defer m.RUnlock()

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	found := false
	for _, val := range m.db {
		if ok := val.IPNet.Contains(parsedIP); ok {
			found = true
			break
		}
	}

	return found
}

func NewMemoryAccessList() accesslist.AccessList {
	return &MemoryAccessList{
		RWMutex: &sync.RWMutex{},
		db:      make(map[string]ListValue),
	}
}
