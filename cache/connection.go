package cache

import (
	"github.com/redis/go-redis/v9"
	"time"
)

// redisCache type is defined with three fields - host which is the Redis server host address, db which is the Redis
// database to be used, and expiryDuration which is the default expiration time for keys.
type redisCache struct {
	host           string
	db             int
	expiryDuration time.Duration
}

// newRedisCache function creates and returns a pointer to a new redisCache object with the given host, db,
// and exp values.
func newRedisCache(host string, db int, exp time.Duration) *redisCache {
	return &redisCache{
		host:           host,
		db:             db,
		expiryDuration: exp,
	}
}

// getClient method returns a new Redis client using the redis.Options struct that is initialized with the values
// from the redisCache object.
func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

// NewClient function creates a new redisCache object using the given host, db, and exp values, and then returns a new
// Redis client by calling the getClient method on the redisCache object.
func NewClient(host string, db int, exp time.Duration) redis.Client {
	cache := newRedisCache(host, db, exp)
	client := cache.getClient()
	return *client
}
