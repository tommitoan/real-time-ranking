package service

import (
	"context"
	"log"
	"real-time-ranking/internal/cache"
	"real-time-ranking/internal/models"
)

func GetGlobalRankings(ctx context.Context, limit int64) ([]models.VideoRanking, error) {
	client := cache.GetClient()
	results, err := client.ZRevRangeWithScores(ctx, "ranking:global", 0, limit-1).Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var rankings []models.VideoRanking
	for _, r := range results {
		rankings = append(rankings, models.VideoRanking{
			VideoID: r.Member.(string),
			Score:   r.Score,
		})
	}
	return rankings, nil
}

func GetUserRankings(ctx context.Context, limit int64, userID string) ([]models.VideoRanking, error) {
	client := cache.GetClient()
	results, err := client.ZRevRangeWithScores(ctx, "ranking:user:"+userID, 0, limit-1).Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var rankings []models.VideoRanking
	for _, r := range results {
		rankings = append(rankings, models.VideoRanking{
			VideoID: r.Member.(string),
			Score:   r.Score,
		})
	}
	return rankings, nil
}
