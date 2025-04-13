package service

import "real-time-ranking/internal/models"

const (
	viewWeight      = 1.0
	likeWeight      = 3.0
	commentWeight   = 5.0
	shareWeight     = 4.0
	watchTimeWeight = 0.1
)

func CalculateVideoScore(req models.InteractionRequest) float64 {
	return float64(req.Views*viewWeight+
		req.Likes*likeWeight+
		req.Comments*commentWeight+
		req.Shares*shareWeight) + req.WatchTime*watchTimeWeight
}
