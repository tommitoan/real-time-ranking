package service

import (
	"context"
	"real-time-ranking/internal/models"
)

func (s *S) CreateVideo(ctx context.Context, v *models.Video) (string, error) {
	return s.db.CreateVideo(ctx, v)
}

func (s *S) GetVideo(ctx context.Context, id string) (*models.Video, error) {
	return s.db.GetVideo(ctx, id)
}

func (s *S) UpdateVideo(ctx context.Context, req models.UpdateVideoRequest, id string) error {
	return s.db.UpdateVideo(ctx, req, id)
}

func (s *S) DeleteVideo(ctx context.Context, id string) error {
	return s.db.DeleteVideo(ctx, id)
}
