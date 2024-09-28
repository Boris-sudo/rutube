package recsys

import (
	"ColdStart/models"
	"gorm.io/gorm"
)

// SaveVideoPreference sets the video preference for a user
func SaveVideoPreference(db *gorm.DB, userId string, videoId string, isLiked bool, isDisliked bool) error {
	// Check if the preference already exists
	var existingPreference models.UserVideoPreference
	if err := db.Where("user_id = ? AND video_id = ?", userId, videoId).First(&existingPreference).Error; err == nil {
		// If it exists, update the preference
		existingPreference.IsLiked = isLiked
		existingPreference.IsDisliked = isDisliked
		return db.Save(&existingPreference).Error
	}

	// If it does not exist, create a new preference
	newPreference := models.UserVideoPreference{
		UserId:     userId,
		VideoId:    videoId,
		IsLiked:    isLiked,
		IsDisliked: isDisliked,
	}
	return db.Create(&newPreference).Error
}

// UpdateVideoPreference updates the video preference for a user
func UpdateVideoPreference(db *gorm.DB, userId string, videoId string, isLiked bool, isDisliked bool) error {
	// Update the preference
	return db.Model(&models.UserVideoPreference{}).
		Where("user_id = ? AND video_id = ?", userId, videoId).
		Update("is_liked", isLiked).Update("is_disliked", isDisliked).Error
}

// GetUserPreferences retrieves all preferences of a user
func GetUserPreferences(db *gorm.DB, userId string) ([]models.UserVideoPreference, error) {
	var preferences []models.UserVideoPreference
	err := db.Where("user_id = ?", userId).Find(&preferences).Error
	if err != nil {
		return nil, err
	}
	return preferences, nil
}

// GetVideoPreference retrieves the user's preference on a single video
func GetVideoPreference(db *gorm.DB, userId string, videoId string) (models.UserVideoPreference, error) {
	var preference models.UserVideoPreference
	err := db.Where("user_id = ? AND video_id = ?", userId, videoId).First(&preference).Error
	if err != nil {
		return models.UserVideoPreference{}, err
	}
	return preference, nil
}
