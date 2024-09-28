package recsys

import (
	"ColdStart/internal/pkg/log"
	"encoding/json"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

// ClearVideoHistoryHandler godoc
// @Summary Clears user's video history
// @Description It does EXACTLY what it says.
// @Tags Video History
// @Accept json
// @Produce json
//
// @Param user_id body string true "User's UUID" example("6686fc28-e98e-4100-a00e-e180f15e5c75")
//
// @Success 200 {object} map[string]string "Video history cleared successfully"
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 500 {object} map[string]string "Failed to clear video history"
// @Router /recsys/history/clear [post]
func (handler *RecSys) ClearVideoHistoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type ClearVideoHistoryRequest struct {
		UserId string `json:"user_id"`
	}

	var request ClearVideoHistoryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := ClearVideoHistory(handler.db, request.UserId); err != nil {
		http.Error(w, "Failed to clear video history", http.StatusInternalServerError)
		return
	}

	// OK
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "Video history cleared successfully"})
	if err != nil {
		handler.logger.Debug("Failed to clear video history", zap.Error(err))
		return
	}
}

// SaveVideoHistoryHandler godoc
// @Summary Save video to user's history
// @Description Adds a video entry to a user's history based on their user ID and video ID.
// @Tags Video History
// @Accept json
// @Produce json
//
// @Param user_id body string true "User's UUID" example("6686fc28-e98e-4100-a00e-e180f15e5c75")
// @Param video_id body string true "Video's UUID" example("8195aaf7-0108-4d8c-be8d-fc4255686feb")
//
// @Success 200 {object} map[string]string "Video history saved successfully"
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 500 {object} map[string]string "Failed to save video history"
// @Router /recsys/history/save [post]
func (handler *RecSys) SaveVideoHistoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type SaveVideoHistoryRequest struct {
		UserId  string `json:"user_id"`
		VideoId string `json:"video_id"`
	}

	var request SaveVideoHistoryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := SaveVideoHistory(handler.db, request.UserId, request.VideoId); err != nil {
		http.Error(w, "Failed to save video history", http.StatusInternalServerError)
		return
	}

	// OK
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "Video history saved successfully"})
	if err != nil {
		handler.logger.Debug("Failed to save video history", zap.Error(err))
		return
	}
}

// GetVideoHistoryHandler godoc
// @Summary Get user's video history
// @Description Retrieves the video history of a user based on their user ID.
// @Tags Video History
// @Accept json
// @Produce json
//
// @Param user_id body string true "User's UUID" example("8195aaf7-0108-4d8c-be8d-fc4255686feb")
//
// @Success 200 {array} models.UserVideoHistory "User's video history"
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 500 {object} map[string]string "Failed to retrieve video history"
// @Router /recsys/history/ [get]
func (handler *RecSys) GetVideoHistoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type GetUserHistoryRequest struct {
		UserId string `json:"user_id"`
	}

	var req GetUserHistoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Retrieve user's video history
	history, err := GetUserHistory(handler.db, req.UserId)
	if err != nil {
		http.Error(w, "Failed to retrieve video history", http.StatusInternalServerError)
		return
	}

	// OK
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(history)
	if err != nil {
		handler.logger.Debug("Failed to save video history", zap.Error(err))
		return
	}
}

func New(logger *log.Logger, db *gorm.DB) *RecSys {
	return &RecSys{
		logger: logger,
		db:     db,
	}
}
