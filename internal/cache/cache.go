package cache

import (
	"context"
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

func Init(addr string) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("failed to connect to cache:", err)
	}

	_c = client
	log.Println("Connected to cache")
}

func GetClient() *redis.Client {
	client, ok := _c.(*redis.Client)
	if !ok {
		return nil
	}
	return client
}

func Close() {
	err := _c.Close()
	if err != nil {
		log.Fatalln("Failed to close redis:", err)
		return
	}
	log.Println("Closed cache")
}
