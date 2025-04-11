package server

import (
	"encoding/json"
	"net/http"
	"real-time-ranking/internal/model"
	"real-time-ranking/internal/redisdb"
	"strconv"
)

// HandleGetRankings godoc
// @Summary Get top-ranked videos
// @Description Returns global top videos based on score
// @Produce  json
// @Success 200 {array} model.VideoRanking
// @Router /rankings [get]
func (s *Server) HandleGetRankings(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	limit := int64(10)

	if limitStr != "" {
		if parsed, err := strconv.ParseInt(limitStr, 10, 64); err == nil {
			limit = parsed
		}
	}

	key := "ranking:global"
	ctx := redisdb.GetContext()
	client := redisdb.GetClient()

	results, err := client.ZRevRangeWithScores(ctx, key, 0, limit-1).Result()
	if err != nil {
		http.Error(w, "Failed to fetch global ranking", http.StatusInternalServerError)
		return
	}

	var rankings []model.VideoRanking
	for _, r := range results {
		rankings = append(rankings, model.VideoRanking{
			VideoID: r.Member.(string),
			Score:   r.Score,
		})
	}

	_ = json.NewEncoder(w).Encode(rankings)
}
