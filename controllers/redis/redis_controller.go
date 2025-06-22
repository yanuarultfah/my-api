package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func RedisSet(c *gin.Context) {
	key := "key-1-set"
	data := "this is a test data"
	err := rdb.Set(ctx, key, data, 0).Err()
	if err != nil {
		fmt.Printf("unable to SET data. error: %v", err)
		return
	}
	log.Println("set operation success")
}

func RedisGet(c *gin.Context) {
	key := "key-1-set"
	data, err := rdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return
	}
	log.Println("get operation success")
	c.String(200, "Data from Redis: %s", data)
}
