package recsys

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

// SaveVideoPreferenceHandler saves the video preference for a user
// @Summary Save Video Preference
// @Description Saves the user's preference for a video based on user ID and video ID.
// @Tags Preferences
// @Accept json
// @Produce json
//
// @Param user_id     body string true "User's unique identifier" example("8195aaf7-0108-4d8c-be8d-fc4255686feb")
// @Param video_id    body string true "Video's unique identifier" example("6686fc28-e98e-4100-a00e-e180f15e5c75")
// @Param is_liked    body string false "Indicates if the video is liked" example("true")
// @Param is_disliked body string false "Indicates if the video is disliked" example("false")
//
// @Success 200 {object} map[string]string "Message indicating the save was successful"
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 500 {object} map[string]string "Failed to save video preference"
// @Router /recsys/preferences/save [post]
func (handler *RecSys) SaveVideoPreferenceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type SaveVideoPreferenceRequest struct {
		UserId     string `json:"user_id"`
		VideoId    string `json:"video_id"`
		IsLiked    bool   `json:"is_liked"`
		IsDisliked bool   `json:"is_disliked"`
	}

	var request SaveVideoPreferenceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := SaveVideoPreference(handler.db, request.UserId, request.VideoId, request.IsLiked, request.IsDisliked); err != nil {
		http.Error(w, "Failed to save video preference", http.StatusInternalServerError)
		return
	}

	// OK
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "Video preference saved successfully"})
	if err != nil {
		handler.logger.Debug("Failed to save video preference", zap.Error(err))
		return
	}
}

// UpdateVideoPreferenceHandler updates the video preference for a user
// @Summary Update Video Preference
// @Description Updates the user's preference for a video based on user ID and video ID.
// @Tags Preferences
// @Accept json
// @Produce json
//
// @Param user_id     body string true "User's unique identifier" example("8195aaf7-0108-4d8c-be8d-fc4255686feb")
// @Param video_id    body string true "Video's unique identifier" example("6686fc28-e98e-4100-a00e-e180f15e5c75")
// @Param is_liked    body string false "Indicates if the video is liked" example("true")
// @Param is_disliked body string false "Indicates if the video is disliked" example("false")
//
// @Success 200 {object} map[string]string "Message indicating the update was successful"
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 500 {object} map[string]string "Failed to update video preference"
// @Router /recsys/preferences/update [post]
func (handler *RecSys) UpdateVideoPreferenceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type UpdateVideoPreferenceRequest struct {
		UserId     string `json:"user_id"`
		VideoId    string `json:"video_id"`
		IsLiked    bool   `json:"is_liked"`
		IsDisliked bool   `json:"is_disliked"`
	}

	var request UpdateVideoPreferenceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := UpdateVideoPreference(handler.db, request.UserId, request.VideoId, request.IsLiked, request.IsDisliked); err != nil {
		http.Error(w, "Failed to update video preference", http.StatusInternalServerError)
		return
	}

	// OK
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "Video preference updated successfully"})
	if err != nil {
		handler.logger.Debug("Failed to update video preference", zap.Error(err))
		return
	}
}

// GetUserPreferencesHandler retrieves all preferences of a user
// @Summary Get User Preferences
// @Description Retrieves all video preferences for a user based on their user ID.
// @Tags Preferences
// @Accept json
// @Produce json
//
// @Param user_id body string true "User's unique identifier" example("8195aaf7-0108-4d8c-be8d-fc4255686feb")
//
// @Success 200 {array} models.UserVideoPreference "List of user's video preferences"
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 500 {object} map[string]string "Failed to retrieve user preferences"
// @Router /recsys/preferences/ [get]
func (handler *RecSys) GetUserPreferencesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type GetUserPreferencesRequest struct {
		UserId string `json:"user_id"`
	}

	var request GetUserPreferencesRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	preferences, err := GetUserPreferences(handler.db, request.UserId)
	if err != nil {
		http.Error(w, "Failed to retrieve user preferences", http.StatusInternalServerError)
		return
	}

	// OK
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(preferences)
	if err != nil {
		handler.logger.Debug("Failed to retrieve user preferences", zap.Error(err))
		return
	}
}

// GetVideoPreferenceHandler retrieves a user's preference on a single video
// @Summary Get User Video Preference
// @Description Retrieves the user's preference for a video based on user ID and video ID.
// @Tags Preferences
// @Accept json
// @Produce json
//
// @Param user_id body string true "User's unique identifier" example("8195aaf7-0108-4d8c-be8d-fc4255686feb")
// @Param video_id body string true "Video's unique identifier" example("6686fc28-e98e-4100-a00e-e180f15e5c75")
//
// @Success 200 {object} models.UserVideoPreference "User's preference for the specified video"
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 500 {object} map[string]string "Failed to retrieve video preference"
// @Router /recsys/preferences/video/ [get]
func (handler *RecSys) GetVideoPreferenceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type GetVideoPreferenceRequest struct {
		UserId  string `json:"user_id"`
		VideoId string `json:"video_id"`
	}

	var request GetVideoPreferenceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	preference, err := GetVideoPreference(handler.db, request.UserId, request.VideoId)
	if err != nil {
		http.Error(w, "Failed to retrieve video preference", http.StatusInternalServerError)
		return
	}

	// OK
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(preference)
	if err != nil {
		handler.logger.Debug("Failed to retrieve video preference", zap.Error(err))
		return
	}
}
