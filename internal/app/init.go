package app

import (
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
	cache.Init(cfg.ConnectionString)
}
