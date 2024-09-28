package models

type Video struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Id          string `json:"video_id"`
	Views       int    `json:"views"`
	Comments    int    `json:"comments"`
	Likes       int    `json:"likes"`
	Dislikes    int    `json:"dislikes"`
	IsLiked     bool   `json:"is_liked"`
	IsDisliked  bool   `json:"is_disliked"`
}
