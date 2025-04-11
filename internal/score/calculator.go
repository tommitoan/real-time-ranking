package score

import "real-time-ranking/internal/model"

const (
	ViewWeight      = 1.0
	LikeWeight      = 3.0
	CommentWeight   = 5.0
	ShareWeight     = 4.0
	WatchTimeWeight = 0.1
)

func CalculateVideoScore(req model.InteractionRequest) float64 {
	return float64(req.Views)*ViewWeight +
		float64(req.Likes)*LikeWeight +
		float64(req.Comments)*CommentWeight +
		float64(req.Shares)*ShareWeight +
		req.WatchTime*WatchTimeWeight
}
