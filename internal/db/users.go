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

func (d *datasource) GetUser(ctx context.Context, id string) (*models.User, error) {
	user := new(models.User)
	err := d.db.NewSelect().Model(user).Where("id = ?", id).Limit(1).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *datasource) UpdateUser(ctx context.Context, id string, username, email string) error {
	_, err := d.db.NewUpdate().
		Model(&models.User{}).
		Set("username = ?", username).
		Set("email = ?", email).
		Where("id = ?", id).
		Exec(ctx)
	return err
}

func (d *datasource) DeleteUser(ctx context.Context, id string) error {
	_, err := d.db.NewDelete().
		Model(&models.User{}).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
