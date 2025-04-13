package models

import (
	"github.com/uptrace/bun"
	"time"
)

type CreateVideoRequest struct {
	Description string `json:"description,omitempty"`
	OwnerId     string `json:"owner_id"`
	Thumbnail   string `json:"thumbnail,omitempty"`
	Title       string `json:"title"`
	Url         string `json:"url"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UpdateUserRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        string    `bun:",pk"`
	Username  string    `bun:",notnull"`
	Email     string    `bun:",notnull,unique"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
