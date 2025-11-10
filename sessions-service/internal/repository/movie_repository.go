// file: sessions-service/internal/repository/movie_repository.go
package repository

import (
	"context"

	"sessions-service/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MovieRepo struct {
	DB *pgxpool.Pool
}

func NewMovieRepo(db *pgxpool.Pool) *MovieRepo {
	return &MovieRepo{DB: db}
}

func (r *MovieRepo) CreateMovie(ctx context.Context, movie models.Movie) (int, error) {
	var movieID int
	query := `INSERT INTO movies (title, description, duration_minutes) 
			  VALUES ($1, $2, $3) 
			  RETURNING movie_id`

	err := r.DB.QueryRow(ctx, query, movie.Title, movie.Description, movie.DurationMinutes).Scan(&movieID)
	if err != nil {
		return 0, err
	}
	return movieID, nil
}
