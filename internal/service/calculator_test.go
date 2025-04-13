package service_test

import (
	"real-time-ranking/internal/models"
	"real-time-ranking/internal/service"
	"testing"
)

func TestCalculateVideoScore(t *testing.T) {
	type args struct {
		req models.InteractionRequest
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "case_1", args: args{req: models.InteractionRequest{
			Views:     1,
			Likes:     1,
			Comments:  1,
			Shares:    1,
			WatchTime: 10,
		}}, want: 14},
		{name: "case_2", args: args{req: models.InteractionRequest{
			Views:     100,
			Likes:     5,
			Comments:  0,
			Shares:    0,
			WatchTime: 200,
		}}, want: 135},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.CalculateVideoScore(tt.args.req); got != tt.want {
				t.Errorf("CalculateVideoScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
