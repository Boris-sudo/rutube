package main

import (
	"ColdStart/internal/app/ColdStart"
	"ColdStart/internal/pkg/config"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"

	_ "ColdStart/docs"
	httpSwagger "github.com/swaggo/http-swagger" // Swagger handler package
)

// @title Cold Start API
// @version 1.0
// @description Backend for Recsys program.

// @host 127.0.0.1:8080
// @BasePath /api/

func main() {
	cfg := config.NewConfig()
	r := chi.NewRouter()
	ColdStart.SetupApp(cfg, r)

	// Swagger API http://localhost:8080/swagger/index.html
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	log.Printf("ColdStart hosted on %v", cfg.GetString("http.base")+":"+cfg.GetString("http.port"))
	log.Fatal(http.ListenAndServe(cfg.GetString("http.base")+":"+cfg.GetString("http.port"), r))
}
