package handler

import (
	"encoding/json"
	"net/http"
	"real-time-ranking/internal/models"
)

func (s *Server) PostVideos(w http.ResponseWriter, r *http.Request) {
	var req models.CreateVideoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	id, err := s.s.CreateVideo(r.Context(), &models.Video{
		Title:       req.Title,
		URL:         req.Url,
		Description: req.Description,
		Thumbnail:   req.Thumbnail,
		OwnerID:     req.OwnerId,
	})
	if err != nil {
		http.Error(w, "Failed to create video", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	writeJSONResponse(w, &VideoResponse{Id: ptr(id)})
}

func (s *Server) GetVideosId(w http.ResponseWriter, r *http.Request, id string) {
	v, err := s.s.GetVideo(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to get video", http.StatusInternalServerError)
		return
	}

	writeJSONResponse(w, &VideoResponse{
		Id:          ptr(id),
		OwnerId:     ptr(v.OwnerID),
		Title:       ptr(v.Title),
		Description: ptr(v.Description),
		Thumbnail:   ptr(v.Thumbnail),
		Url:         ptr(v.URL),
		Comments:    ptr(v.Comments),
		Likes:       ptr(v.Likes),
		Views:       ptr(v.Views),
		WatchTime:   ptr(v.WatchTime),
		Shares:      ptr(v.Shares),
		CreatedAt:   ptr(v.CreatedAt),
		UpdatedAt:   ptr(v.UpdatedAt),
	})
}

func (s *Server) PutVideosId(w http.ResponseWriter, r *http.Request, id string) {
	var req models.UpdateVideoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := s.s.UpdateVideo(r.Context(), req, id)
	if err != nil {
		http.Error(w, "Failed to update video", http.StatusInternalServerError)
		return
	}

	writeJSONResponse(w, &UserResponse{Id: ptr(id)})
}

func (s *Server) DeleteVideosId(w http.ResponseWriter, r *http.Request, id string) {
	err := s.s.DeleteVideo(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to delete video", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
