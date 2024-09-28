package recsys

import (
	"ColdStart/models"
	"gorm.io/gorm"
	"time"
)

// ClearVideoHistory clears user's video history
func ClearVideoHistory(db *gorm.DB, userId string) error {
	if err := db.Where("user_id = ?", userId).Delete(&models.UserVideoHistory{}).Error; err != nil {
		return err
	}
	return nil
}

// SaveVideoHistory saves a video in the user's history
func SaveVideoHistory(db *gorm.DB, userId string, videoId string) error {
	history := models.UserVideoHistory{
		UserId:    userId,
		VideoId:   videoId,
		CreatedAt: time.Now(),
	}

	// Save the history entry to the database
	err := db.Create(&history).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUserHistory retrieves a user's video history based on their user ID (TODO pagination)
func GetUserHistory(db *gorm.DB, userId string) ([]models.UserVideoHistory, error) {
	var history []models.UserVideoHistory

	// Fetch all video entries
	err := db.Where("user_id = ?", userId).Order("created_at DESC").Find(&history).Error
	if err != nil {
		return nil, err
	}

	return history, nil
}
