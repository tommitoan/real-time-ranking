package handler

import (
	"encoding/json"
	"net/http"
	"real-time-ranking/internal/models"
)

func writeJSONResponse(w http.ResponseWriter, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func convertToVideoRank(vr models.VideoRanking) VideoRank {
	return VideoRank{
		VideoId: ptr(vr.VideoID),
		Score:   ptr(float32(vr.Score)),
	}
}

func ptr[T any](v T) *T {
	return &v
}
