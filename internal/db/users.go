package db

import (
	"context"
	"github.com/google/uuid"
	"real-time-ranking/internal/models"
	"time"
)

func (d *datasource) CreateUser(ctx context.Context, username, email string) (string, error) {
	user := &models.User{
		ID:        uuid.NewString(),
		Username:  username,
		Email:     email,
		CreatedAt: time.Now().UTC(),
	}

	_, err := d.db.NewInsert().
		Model(user).
		Exec(ctx)
	return user.ID, err
}
