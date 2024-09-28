package models

type Video struct {
	Id          string  `json:"uuid" gorm:"not null;primary_key"`
	PubTime     string  `json:"pub_time" gorm:"not null"`
	Title       string  `json:"title" gorm:"not null"`
	Description string  `json:"description"`
	Likes       int     `json:"likes" gorm:"not null"`
	Dislikes    int     `json:"dislikes" gorm:"not null"`
	Duration    float32 `json:"duration" gorm:"not null"`
	CategoryID  string  `json:"category_id"`
}
