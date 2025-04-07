package storage

import (
	"movie-crud/config"
	"movie-crud/storage/postgres"
	"movie-crud/storage/repo"
)

type IStorage interface {
	UserService() repo.UserRepoInterface
	MovieService() repo.MovieRepoInterface
	AuthService() repo.AuthRepoInteface
}

type storage struct {
	db *postgres.Postgres

	userService  repo.UserRepoInterface
	movieService repo.MovieRepoInterface
	authService  repo.AuthRepoInteface
}

func NewStorage(db *postgres.Postgres, cfg config.Config) IStorage {
	return &storage{
		db:           db,
		userService:  postgres.NewUserRepository(db),
		movieService: postgres.NewMovieRepository(db),
		authService:  postgres.NewAuthRepository(db, cfg),
	}
}

func (s *storage) UserService() repo.UserRepoInterface {
	return s.userService
}

func (s *storage) MovieService() repo.MovieRepoInterface {
	return s.movieService
}

func (s *storage) AuthService() repo.AuthRepoInteface {
	return s.authService
}
