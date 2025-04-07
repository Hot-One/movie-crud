package storage

import (
	"movie-crud/storage/postgres"
	"movie-crud/storage/repo"
)

type IStorage interface {
	UserService() repo.UserRepoInterface
	MovieService() repo.MovieRepoInterface
}

type storage struct {
	db *postgres.Postgres

	userService  repo.UserRepoInterface
	movieService repo.MovieRepoInterface
}

func NewStorage(db *postgres.Postgres) IStorage {
	return &storage{
		db:           db,
		userService:  postgres.NewUserRepository(db),
		movieService: postgres.NewMovieRepository(db),
	}
}

func (s *storage) UserService() repo.UserRepoInterface {
	return s.userService
}

func (s *storage) MovieService() repo.MovieRepoInterface {
	return s.movieService
}
