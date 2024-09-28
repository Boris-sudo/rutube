package models

type User struct {
	Id    string `json:"uuid" gorm:"not null;primary_key"`
	Login string `json:"login" gorm:"unique;not null;default:null"`
	Email string `json:"email" gorm:"unique"`

	Name    string `json:"name" gorm:"default:null"`
	Surname string `json:"surname" gorm:"default:null"`
	Region  string `json:"region" gorm:"default:null"`
	City    string `json:"city" gorm:"default:null"`

	Password string `json:"-" gorm:"not null;default:null"`
	Salt     []byte `json:"-" gorm:"not null;default:null"`

	// Stats
	VideoPreferences []UserVideoPreference `json:"video_preferences,omitempty" gorm:"foreignKey:UserId"`
	VideoHistory     []UserVideoHistory    `json:"video_history,omitempty" gorm:"foreignKey:UserId"`
}
