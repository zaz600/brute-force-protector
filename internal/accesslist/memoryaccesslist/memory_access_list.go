package memoryaccesslist

import (
	"fmt"
	"net"
	"sync"
)

type ListValue struct {
	IP    net.IP
	IPNet *net.IPNet
}

// MemoryAccessList реализация списка доступа с хранением элементов в памяти.
type MemoryAccessList struct {
	*sync.RWMutex
	db map[string]ListValue
}

// Add добавление подсети в список доступа.
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

// Remove удаление подсети из списка доступа.
func (m *MemoryAccessList) Remove(networkCIDR string) error {
	m.Lock()
	defer m.Unlock()
	delete(m.db, networkCIDR)
	return nil
}

// Len количество элементов в списке доступа.
func (m *MemoryAccessList) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.db)
}

// Exists проверка, что подсеть есть в списке доступа.
func (m *MemoryAccessList) Exists(networkCIDR string) bool {
	m.RLock()
	defer m.RUnlock()
	if _, ok := m.db[networkCIDR]; ok {
		return true
	}
	return false
}

// IsInList проверяет, что IP входит в одну из подсетей списка доступа.
func (m *MemoryAccessList) IsInList(ip string) bool {
	m.RLock()
	defer m.RUnlock()
	return m.isInList(net.ParseIP(ip))
}

func (m *MemoryAccessList) isInList(ip net.IP) bool {
	if ip == nil {
		return false
	}
	found := false
	for _, val := range m.db {
		if ok := val.IPNet.Contains(ip); ok {
			found = true
			break
		}
	}
	return found
}

// GetAll возвращает все элементы списка доступа.
func (m *MemoryAccessList) GetAll() []string {
	m.RLock()
	defer m.RUnlock()
	result := make([]string, 0, len(m.db))
	for k := range m.db {
		result = append(result, k)
	}
	return result
}

// NewMemoryAccessList создает список доступа с хранением элементов в памяти.
func NewMemoryAccessList() *MemoryAccessList {
	return &MemoryAccessList{
		RWMutex: &sync.RWMutex{},
		db:      make(map[string]ListValue),
	}
}
