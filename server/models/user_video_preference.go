package models

import (
	"time"
)

type UserVideoPreference struct {
	UserId     string `gorm:"not null;index:,unique,composite:idx_user_video"`
	VideoId    string `gorm:"not null;index:,unique,composite:idx_user_video"`
	IsLiked    bool   `gorm:"default:false"`
	IsDisliked bool   `gorm:"default:false"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
