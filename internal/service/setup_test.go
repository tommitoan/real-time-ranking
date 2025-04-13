package service_test

import (
	"context"
	"os"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"real-time-ranking/internal/cache"
)

var (
	testRedis *miniredis.Miniredis
	ctx       = context.Background()
)

func TestMain(m *testing.M) {
	var err error
	testRedis, err = miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer testRedis.Close()

	cache.Init(testRedis.Addr())

	code := m.Run()
	os.Exit(code)
}
