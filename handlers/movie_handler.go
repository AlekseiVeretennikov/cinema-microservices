// file: sessions-service/internal/handlers/movie_handler.go
package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"sessions-service/internal/models"
)

type MovieCreator interface {
	CreateMovie(ctx context.Context, movie models.Movie) (int, error)
}

type MovieHandler struct {
	repo MovieCreator
}

func NewMovieHandler(repo MovieCreator) *MovieHandler {
	return &MovieHandler{repo: repo}
}

func (h *MovieHandler) CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	movieID, err := h.repo.CreateMovie(r.Context(), movie)
	if err != nil {
		http.Error(w, "Failed to create movie", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Movie created successfully",
		"movie_id": movieID,
	})
}
