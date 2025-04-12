package cache

import (
	"github.com/redis/go-redis/v9"
	"io"
	"log"
)

var _c Redis

type (
	Redis interface {
		redis.Cmdable
		io.Closer
	}
)

func Init(c Redis) {
	if c == nil {
		log.Fatal("Redis client is nil")
	}
	_c = c
	log.Println("Connected to Redis")
}

func GetClient() *redis.Client {
	client, ok := _c.(*redis.Client)
	if !ok {
		return nil
	}
	return client
}

func Close() { _c.Close() }
