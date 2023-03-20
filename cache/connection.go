package cache

import (
	"github.com/redis/go-redis/v9"
	"time"
)

type redisCache struct {
	host           string
	db             int
	expiryDuration time.Duration
}

func newRedisCache(host string, db int, exp time.Duration) *redisCache {
	return &redisCache{
		host:           host,
		db:             db,
		expiryDuration: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func NewClient(host string, db int, exp time.Duration) redis.Client {
	cache := newRedisCache(host, db, exp)
	client := cache.getClient()
	return *client
}
