// file: sessions-service/internal/models/movie.go
package models

type Movie struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	DurationMinutes int    `json:"duration_minutes"`
}
