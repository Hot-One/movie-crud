package storage

import (
	"movie-crud/storage/postgres"
	"movie-crud/storage/repo"
)

type IStorage interface {
	UserService() repo.UserRepoInterface
}

type storage struct {
	db *postgres.Postgres

	userService repo.UserRepoInterface
}

func NewStorage(db *postgres.Postgres) IStorage {
	return &storage{
		db:          db,
		userService: postgres.NewUserRepository(db),
	}
}

func (s *storage) UserService() repo.UserRepoInterface {
	return s.userService
}
