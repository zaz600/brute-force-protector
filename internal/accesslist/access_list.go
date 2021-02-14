package accesslist

type AccessList interface {
	Add(networkCIDR string) error
	Remove(networkCIDR string) error
	Exists(networkCIDR string) bool
	IsInList(ip string) bool
	Len() int
	GetAll() []string
}
