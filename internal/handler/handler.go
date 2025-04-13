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

func (s *Server) GetRankingsTop(w http.ResponseWriter, req *http.Request, limit int64) {
	data, err := service.GetGlobalRankings(req.Context(), limit)
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

func (s *Server) GetRankingsTopUserID(w http.ResponseWriter, req *http.Request, limit int64, userID string) {
	data, err := service.GetUserRankings(req.Context(), limit, userID)
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

func (s *Server) PostUsers(w http.ResponseWriter, req *http.Request) {
	var r models.CreateUserRequest
	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	id, err := service.CreateUser(req.Context(), r)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	writeJSONResponse(w, &UserResponse{Id: ptr(id), Username: ptr(r.Username)})
}
