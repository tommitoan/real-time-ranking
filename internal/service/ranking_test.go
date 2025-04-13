package service_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"

	"real-time-ranking/internal/cache"
	"real-time-ranking/internal/service"
)

func TestGetGlobalRankings_Integration(t *testing.T) {
	client := cache.GetClient()
	key := "ranking:global"

	client.ZIncrBy(ctx, key, 10, "video-a")
	client.ZIncrBy(ctx, key, 20, "video-b")

	result, err := service.GetGlobalRankings(ctx, 10)
	require.NoError(t, err)
	require.Len(t, result, 2)
	require.Equal(t, "video-b", result[0].VideoID)

	testRedis.FlushAll()
}

func TestGetUserRankings_Integration(t *testing.T) {
	client := cache.GetClient()
	userID := "test-user"
	key := fmt.Sprintf("ranking:user:%s", userID)

	client.ZIncrBy(ctx, key, 10, "video-d")
	client.ZIncrBy(ctx, key, 20, "video-e")
	client.ZIncrBy(ctx, key, 30, "video-f")

	result, err := service.GetUserRankings(ctx, 10, userID)
	require.NoError(t, err)
	require.Len(t, result, 3)
	require.Equal(t, "video-f", result[0].VideoID)

	testRedis.FlushAll()
}
