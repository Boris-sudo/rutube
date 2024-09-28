package recsys

import (
	"ColdStart/internal/pkg/log"
	"ColdStart/models"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type RecSys struct {
	logger *log.Logger
	db     *gorm.DB
}

// GetUserByIDHandler godoc
// @Summary retrieves a user with history by their ID
// @Description It does EXACTLY what it says.
// @Tags Video History
// @Accept json
// @Produce json
//
// @Router /recsys/user [post]
func (handler *RecSys) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract user ID from query parameters or request body
	type GetUserByIDRequest struct {
		UserId string `json:"user_id"`
	}

	var request GetUserByIDRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the GetUserByID function to retrieve the user
	user, err := GetUserByID(handler.db, request.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		}
		return
	}

	// Return the user as JSON
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// GetUserByID retrieves a user by ID from the database
func GetUserByID(db *gorm.DB, userId string) (*models.User, error) {
	var user models.User
	err := db.Where("id = ?", userId).
		Preload("VideoPreferences").
		Preload("VideoHistory").
		First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
