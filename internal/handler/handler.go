package handler

import (
	"encoding/json"
	"net/http"
	"real-time-ranking/internal/models"
	"real-time-ranking/internal/service"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) PostInteractions(w http.ResponseWriter, req *http.Request) {
	var r models.InteractionRequest
	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := service.PostInteractions(req.Context(), r)
	if err != nil {
		http.Error(w, "Failed to update score", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Score updated successfully"))
}

func (s *Server) GetRankings(w http.ResponseWriter, req *http.Request, params GetRankingsParams) {
	data, err := service.GetGlobalRankings(req.Context(), val(params.Limit))
	if err != nil {
		http.Error(w, "Failed to fetch global rankings", http.StatusInternalServerError)
		return
	}

	var res []VideoRank
	for _, d := range data {
		res = append(res, convertToVideoRank(d))
	}

	writeJSONResponse(w, &RankingsResponse{Rankings: &res})
}

func (s *Server) GetRankingsUserUserID(w http.ResponseWriter, req *http.Request, userID string, params GetRankingsUserUserIDParams) {
	data, err := service.GetUserRankings(req.Context(), val(params.Limit), userID)
	if err != nil {
		http.Error(w, "Failed to fetch user rankings", http.StatusInternalServerError)
		return
	}

	var res []VideoRank
	for _, d := range data {
		res = append(res, convertToVideoRank(d))
	}

	writeJSONResponse(w, &RankingsResponse{Rankings: &res})
}
