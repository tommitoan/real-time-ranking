package service

import (
	"context"
	"real-time-ranking/internal/models"
)

func (s *S) CreateUser(ctx context.Context, req models.CreateUserRequest) (string, error) {
	return s.db.CreateUser(ctx, req.Username, req.Email)
}

func (s *S) GetUser(ctx context.Context, id string) (*models.User, error) {
	return s.db.GetUser(ctx, id)
}

func (s *S) UpdateUser(ctx context.Context, id string, username, email string) error {
	return s.db.UpdateUser(ctx, id, username, email)
}

func (s *S) DeleteUser(ctx context.Context, id string) error {
	return s.db.DeleteUser(ctx, id)
}
