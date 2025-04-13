package db

import (
	"context"
	"github.com/uptrace/bun"
	"real-time-ranking/internal/models"
	"time"
)

func (t *testdatasource) CreateUser(ctx context.Context, username, email string) (string, error) {
	return "789", nil
}

func (t *testdatasource) GetUser(ctx context.Context, id string) (*models.User, error) {
	return &models.User{
		BaseModel: bun.BaseModel{},
		ID:        "789",
		Username:  "test-user-1",
		Email:     "test-user-1@example.com",
		CreatedAt: time.Time{},
	}, nil
}

func (t *testdatasource) UpdateUser(ctx context.Context, id string, username, email string) error {
	return nil
}

func (t *testdatasource) DeleteUser(ctx context.Context, id string) error {
	return nil
}
