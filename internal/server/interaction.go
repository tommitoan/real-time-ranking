package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"real-time-ranking/internal/cache"
	"real-time-ranking/internal/model"
	"real-time-ranking/internal/score"
)

// InteractionRequest represents a user interaction.
// swagger:model
type InteractionRequest struct {
	VideoID   string  `json:"video_id"`
	UserID    string  `json:"user_id"`
	Views     int     `json:"views"`
	Likes     int     `json:"likes"`
	Comments  int     `json:"comments"`
	Shares    int     `json:"shares"`
	WatchTime float64 `json:"watch_time"`
}

// HandleInteraction godoc
// @Summary Record a user interaction with a video
// @Description Updates the video score based on user interactions
// @Tags interactions
// @Accept json
// @Produce json
// @Param interaction body InteractionRequest true "Interaction Data"
// @Success 200 {string} string "Score updated successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Failed to update score"
// @Router /interactions [post]
func (s *Server) HandleInteraction(w http.ResponseWriter, r *http.Request) {
	var req model.InteractionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	vidScore := score.CalculateVideoScore(req)
	key := "ranking:global"
	ctx := context.Background()
	client := cache.GetClient()

	if err := client.ZIncrBy(ctx, key, vidScore, req.VideoID).Err(); err != nil {
		log.Println("Redis update failed (global):", err)
		http.Error(w, "Failed to update score", http.StatusInternalServerError)
		return
	}

	userKey := fmt.Sprintf("ranking:user:%s", req.UserID)
	if err := client.ZIncrBy(ctx, userKey, vidScore, req.VideoID).Err(); err != nil {
		log.Println("Redis update failed (user):", err)
		http.Error(w, "Failed to update user score", http.StatusInternalServerError)
		return
	}

	log.Printf("Updated score for video %s by %.2f", req.VideoID, vidScore)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Score updated successfully"))
}
