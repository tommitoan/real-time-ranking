package db

import (
	"context"
	"real-time-ranking/internal/models"
)

var _testDataSource *testdatasource

type (
	TestServiceDataSource interface {
		TestUser
	}

	TestUser interface {
		CreateUser(ctx context.Context, username, email string) (string, error)
		GetUser(ctx context.Context, id string) (*models.User, error)
		UpdateUser(ctx context.Context, id string, username, email string) error
		DeleteUser(ctx context.Context, id string) error
	}

	TestVideo interface {
		CreateVideo(ctx context.Context, v *models.Video) (string, error)
		GetVideo(ctx context.Context, id string) (*models.Video, error)
		UpdateVideo(ctx context.Context, req models.UpdateVideoRequest, id string) error
		UpdateVideoReactions(ctx context.Context, id string, r *models.UpdateVideoReactionsRequest) (*models.Video, error)
		DeleteVideo(ctx context.Context, id string) error
	}
)

type testdatasource struct{}

func TestDataSource() ServiceDataSource {
	return _testDataSource
}
