package service

import (
	"context"
	"real-time-ranking/internal/db"
	"real-time-ranking/internal/models"
)

func (s *S) CreateUser(ctx context.Context, req models.CreateUserRequest) (string, error) {
	return db.DataSource().CreateUser(ctx, req.Username, req.Email)
}

func (s *S) GetUser(ctx context.Context, id string) (*models.User, error) {
	return db.DataSource().GetUser(ctx, id)
}

func (s *S) UpdateUser(ctx context.Context, id string, username, email string) error {
	return db.DataSource().UpdateUser(ctx, id, username, email)
}

func (s *S) DeleteUser(ctx context.Context, id string) error {
	return db.DataSource().DeleteUser(ctx, id)
}
