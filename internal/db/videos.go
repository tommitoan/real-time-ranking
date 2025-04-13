package db

import (
	"context"
	"github.com/google/uuid"
	"real-time-ranking/internal/models"
	"time"
)

func (d *datasource) CreateVideo(ctx context.Context, v *models.Video) (string, error) {
	v.ID = uuid.NewString()
	v.CreatedAt = time.Now().UTC()
	v.UpdatedAt = time.Now().UTC()
	_, err := d.db.NewInsert().Model(v).Exec(ctx)
	return v.ID, err
}

func (d *datasource) GetVideo(ctx context.Context, id string) (*models.Video, error) {
	video := new(models.Video)
	err := d.db.NewSelect().
		Model(video).
		Where("id = ?", id).
		Limit(1).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (d *datasource) UpdateVideo(ctx context.Context, req models.UpdateVideoRequest, id string) error {
	_, err := d.db.NewUpdate().
		Model(&models.Video{}).
		Set("description = ?", req.Description).
		Set("thumbnail = ?", req.Thumbnail).
		Set("title = ?", req.Title).
		Set("updated_at = ?", time.Now().UTC()).
		Where("id = ?", id).
		Exec(ctx)
	return err
}

func (d *datasource) UpdateVideoReactions(ctx context.Context, id string, r *models.UpdateVideoReactionsRequest) (*models.Video, error) {
	q := d.db.NewUpdate().Table("videos").Where("id = ?", id)

	if r.Likes != 0 {
		q = q.Set("likes = likes + ?", r.Likes)
	}
	if r.Comments != 0 {
		q = q.Set("comments = comments + ?", r.Comments)
	}
	if r.Shares != 0 {
		q = q.Set("shares = shares + ?", r.Shares)
	}
	if r.Views != 0 {
		q = q.Set("views = views + ?", r.Views)
	}
	if r.WatchTime != 0 {
		q = q.Set("watch_time = watch_time + ?", r.WatchTime)
	}

	q = q.Set("updated_at = ?", time.Now().UTC())

	var video models.Video
	if err := q.Returning("*").Scan(ctx, &video); err != nil {
		return nil, err
	}
	return &video, nil
}

func (d *datasource) DeleteVideo(ctx context.Context, id string) error {
	_, err := d.db.NewDelete().
		Model((*models.Video)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
