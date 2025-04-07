package repo

import "movie-crud/models"

type AuthRepoInteface interface {
	Register(request *models.Register) (string, error)
}
