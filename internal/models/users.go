package models

import (
	"github.com/uptrace/bun"
	"time"
)

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
