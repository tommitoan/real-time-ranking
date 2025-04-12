package app

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"real-time-ranking/internal/cache"
	"real-time-ranking/internal/config"
	"sync"
)

var once sync.Once

func Init(cfg *config.Service) {
	once.Do(func() {
		InitRedis(cfg.Redis)
	})
}

func InitRedis(cfg config.Redis) {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.ConnectionString,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("failed to connect to redis: ", err)
	}

	log.Println("Redis ping succeeded:", pong)
	cache.Init(client)
}
