package caching

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

// Cache do caching
type Cache struct {
	redisClient redis.Conn
}

// SetUint set Uint
func (cache *Cache) SetUint(key, value uint64) (uint64, error) {
	return redis.Uint64(cache.redisClient.Do("SET", key, value))
}

// GetUint get uint
func (cache *Cache) GetUint(key uint64) (uint64, error) {
	return redis.Uint64(cache.redisClient.Do("GET", key))
}

// NewCacheConnection get cache instance
func NewCacheConnection(address string) *Cache {
	redisClient, err := redis.Dial("tcp", address)
	if err != nil {
		log.Fatalf("failed to dial redis: %v", err)
	}

	cache := &Cache{
		redisClient: redisClient,
	}

	return cache
}
