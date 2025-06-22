package redis

import "github.com/go-redis/redis"

func newRedisClient() *redis.Client {
	// Initialize the Redis client here
	// This function should set up the connection to the Redis server
	// and handle any necessary configuration.
	// Example:
	var redisHost = "localhost:6379"
	var redisPassword = ""
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       0,
	})
	return client
}
