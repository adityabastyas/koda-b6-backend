package lib

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client
var Ctx = context.Background()

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := RDB.Ping(Ctx).Result()
	if err != nil {
		panic("Redis tidak terhubung")
	}

	fmt.Println("Redis connected")
}
