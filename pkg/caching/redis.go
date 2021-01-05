package caching

import (
	"log"
	"math/big"

	"github.com/garyburd/redigo/redis"
)

// Cache do caching
type Cache struct {
	redisClient redis.Conn
}

// SetBigInt set bigint
func (cache *Cache) SetBigInt(key uint32, value *big.Int) ([]byte, error) {
	return redis.Bytes(cache.redisClient.Do("SET", key, value.Bytes()))
}

// GetBigInt get bigint
func (cache *Cache) GetBigInt(key uint32) (*big.Int, error) {
	rv, err := redis.Bytes(cache.redisClient.Do("GET", key))
	if err != nil {
		return new(big.Int), err
	}

	cachedVal := new(big.Int).SetBytes(rv)

	return cachedVal, nil
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
