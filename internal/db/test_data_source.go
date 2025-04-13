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
)

type testdatasource struct{}

func TestDataSource() ServiceDataSource {
	return _testDataSource
}
