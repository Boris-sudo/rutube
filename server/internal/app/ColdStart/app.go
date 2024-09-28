package ColdStart

import (
	"ColdStart/internal/pkg/auth"
	"ColdStart/internal/pkg/dbha"
	"ColdStart/internal/pkg/log"
	"ColdStart/internal/pkg/mw"
	"ColdStart/internal/pkg/recsys"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var configGlobal *viper.Viper
var loggerGlobal *log.Logger

var middleware *mw.Middleware
var authHandler *auth.Auth
var recSys *recsys.RecSys

func setupPostgres() *gorm.DB {
	db := dbha.ConnectPostgres(loggerGlobal, configGlobal)
	return db
}

func setupMiddleware() {
	middleware = mw.New(loggerGlobal, configGlobal)
}

func setupAuth(db *gorm.DB) {
	authHandler = auth.New(loggerGlobal, configGlobal, db)
}

func setupRecSys(db *gorm.DB) {
	recSys = recsys.New(loggerGlobal, db)
}

func setupRoutes(r *chi.Mux) {
	loggerGlobal.Info("Initializing routes")
	r.Use(mw.LoggingMiddleware)
	r.Use(mw.CORSMiddleware)
	// <-- Auth -->
	// Input: Name, Email, Password
	r.Post("/api/accounts/register", authHandler.Register)
	// Input: Email, Password
	r.Post("/api/accounts/login", authHandler.Login)
	// Should be authenticated
	r.Post("/api/accounts/logout", authHandler.Logout)
	// Should be authenticated. Gets current user
	r.Get("/api/accounts/user", authHandler.Authenticate)

	// <-- History -->
	r.Post("/api/recsys/history/clear", recSys.ClearVideoHistoryHandler)
	r.Post("/api/recsys/history/save", recSys.SaveVideoHistoryHandler)
	r.Get("/api/recsys/history", recSys.GetVideoHistoryHandler)
	r.Get("/api/recsys/user", recSys.GetUserByIDHandler)

	// <-- Preferences -->
	r.Post("/api/recsys/preferences/save", recSys.SaveVideoPreferenceHandler)
	r.Post("/api/recsys/preferences/update", recSys.UpdateVideoPreferenceHandler)
	r.Get("/api/recsys/preferences/video", recSys.GetVideoPreferenceHandler)
	r.Get("/api/recsys/preferences", recSys.GetUserPreferencesHandler)

	// <-- Recsys -->
	r.Post("/api/recsys/videos", recSys.GetVideosByUserID)

	loggerGlobal.Info("Routes initialized")
}

func SetupApp(cfg *viper.Viper, r *chi.Mux) {
	configGlobal = cfg
	loggerGlobal = log.NewLogger(configGlobal)
	db := setupPostgres()
	setupMiddleware()
	setupAuth(db)
	setupRecSys(db)
	setupRoutes(r)
}
