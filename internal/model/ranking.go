package model

type VideoRanking struct {
	VideoID string  `json:"video_id"`
	Score   float64 `json:"score"`
}
