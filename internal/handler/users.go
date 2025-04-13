package handler

import (
	"encoding/json"
	"net/http"
	"real-time-ranking/internal/models"
)

func (s *Server) PostUsers(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	id, err := s.s.CreateUser(r.Context(), req)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	writeJSONResponse(w, &UserResponse{Id: ptr(id), Username: ptr(req.Username)})
}

func (s *Server) GetUsersId(w http.ResponseWriter, r *http.Request, id string) {
	user, err := s.s.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusNotFound)
		return
	}
	writeJSONResponse(w, &UserResponse{Id: ptr(id), Username: ptr(user.Username), Email: ptr(user.Email)})
}

func (s *Server) PutUsersId(w http.ResponseWriter, r *http.Request, id string) {
	var req models.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := s.s.UpdateUser(r.Context(), id, req.Username, req.ID)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	writeJSONResponse(w, &UserResponse{Id: ptr(id), Username: ptr(req.Username)})
}

func (s *Server) DeleteUsersId(w http.ResponseWriter, r *http.Request, id string) {
	err := s.s.DeleteUser(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
