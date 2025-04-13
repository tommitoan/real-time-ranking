package db

import (
	"context"
	"errors"
	"github.com/uptrace/bun"
	"real-time-ranking/internal/models"
)

func (t *testdatasource) CreateVideo(ctx context.Context, v *models.Video) (string, error) {
	return "video-456", nil
}

func (t *testdatasource) GetVideo(ctx context.Context, id string) (*models.Video, error) {
	if id != "video-456" {
		return nil, errors.New("not found")
	}
	return &models.Video{
		BaseModel:   bun.BaseModel{},
		ID:          "video-456",
		Title:       "test-title",
		Description: "test-description",
		OwnerID:     "789",
	}, nil
}

func (t *testdatasource) UpdateVideo(ctx context.Context, req models.UpdateVideoRequest, id string) error {
	return nil
}

func (t *testdatasource) UpdateVideoReactions(ctx context.Context, id string, r *models.UpdateVideoReactionsRequest) (*models.Video, error) {
	return nil, nil
}

func (t *testdatasource) DeleteVideo(ctx context.Context, id string) error {
	if id != "video-456" {
		return errors.New("not found")
	}
	return nil
}
