package redisdb

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client
var ctx = context.Background()

func InitRedis() {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}

	RDB = redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := RDB.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("✅ Connected to Redis:", addr)
}

func GetClient() *redis.Client {
	return RDB
}

func GetContext() context.Context {
	return ctx
}
