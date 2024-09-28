package models

import "time"

type UserVideoHistory struct {
	ID        uint   `gorm:"primaryKey"`
	UserId    string `gorm:"not null;index:idx_user_video"`
	VideoId   string `gorm:"not null;index:idx_user_video"`
	CreatedAt time.Time
}
