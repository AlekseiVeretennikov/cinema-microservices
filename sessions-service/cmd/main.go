// file: sessions-service/main.go
package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"

	"sessions-service/internal/handlers"
	"sessions-service/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbpool *pgxpool.Pool

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db_url := os.Getenv("DATABASE_URL")
	if db_url == "" {
		logger.Error("DATABASE_URL environment variable is not set")
		os.Exit(1)
	}
	var err error
	dbpool, err = pgxpool.New(context.Background(), db_url)
	if err != nil {
		logger.Error("Unable to connect to database", "error", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	movieRepo := repository.NewMovieRepo(dbpool)
	movieHandler := handlers.NewMovieHandler(movieRepo)

	http.HandleFunc("/healthcheck", healthCheckHandler)
	http.HandleFunc("POST /movies", movieHandler.CreateMovieHandler)

	port := ":8080"
	logger.Info("Starting sessions-service", "port", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		logger.Error("Could not start server", "error", err)
		os.Exit(1)
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	dbStatus := "ok"
	if err := dbpool.Ping(ctx); err != nil {
		dbStatus = "failed"
	}
	response := map[string]string{
		"status":          "ok",
		"service":         "sessions-service",
		"database_status": dbStatus,
	}
	w.Header().Set("Content-Type", "application/json")
	if dbStatus == "failed" {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(response)
}
