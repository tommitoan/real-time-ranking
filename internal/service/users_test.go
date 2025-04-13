package service_test

import (
	"github.com/uptrace/bun"
	"real-time-ranking/internal/models"
	"reflect"
	"testing"
	"time"
)

func TestS_CreateUser(t *testing.T) {
	tests := []struct {
		name    string
		req     models.CreateUserRequest
		want    string
		wantErr bool
	}{
		{name: "success", req: models.CreateUserRequest{}, want: "789", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateUser(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestS_DeleteUser(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{name: "success", id: "789", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.DeleteUser(ctx, tt.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestS_GetUser(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		want    *models.User
		wantErr bool
	}{
		{name: "success", id: "789", want: &models.User{
			BaseModel: bun.BaseModel{},
			ID:        "789",
			Username:  "test-user-1",
			Email:     "test-user-1@example.com",
			CreatedAt: time.Time{},
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetUser(ctx, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestS_UpdateUser(t *testing.T) {
	type args struct {
		id       string
		username string
		email    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "success", args: args{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.UpdateUser(ctx, tt.args.id, tt.args.username, tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
