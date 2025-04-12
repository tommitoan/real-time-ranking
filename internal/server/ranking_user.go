package server

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"real-time-ranking/internal/cache"
	"real-time-ranking/internal/model"
)

// HandleGetUserRankings godoc
// @Summary Get user-personalized rankings
// @Description Returns top videos for a specific user
// @Param user_id path string true "User ID"
// @Produce  json
// @Success 200 {array} model.VideoRanking
// @Router /rankings/user/{user_id} [get]
func (s *Server) HandleGetUserRankings(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	if userID == "" {
		http.Error(w, "Missing userID", http.StatusBadRequest)
		return
	}

	key := "ranking:user:" + userID
	ctx := context.Background()
	client := cache.GetClient()

	// Get top 10 videos
	results, err := client.ZRevRangeWithScores(ctx, key, 0, 9).Result()
	if err != nil {
		http.Error(w, "Failed to fetch user ranking", http.StatusInternalServerError)
		return
	}

	var response []model.VideoRanking
	for _, r := range results {
		response = append(response, model.VideoRanking{
			VideoID: r.Member.(string),
			Score:   r.Score,
		})
	}

	json.NewEncoder(w).Encode(response)
}
