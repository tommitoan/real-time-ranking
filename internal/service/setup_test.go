package service_test

import (
	"context"
	"os"
	"real-time-ranking/internal/db"
	"real-time-ranking/internal/service"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"real-time-ranking/internal/cache"
)

var (
	testRedis *miniredis.Miniredis
	ctx       = context.Background()
	s         service.Service
)

func TestMain(m *testing.M) {
	var err error
	testRedis, err = miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer testRedis.Close()

	cache.Init(testRedis.Addr())
	s = service.NewService(db.TestDataSource())

	code := m.Run()
	os.Exit(code)
}
