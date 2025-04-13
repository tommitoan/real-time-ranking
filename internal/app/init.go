package app

import (
	"context"
	"real-time-ranking/internal/cache"
	"real-time-ranking/internal/config"
	"real-time-ranking/internal/db"
	"sync"
)

var once sync.Once

func Init(cfg *config.Service) {
	once.Do(func() {
		InitRedis(cfg.Redis)
		InitDatabase(cfg.Database)
	})
}

func InitRedis(cfg config.Redis) {
	cache.Init(cfg.ConnectionString)
}

func InitDatabase(cfg config.Database) {
	ctx := context.Background()
	db.InitDataSource(ctx, cfg)
}

func Close() {
	cache.Close()
	db.Close()
}
