package redisaccesslist

import (
	"context"
	"log"
	"net"

	"github.com/go-redis/redis/v8"
)

// RedisAccessList список доступа с хранением элементов в redis
type RedisAccessList struct {
	redisClient *redis.Client
	hashName    string
}

// NewRedisAccessList возвращает список доступа с хранением элементов в redis
func NewRedisAccessList(hashName string, redisClient *redis.Client) *RedisAccessList {
	return &RedisAccessList{
		redisClient: redisClient,
		hashName:    hashName,
	}
}

// Add добавление подсети в список доступа
func (r *RedisAccessList) Add(networkCIDR string) error {
	err := r.redisClient.HSet(context.Background(), r.hashName, networkCIDR, "").Err()
	if err != nil {
		return err
	}
	return nil
}

// Remove удаление подсети из списка доступа
func (r *RedisAccessList) Remove(networkCIDR string) error {
	err := r.redisClient.HDel(context.Background(), r.hashName, networkCIDR).Err()
	if err != nil {
		return err
	}
	return nil
}

// Exists проверка, что подсеть есть в списке доступа
func (r *RedisAccessList) Exists(networkCIDR string) bool {
	ok, err := r.redisClient.HExists(context.Background(), r.hashName, networkCIDR).Result()
	if err != nil || !ok {
		return false
	}
	return true
}

// IsInList проверяет, что IP входит в одну из подсетей списка доступа
func (r *RedisAccessList) IsInList(ip string) bool {
	items := r.GetAll()
	return r.isInList(net.ParseIP(ip), items)
}

func (r *RedisAccessList) isInList(ip net.IP, items []string) bool {
	if ip == nil {
		return false
	}
	for _, network := range items {
		// TODO кешировать
		_, ipv4Net, err := net.ParseCIDR(network)
		if err != nil {
			// return fmt.Errorf("can't add value to list: %w", err)
			continue
		}

		if ok := ipv4Net.Contains(ip); ok {
			return true
		}
	}
	return false
}

// Len количество элементов в списке доступа
func (r *RedisAccessList) Len() int {
	size, err := r.redisClient.HLen(context.Background(), r.hashName).Result()
	if err != nil {
		log.Println(err)
	}
	return int(size)
}

// GetAll возвращает все элементы списка доступа
func (r *RedisAccessList) GetAll() []string {
	items, err := r.redisClient.HGetAll(context.Background(), r.hashName).Result()

	if err != nil {
		log.Println(err)
		return []string{}
	}

	var result = make([]string, 0, len(items))
	for key := range items {
		result = append(result, key)
	}
	return result
}
