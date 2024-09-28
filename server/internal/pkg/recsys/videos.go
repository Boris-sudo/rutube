package recsys

import (
	"ColdStart/models"
	"bytes"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

// GetVideosByUserID godoc
// @Summary      Get recommended videos for a user
// @Description  This API sends a user object to ML backend providing recommendations based on the user's history
// @Tags         recommendations
// @Accept       json
// @Produce      json
// @Param        user_id  body string  true  "User's ID"
// @Success      200   {array} models.Video "List of recommended videos"
// @Failure      400   {object} map[string]string "Invalid request payload"
// @Failure      500   {object} map[string]string "Server error"
// @Router       /recsys/videos [post]
func (handler *RecSys) GetVideosByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type userRequest struct {
		Id string `json:"user_id"`
	}

	var UR userRequest
	if err := json.NewDecoder(r.Body).Decode(&UR); err != nil {
		handler.logger.Debug("Error decoding request", zap.Error(err))
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := GetUserByID(handler.db, UR.Id)
	if err != nil {
		handler.logger.Debug("Error getting user", zap.Error(err))
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Convert user model to JSON
	userJson, err := json.Marshal(user)
	if err != nil {
		handler.logger.Debug("Error marshaling user model", zap.Error(err))
		http.Error(w, "Failed to process user data", http.StatusInternalServerError)
		return
	}

	handler.logger.Debug("Sending user to analytics", zap.Any("User", user))
	// Create a request to the external API
	apiUrl := "http://localhost:5000/api/predicted_videos"
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(userJson))
	if err != nil {
		handler.logger.Debug("Error creating request", zap.Error(err))
		http.Error(w, "Failed to create request to external API", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		handler.logger.Debug("Error making API request", zap.Error(err))
		http.Error(w, "Failed to contact external API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read and process the response
	if resp.StatusCode != http.StatusOK {
		handler.logger.Debug("Error response from external API", zap.Int("status_code", resp.StatusCode))
		http.Error(w, "Error from external API", resp.StatusCode)
		return
	}

	var recommendedVideos []models.Video
	if err := json.NewDecoder(resp.Body).Decode(&recommendedVideos); err != nil {
		handler.logger.Debug("Error decoding API response", zap.Error(err))
		http.Error(w, "Failed to decode API response", http.StatusInternalServerError)
		return
	}

	// Fetch user's video preferences from the database
	var preferences []models.UserVideoPreference
	if err := handler.db.Where("user_id = ?", user.Id).Find(&preferences).Error; err != nil {
		handler.logger.Debug("Error fetching video preferences", zap.Error(err))
		http.Error(w, "Failed to fetch user preferences", http.StatusInternalServerError)
		return
	}

	// Map preferences
	prefMap := make(map[string]models.UserVideoPreference)
	for _, pref := range preferences {
		prefMap[pref.VideoId] = pref
	}

	for i, video := range recommendedVideos {
		if pref, found := prefMap[video.Id]; found {
			recommendedVideos[i].IsLiked = pref.IsLiked
			recommendedVideos[i].IsDisliked = pref.IsDisliked
		}
	}

	handler.logger.Debug("Received videos", zap.Any("recommended_videos", recommendedVideos))

	// Send the recommended videos back
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(recommendedVideos); err != nil {
		handler.logger.Debug("Error sending response", zap.Error(err))
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
	}
}
