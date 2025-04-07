package repo

import "movie-crud/models"

type MovieRepoInterface interface {
	CreateMovie(movie *models.Movie) error
	GetMovieByID(id uint) (*models.Movie, error)
	UpdateMovie(movie *models.Movie) (*models.Movie, error)
	DeleteMovie(id uint) error
	GetAllMovies(request models.MovieGetListRequest) (*models.MovieGetListResponse, error)
}
