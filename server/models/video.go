package models

type Video struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Id          string `json:"video_id"`
	Views       int    `json:"views"`
	Likes       int    `json:"likes"`
	Dislikes    int    `json:"dislikes"`
}
