package handler

import (
	"encoding/json"
	"net/http"
	"real-time-ranking/internal/models"
)

func (s *Server) PatchVideosIdReactions(w http.ResponseWriter, r *http.Request, id string) {
	var req models.UpdateVideoReactionsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	v, err := s.s.UpdateVideoReactions(r.Context(), id, &req)
	if err != nil {
		http.Error(w, "Failed to update video", http.StatusInternalServerError)
		return
	}

	err = s.s.PostInteractions(r.Context(), models.InteractionRequest{
		VideoID:   v.ID,
		UserID:    v.OwnerID,
		Views:     req.Views,
		Likes:     req.Likes,
		Comments:  req.Comments,
		Shares:    req.Shares,
		WatchTime: float64(req.WatchTime),
	})
	if err != nil {
		http.Error(w, "Failed to update score", http.StatusInternalServerError)
		return
	}

	writeJSONResponse(w, &UserResponse{Id: ptr(id)})
}
