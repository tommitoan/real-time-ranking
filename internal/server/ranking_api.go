package server

import (
	"context"
	"encoding/json"
	"net/http"
	"real-time-ranking/internal/api"
	"real-time-ranking/internal/cache"
	"strconv"
)

type RankingAPI struct{}

func (r *RankingAPI) PostInteractions(w http.ResponseWriter, req *http.Request) {
	// Implement your logic for handling interactions
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Interaction processed"))
}

func (r *RankingAPI) GetRankings(w http.ResponseWriter, req *http.Request) {
	// Default limit for the number of top videos returned
	limitStr := req.URL.Query().Get("limit")
	limit := int64(10) // Default to top 10

	// If a limit is provided as a query parameter, parse it
	if limitStr != "" {
		if parsed, err := strconv.ParseInt(limitStr, 10, 64); err == nil {
			limit = parsed
		} else {
			// If limit is invalid, respond with a bad request
			http.Error(w, "Invalid limit value", http.StatusBadRequest)
			return
		}
	}

	// The key used to store global rankings in Redis
	key := "ranking:global"
	// Get context and client from the Redis DB package
	ctx := context.Background()
	client := cache.GetClient()

	// Fetch the global ranking from Redis sorted set
	results, err := client.ZRevRangeWithScores(ctx, key, 0, limit-1).Result()
	if err != nil {
		// If there's an error fetching data, return internal server error
		http.Error(w, "Failed to fetch global ranking", http.StatusInternalServerError)
		return
	}

	// Prepare the response with the rankings in the required format
	var rankings []api.VideoRank
	for _, r := range results {
		rankings = append(rankings, api.VideoRank{
			VideoId: ptr(r.Member.(string)), // Assuming r.Member is a string, using a pointer to it
			Score:   ptr(float32(r.Score)),
		})
	}

	// Prepare the response object
	resp := &api.RankingsResponse{
		Rankings: &rankings,
	}

	// Send the JSON response
	writeJSONResponse(w, resp)
}

func (r *RankingAPI) GetRankingsUserUserID(w http.ResponseWriter, req *http.Request, userID string) {
	// Implement your logic for fetching user-specific rankings
	resp := &api.UserRankingResponse{
		UserId: ptr(userID),
		Rank:   ptr(1),
	}
	writeJSONResponse(w, resp)
}

// Helper function to write JSON response
func writeJSONResponse(w http.ResponseWriter, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func ptr[T any](v T) *T {
	return &v
}
