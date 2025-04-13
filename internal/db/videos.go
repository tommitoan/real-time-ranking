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

func (d *datasource) DeleteVideo(ctx context.Context, id string) error {
	_, err := d.db.NewDelete().
		Model((*models.Video)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
