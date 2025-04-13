package models

type VideoRanking struct {
	VideoID string  `json:"video_id"`
	Score   float64 `json:"score"`
}

type InteractionRequest struct {
	VideoID   string  `json:"video_id"`
	UserID    string  `json:"user_id"`
	Views     int64   `json:"views"`
	Likes     int64   `json:"likes"`
	Comments  int64   `json:"comments"`
	Shares    int64   `json:"shares"`
	WatchTime float64 `json:"watch_time"`
}
