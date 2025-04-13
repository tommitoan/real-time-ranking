package models

import (
	"github.com/uptrace/bun"
	"time"
)

type UpdateVideoRequest struct {
	Description string `json:"description,omitempty"`
	Thumbnail   string `json:"thumbnail,omitempty"`
	Title       string `json:"title,omitempty"`
}

type UpdateVideoReactionsRequest struct {
	Comments  int64 `json:"comments,omitempty"`
	Likes     int64 `json:"likes,omitempty"`
	Shares    int64 `json:"shares,omitempty"`
	Views     int64 `json:"views,omitempty"`
	WatchTime int64 `json:"watch_time,omitempty"`
}

type Video struct {
	bun.BaseModel `bun:"table:videos"`

	ID          string `bun:",pk"`
	Title       string `bun:",notnull"`
	Description string
	URL         string `bun:",notnull"`
	Thumbnail   string
	OwnerID     string `bun:",notnull"` // Foreign key to users(id)

	Views     int64 `bun:",default:0"`
	Likes     int64 `bun:",default:0"`
	Comments  int64 `bun:",default:0"`
	Shares    int64 `bun:",default:0"`
	WatchTime int64 `bun:",default:0"` // total watch time in seconds (or ms if you prefer)

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`

	Owner *User `bun:"rel:belongs-to,join:owner_id=id"`
}
