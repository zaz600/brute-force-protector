package accesslist

type AccessList interface {
	Add(networkCIDR string) error
	Remove(networkCIDR string)
	Exists(networkCIDR string) bool
	IsInList(ip string) bool
	Len() int
	Clear()
	GetAll() []string
}
