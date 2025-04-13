package service_test

import (
	"github.com/uptrace/bun"
	"real-time-ranking/internal/models"
	"reflect"
	"testing"
)

func TestS_CreateVideo(t *testing.T) {
	tests := []struct {
		name    string
		v       *models.Video
		want    string
		wantErr bool
	}{
		{name: "success", v: &models.Video{}, want: "video-456", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateVideo(ctx, tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateVideo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateVideo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestS_DeleteVideo(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{name: "success", id: "video-456", wantErr: false},
		{name: "failed", id: "video-000", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.DeleteVideo(ctx, tt.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteVideo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestS_GetVideo(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		want    *models.Video
		wantErr bool
	}{
		{name: "success", id: "video-456", want: &models.Video{
			BaseModel:   bun.BaseModel{},
			ID:          "video-456",
			Title:       "test-title",
			Description: "test-description",
			OwnerID:     "789",
		}, wantErr: false},
		{name: "failed", id: "video-000", wantErr: true, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetVideo(ctx, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVideo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVideo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestS_UpdateVideo(t *testing.T) {
	tests := []struct {
		name    string
		req     models.UpdateVideoRequest
		id      string
		wantErr bool
	}{
		{name: "success", req: models.UpdateVideoRequest{}, id: "video-456", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.UpdateVideo(ctx, tt.req, tt.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateVideo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
