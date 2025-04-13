package service

import (
	"context"
	"real-time-ranking/internal/db"
	"real-time-ranking/internal/models"
)

var _ Service = (*S)(nil)

type (
	Service interface {
		User
		Video
		Interactions
		Ranking
	}

	User interface {
		CreateUser(ctx context.Context, req models.CreateUserRequest) (string, error)
		GetUser(ctx context.Context, id string) (*models.User, error)
		UpdateUser(ctx context.Context, id string, username, email string) error
		DeleteUser(ctx context.Context, id string) error
	}

	Video interface {
		CreateVideo(ctx context.Context, v *models.Video) (string, error)
		GetVideo(ctx context.Context, id string) (*models.Video, error)
		UpdateVideo(ctx context.Context, req models.UpdateVideoRequest, id string) error
		DeleteVideo(ctx context.Context, id string) error
	}

	Interactions interface {
		PostInteractions(ctx context.Context, req models.InteractionRequest) error
	}

	Ranking interface {
		GetGlobalRankings(ctx context.Context, limit int64) ([]models.VideoRanking, error)
		GetUserRankings(ctx context.Context, limit int64, userID string) ([]models.VideoRanking, error)
	}
)

type S struct {
	db db.ServiceDataSource
}

func NewService(db db.ServiceDataSource) *S {
	return &S{
		db: db,
	}
}
