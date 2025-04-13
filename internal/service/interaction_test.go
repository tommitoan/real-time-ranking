package service_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"real-time-ranking/internal/cache"
	"real-time-ranking/internal/models"
	"real-time-ranking/internal/service"
)

func TestPostInteractions_Integration(t *testing.T) {
	client := cache.GetClient()

	userID := "user-123"
	videoID := "video-xyz"

	req := models.InteractionRequest{
		UserID:    userID,
		VideoID:   videoID,
		Views:     1,
		Likes:     1,
		Shares:    1,
		Comments:  1,
		WatchTime: 10,
	}

	globalKey := "ranking:global"
	userKey := fmt.Sprintf("ranking:user:%s", userID)

	err := service.PostInteractions(ctx, req)
	require.NoError(t, err)

	globalScore, err := client.ZScore(ctx, globalKey, videoID).Result()
	require.NoError(t, err)
	require.NotZero(t, globalScore)

	userScore, err := client.ZScore(ctx, userKey, videoID).Result()
	require.NoError(t, err)
	require.Equal(t, globalScore, userScore)

	testRedis.FlushAll()
}
