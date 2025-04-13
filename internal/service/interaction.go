package service

import (
	"context"
	"fmt"
	"log"
	"real-time-ranking/internal/cache"
	"real-time-ranking/internal/models"
)

func (s *S) PostInteractions(ctx context.Context, req models.InteractionRequest) error {
	vidScore := CalculateVideoScore(req)
	key := "ranking:global"

	client := cache.GetClient()
	if err := client.ZIncrBy(ctx, key, vidScore, req.VideoID).Err(); err != nil {
		log.Println("Redis update failed (global):", err)
		return err
	}

	userKey := fmt.Sprintf("ranking:user:%s", req.UserID)
	if err := client.ZIncrBy(ctx, userKey, vidScore, req.VideoID).Err(); err != nil {
		log.Println("Redis update failed (user):", err)
		return err
	}

	log.Printf("Updated score for video %s by %.2f", req.VideoID, vidScore)
	return nil
}
