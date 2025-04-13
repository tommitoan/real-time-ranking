package handler

import (
	"encoding/json"
	"net/http"
	"real-time-ranking/internal/models"
)

func (s *Server) PostInteractions(w http.ResponseWriter, req *http.Request) {
	var r models.InteractionRequest
	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := s.s.PostInteractions(req.Context(), r)
	if err != nil {
		http.Error(w, "Failed to update score", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Score updated successfully"))
}
