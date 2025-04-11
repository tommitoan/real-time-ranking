package model

type InteractionRequest struct {
	VideoID   string  `json:"video_id"`
	UserID    string  `json:"user_id"`
	Views     int     `json:"views"`
	Likes     int     `json:"likes"`
	Comments  int     `json:"comments"`
	Shares    int     `json:"shares"`
	WatchTime float64 `json:"watch_time"`
}
