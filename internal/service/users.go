package service

import (
	"context"
	"log"
	"real-time-ranking/internal/db"
	"real-time-ranking/internal/models"
)

func CreateUser(ctx context.Context, req models.CreateUserRequest) (string, error) {
	client := db.DataSource()
	id, err := client.CreateUser(ctx, req.Username, req.Email)
	if err != nil {
		log.Println(err)
	}
	return id, err
}
