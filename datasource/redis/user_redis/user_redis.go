package redis

import (
	"github.com/go-redis/redis/v8"
)

// type redisCache struct {
// 	host string
// 	port string
// 	exp  time.Duration
// }

func NewRedisCache() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

// func (cache redisCache) getClient() *redis.Client {
// 	return redis.NewClient(&redis.Options{
// 		host: cache.host,
// 		port: cache.port,
// 		exp:  cache.exp,
// 	})
// }
